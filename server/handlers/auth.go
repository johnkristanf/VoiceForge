package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/johnkristanf/VoiceForge/server/auth"
	"github.com/johnkristanf/VoiceForge/server/types"
	"github.com/johnkristanf/VoiceForge/server/utils"
	"golang.org/x/crypto/bcrypt"
)

func (s *ApiServer) SignUpHandler(res http.ResponseWriter, req *http.Request) error {

	if err := utils.HttpMethod(http.MethodPost, req); err != nil {
		return err
	}

	var signUpCredentials *types.SignupCredentials

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &signUpCredentials); err != nil {
		return err
	}

	if err := s.database.SignUp(signUpCredentials); err != nil {
		return err
	}

	return utils.WriteJson(res, http.StatusOK, map[string]bool{"Signup": true})

}

func (s *ApiServer) LoginHandler(res http.ResponseWriter, req *http.Request) error {

	if err := utils.HttpMethod(http.MethodPost, req); err != nil {
		return err
	}

	var loginCredentials *types.LoginCredentials
	if err := json.NewDecoder(req.Body).Decode(&loginCredentials); err != nil {
		return err
	}

	user, emailErr := s.database.CheckEmailExist(loginCredentials.Email)
	if emailErr != nil {
		return emailErr
	}

	if user == nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginCredentials.Password)) != nil {
		return utils.WriteJson(res, http.StatusNotFound, map[string]string{"Invalid_Credentials": "Incorrect Email or Password"})
	}

	if user.Verification_Token == "Unverified" {

		verificationCode, err := s.smtpClient.SendVerificationEmail(loginCredentials.Email)
		if err != nil {
			return err
		}

		hashedCode, err := bcrypt.GenerateFromPassword([]byte(strconv.Itoa(int(verificationCode))), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		verificationToken, err := auth.GenerateVerificationToken(user.ID, user.Email, string(hashedCode))
		if err != nil {
			return err
		}

		utils.SetCookie(res, verificationToken, time.Now().Add(30 * time.Minute), "Verification_Token")

		return utils.WriteJson(res, http.StatusUnauthorized, map[string]bool{"Need_Verification": true})

	}

	access_token, err := auth.GenerateAccessToken(user.ID, user.Email)
	if err != nil {
		return err
	}

	refreshToken, refreshErr := auth.GenerateRefreshToken(user.ID, user.Email)
	if refreshErr != nil {
		return refreshErr
	}

	utils.SetCookie(res, access_token, time.Now().Add(30 * time.Minute), "Access_Token")

	utils.SetCookie(res, refreshToken, time.Now().Add(3 * 24 * time.Hour), "Refresh_Token")

	return utils.WriteJson(res, http.StatusOK, map[string]bool{"Login": true})

}

func (s *ApiServer) FetchUserDatahandler(res http.ResponseWriter, req *http.Request) error {

	if err := utils.HttpMethod(http.MethodGet, req); err != nil {
		return err
	}

	userPayload, ok := req.Context().Value("jwt_payload").(*types.JWTPayloadClaims)
	if !ok {
		return fmt.Errorf("failed to retrieve user data from context")
	}

	return utils.WriteJson(res, http.StatusOK, map[string]any{
		"user_id": userPayload.ID,
		"email":   userPayload.Email,
	})
}


// --------------------------------TOKEN HANDLERS -----------------------------------------

func (s *ApiServer) RefreshTokenHandler(res http.ResponseWriter, req *http.Request) {

	refreshTokenCookie, err := req.Cookie("Refresh_Token")
	if err != nil {
		utils.WriteJson(res, http.StatusUnauthorized, "Unauthorized: Refresh Token Not Found")
		return
	}

	token, err := jwt.Parse(refreshTokenCookie.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("REFRESH_TOKEN_JWTSECRET")), nil
	})

	if err != nil || !token.Valid {
		utils.WriteJson(res, http.StatusUnauthorized, "Unauthorized: Access is Denied Due to Invalid Credentials")
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(res, "Invalid token claims", http.StatusInternalServerError)
		return
	}

	userID := int64(claims["user_id"].(float64))
	email := claims["email"].(string)

	access_token, err := auth.GenerateAccessToken(userID, email)
	if err != nil {
		utils.WriteJson(res, http.StatusUnauthorized, "Failed to generate access token")
		return
	}

	utils.SetCookie(res, access_token, time.Now().Add(30 * time.Minute), "Access_Token")

	utils.WriteJson(res, http.StatusOK, map[string]bool{"New_Access_Token_Generated": true})

}

