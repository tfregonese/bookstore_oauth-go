package oauth

import "net/http"

func IsPublic(request *http.Request) bool {
	if request == nil {
		return true
	}
	return request.Header.Get("X-Public") == "true"
}
