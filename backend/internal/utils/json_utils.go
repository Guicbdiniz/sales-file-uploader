package utils

import (
	"encoding/json"
)

// Generic json response with data and a error text.
type JsonResponse[T any] struct {
	Data      T      `json:"data"`
	ErrorText string `json:"errorText"`
}

// UnmarshalJsonResponse generates a json object from a byte array.
func UnmarshalJsonResponse[T any](data []byte) (JsonResponse[T], error) {
	var t JsonResponse[T]
	err := json.Unmarshal(data, &t)

	if err != nil {
		return t, err
	}

	return t, nil
}

// MarshalJsonResponse generates a json object as an byte array with the
// required data.
func MarshalJsonResponse[T any](t T) ([]byte, error) {
	jsonResponse := JsonResponse[T]{Data: t, ErrorText: ""}

	jsonData, err := json.Marshal(jsonResponse)

	if err != nil {
		return []byte(""), err
	}

	return jsonData, nil
}

// MarshalJsonErrorResponse generates a json object as an byte array with the
// required error text.
func MarshalJsonErrorResponse(errorText string) ([]byte, error) {
	jsonResponse := JsonResponse[string]{Data: "", ErrorText: errorText}

	jsonData, err := json.Marshal(jsonResponse)

	if err != nil {
		return []byte(""), err
	}

	return jsonData, nil
}
