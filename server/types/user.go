package types
import (
	"github.com/golang-jwt/jwt/v5"
)

type SignupCredentials struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginEmail struct{
	Email string `json:"email"`
}

type LoginCredentials struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserEmailExist struct{
	ID int64 `json:"user_id"`
	Email string `json:"email"`
	Password string `json:"password"`
	Verification_Token string `json:"verification_token"`
}

type JWTPayloadClaims struct{
	ID int64 `json:"user_id"`
	Email string `json:"email"`
	ExpiresAt int64 `json:"exp"`
	jwt.RegisteredClaims
}


type VerificationCode struct{
	Code string `json:"verification_code"`
}

type ParseVerificationClaims struct{
	ID int64 `json:"user_id"`
	Email string `json:"email"`
	HashedCode string `json:"hashedCode"`
	ExpiresAt int64 `json:"exp"`
	jwt.RegisteredClaims
}