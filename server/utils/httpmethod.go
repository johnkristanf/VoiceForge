package utils

import (
	"fmt"
	"net/http"
)

func HttpMethod(method string, req *http.Request) error {
	if req.Method != method{
	   return fmt.Errorf("Method not Allowed %s", req.Method)
	}

	return nil
}