package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errorResponse struct {
	Message string `json:"msg"`
}

func newErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	fmt.Println(message)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"msg": message,
	})
}