func (s *ApiServer) VerifyUserHandler(res http.ResponseWriter, req *http.Request) error {

	if err := utils.HttpMethod(http.MethodPost, req); err != nil {
		return err
	}

	var code *types.VerificationCode
	errorChan := make(chan error, 1)

	codeCookie, err := req.Cookie("Verification_Token")
	if err != nil {
		return utils.WriteJson(res, http.StatusUnauthorized, map[string]string{"ERROR": "Verification Token Not Found"})
	}

	userClaims, parseErr := ParseVerificationClaims(codeCookie)
	if parseErr != nil {
		return utils.WriteJson(res, http.StatusUnauthorized, map[string]string{"ERROR": parseErr.Error()})
	}

	body, readErr := io.ReadAll(req.Body)
	if readErr != nil {
		return readErr
	}

	if err := json.Unmarshal(body, &code); err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userClaims.HashedCode), []byte(code.Code)); err != nil {
		return utils.WriteJson(res, http.StatusBadRequest, map[string]string{"ERROR": "Incorrect Verification Code"})
	}

	access_token, err := auth.GenerateAccessToken(userClaims.ID, userClaims.Email)
	if err != nil {
		return err
	}

	refreshToken, refreshErr := auth.GenerateRefreshToken(userClaims.ID, userClaims.Email)
	if refreshErr != nil {
		return refreshErr
	}

	utils.SetCookie(res, access_token, time.Now().Add(30 *time.Minute), "Access_Token")

	utils.SetCookie(res, refreshToken, time.Now().Add(3*24*time.Hour), "Refresh_Token")

	go func() {
		if err := s.database.VerifyUser(userClaims.ID, userClaims.HashedCode); err != nil {
			errorChan <- err
		}
	}()
	close(errorChan)

	if err := <-errorChan; err != nil {
		return err
	}

	Verification_TokenCookie := &http.Cookie{
		Name:    "Verification_Token",
		Value:   "",
		Expires: time.Now().AddDate(0, 0, -1),
		Path:    "/",
	}

	http.SetCookie(res, Verification_TokenCookie)

	return utils.WriteJson(res, http.StatusOK, map[string]bool{"Verified": true})

}

func ParseVerificationClaims(cookie *http.Cookie) (*types.ParseVerificationClaims, error) {

	token, err := jwt.ParseWithClaims(cookie.Value, &types.ParseVerificationClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("VERIFICATION_TOKEN_JWTSECRET")), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("Unauthorized: Access is denied due to invalid credentials")
	}

	user, ok := token.Claims.(*types.ParseVerificationClaims)
	if !ok {
		return nil, fmt.Errorf("Failed to Claim Payload From Struct")
	}

	if time.Unix(user.ExpiresAt, 0).Sub(time.Now()) < 0 {
		return nil, fmt.Errorf("Unauthorized")
	}

	return &types.ParseVerificationClaims{
		ID:         user.ID,
		Email:      user.Email,
		HashedCode: user.HashedCode,
	}, nil

}

func (s *ApiServer) LogoutHandler(res http.ResponseWriter, req *http.Request) error {

	if err := utils.HttpMethod(http.MethodPost, req); err != nil {
		return err
	}

	Access_TokenCookie := &http.Cookie{
		Name:    "Access_Token",
		Value:   "",
		Expires: time.Now().AddDate(0, 0, -1),
		Path:    "/",
	}

	Refresh_TokenCookie := &http.Cookie{
		Name:    "Refresh_Token",
		Value:   "",
		Expires: time.Now().AddDate(0, 0, -1),
		Path:    "/",
	}

	http.SetCookie(res, Access_TokenCookie)
	http.SetCookie(res, Refresh_TokenCookie)

	return utils.WriteJson(res, http.StatusOK, true)
}
