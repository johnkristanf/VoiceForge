package utils

import (
	"net/http"
	"time"
)

func SetCookie(res http.ResponseWriter, value string, expire time.Time, cookieName string){

	cookie := &http.Cookie{
		Name: cookieName,
		Value: value,
		Expires: expire,
		Path: "/",
		HttpOnly: true,
	}

	http.SetCookie(res, cookie)
}