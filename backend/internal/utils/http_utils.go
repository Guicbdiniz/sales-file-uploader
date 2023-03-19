package utils

import (
	"fmt"
	"net/http"
)

// SendTextResponse writes a HTTP code and a string message to the response.
func SendTextResponse(res http.ResponseWriter, httpCode int, message string) {
	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(httpCode)
	res.Write([]byte(message))
}

// SendJsonResponse writes a HTTP code and a json message to the response.
func SendJsonResponse(res http.ResponseWriter, httpCode int, message []byte) {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Content-Length", fmt.Sprintf("%d", len(message)))
	res.WriteHeader(httpCode)
	res.Write(message)
}

// SendJsonErrorResponse writes a HTTP code and a json message with an error to the response.
func SendJsonErrorResponse(res http.ResponseWriter, httpCode int, err error) {
	jsonErrorResponse, _ := MarshalJsonErrorResponse(err.Error())

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Content-Length", fmt.Sprintf("%d", len(jsonErrorResponse)))
	res.WriteHeader(httpCode)
	res.Write([]byte(jsonErrorResponse))
}
