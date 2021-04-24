package controllers

import (
	"encoding/json"
	"net/http"
)

func sendSuccessResponse(w http.ResponseWriter) {
	var response Response
	response.Status = 200
	response.Message = "success"
	w.Header().Set("Content-Type", "application/json")
}

func sendErrorResponse(w http.ResponseWriter) {
	var response Response
	response.Status = 400
	response.Message = "Failed"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendUnAuthorizedResponse(w http.ResponseWriter) {
	var response Response
	response.Status = 401
	response.Message = "Unauthorized Access"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
