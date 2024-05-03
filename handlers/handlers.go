package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-api/services"
)

func Register(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var reqData services.RegisterRequest
	json.NewDecoder(r.Body).Decode(&reqData)

	// Call the register service
	err := services.Register(reqData.Username, reqData.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Return the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func Minter(w http.ResponseWriter, r *http.Request) {
	var reqData services.MintRequest
	json.NewDecoder(r.Body).Decode(&reqData)

	err := services.Mint(reqData.Username, reqData.Value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
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
	json.NewDecoder(r.Body).Decode(&reqData)

	balance, err := services.ClientAccountBalance(reqData.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
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
	json.NewDecoder(r.Body).Decode(&reqData)

	err := services.Transfer(reqData.Username, reqData.Receiver, reqData.Value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
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
	json.NewDecoder(r.Body).Decode(&reqData)

	clientID, err := services.ClientAccountID(reqData.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
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
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Service initialized successfully"))
}

// Add other handler functions...