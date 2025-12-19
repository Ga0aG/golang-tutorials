package handler

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse 错误响应
type ErrorResponse struct {
	Error string `json:"error"`
}

// respondWithError 返回错误响应
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, ErrorResponse{Error: message})
}

// respondWithJSON 返回 JSON 响应
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Failed to marshal response"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
