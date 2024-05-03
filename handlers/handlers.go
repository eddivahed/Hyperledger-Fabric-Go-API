package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-api/errors"
	"go-api/logging"
	"go-api/services"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the API!"))
}

func Register(w http.ResponseWriter, r *http.Request) {
	var reqData services.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to decode request payload")
		errors.HandleError(w, errors.NewAPIError(http.StatusBadRequest, "Invalid request payload"))
		return
	}

	err = services.Register(reqData.Username, reqData.Password)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to register user")
		errors.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func Minter(w http.ResponseWriter, r *http.Request) {
	var reqData services.MintRequest
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to decode request payload")
		errors.HandleError(w, errors.NewAPIError(http.StatusBadRequest, "Invalid request payload"))
		return
	}

	err = services.Mint(reqData.Username, reqData.Value)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to mint tokens")
		errors.HandleError(w, err)
		return
	}

	response := services.MintResponse{
		Message:  "ok",
		Username: reqData.Username,
		Value:    strconv.Itoa(reqData.Value),
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func Balancer(w http.ResponseWriter, r *http.Request) {
	var reqData services.BalanceRequest
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to decode request payload")
		errors.HandleError(w, errors.NewAPIError(http.StatusBadRequest, "Invalid request payload"))
		return
	}

	balance, err := services.ClientAccountBalance(reqData.Username)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to get account balance")
		errors.HandleError(w, err)
		return
	}

	response := services.BalanceResponse{
		Message: "ok",
		Value:   balance,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func Transferer(w http.ResponseWriter, r *http.Request) {
	var reqData services.TransferRequest
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to decode request payload")
		errors.HandleError(w, errors.NewAPIError(http.StatusBadRequest, "Invalid request payload"))
		return
	}

	err = services.Transfer(reqData.Username, reqData.Receiver, reqData.Value)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to transfer tokens")
		errors.HandleError(w, err)
		return
	}

	response := services.TransferResponse{
		Message:  "ok",
		Sender:   reqData.Username,
		Receiver: reqData.Receiver,
		Value:    strconv.Itoa(reqData.Value),
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func ClientAccountIDer(w http.ResponseWriter, r *http.Request) {
	var reqData services.AccountIDRequest
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to decode request payload")
		errors.HandleError(w, errors.NewAPIError(http.StatusBadRequest, "Invalid request payload"))
		return
	}

	clientID, err := services.ClientAccountID(reqData.Username)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to get client account ID")
		errors.HandleError(w, err)
		return
	}

	response := services.AccountIDResponse{
		Message: "ok",
		ID:      clientID,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func Initializer(w http.ResponseWriter, r *http.Request) {
	err := services.Initialize()
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to initialize service")
		errors.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Service initialized successfully"))
}