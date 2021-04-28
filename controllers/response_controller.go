package controllers

import (
	"encoding/json"
	"net/http"

	models "github.com/tubes/models"
)

func sendSuccessResponse(w http.ResponseWriter) {
	var response models.UserResponse
	response.Status = 200
	response.Message = "Success"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendUserSuccessResponse(w http.ResponseWriter, users []models.User) {
	var response models.UserResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = users
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendFilmSuccessResponse(w http.ResponseWriter, films []models.Film) {
	var response models.FilmResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = films
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendSubscriptionSuccessResponse(w http.ResponseWriter, subscriptions []models.Subscription) {
	var response models.SubscriptionResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = subscriptions
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendHistorySuccessResponse(w http.ResponseWriter, histories []models.History) {
	var response models.HistoryResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = histories
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter) {
	var response models.ErrorResponse
	response.Status = 400
	response.Message = "Failed"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendWatchErrorResponse(w http.ResponseWriter) {
	var response models.ErrorResponse
	response.Status = 400
	response.Message = "Failed, not yet subscribe"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendUnauthorizedResponse(w http.ResponseWriter) {
	var response models.ErrorResponse
	response.Status = 401
	response.Message = "Unauthorized Access"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
