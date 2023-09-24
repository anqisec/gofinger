package request

import (
	"net/http"
	"strings"
)

func GetHeader(response *http.Response) string {
	var headers string
	for key, values := range response.Header {
		headers += key + ": " + strings.Join(values, ", ") + "\n"
	}
	return headers
}
