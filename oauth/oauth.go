package oauth

import (
	"github.com/tfregonese/bookstore_utils-go/rest_errors"
	"net/http"
	"strconv"
)

const headerXClientId = "X-Client-Id"

func IsPublic(request *http.Request) bool {
	if request == nil {
		return true
	}
	return request.Header.Get("X-Public") == "true"
}

func GetClientId(request *http.Request) int64 {
	if request == nil {
		return 0
	}
	clientId, err := strconv.ParseInt(request.Header.Get(headerXClientId), 10, 64)
	if err != nil {
		return 0
	}
	return clientId
}

func AuthenticateRequest(request *http.Request) *rest_errors.RestErr {
	if request == nil {
		return nil
	}

	return nil
}
