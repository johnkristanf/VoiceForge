package auth

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/johnkristanf/VoiceForge/server/types"
)

var (
	accessTokenDuration   = 15 * time.Minute
	refreshTokenDuration  = 3 * 24 * time.Hour
)

func GenerateAccessToken(user_id int64, email string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user_id,
		"email": email,
		"exp" : time.Now().Add(accessTokenDuration).Unix(),  
	})


	return token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_JWTSECRET")))
}


func GenerateVerificationToken(user_id int64, email string, hashedCode string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user_id,
		"email": email,
		"hashedCode": hashedCode,
		"exp" : time.Now().Add(accessTokenDuration).Unix(),  
	})


	return token.SignedString([]byte(os.Getenv("VERIFICATION_TOKEN_JWTSECRET")))
}


func GenerateRefreshToken(user_id int64, email string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user_id,
		"email": email,
		"exp": time.Now().Add(refreshTokenDuration).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("REFRESH_TOKEN_JWTSECRET")))
}


func AuthenticationMiddleWare(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func (res http.ResponseWriter, req *http.Request)  {

		cookie, cookieErr := req.Cookie("Access_Token")
	    if cookieErr != nil{
		    http.Error(res, "Unauthorized: Access is denied due to invalid credentials", http.StatusUnauthorized)
            return
	    }

	    token, err := jwt.ParseWithClaims(cookie.Value, &types.JWTPayloadClaims{}, func(t *jwt.Token) (interface{}, error) {
		    return []byte(os.Getenv("ACCESS_TOKEN_JWTSECRET")), nil
	    })


	    if err != nil || !token.Valid{
			http.Error(res, "Unauthorized: Access is denied due to invalid credentials", http.StatusUnauthorized)
			return
	    }

	    user, ok := token.Claims.(*types.JWTPayloadClaims)
	    if !ok {
		    http.Error(res, "Failed to Claim Payload From Struct", http.StatusInternalServerError)
			return
	    }

		if time.Unix(user.ExpiresAt, 0).Sub(time.Now()) < 0 {
			http.Error(res, "Unauthorized", http.StatusUnauthorized)
			return
		}

	    ctx := context.WithValue(req.Context(), "jwt_payload", user)

	    req = req.WithContext(ctx)

		next.ServeHTTP(res, req)

	})

	
}