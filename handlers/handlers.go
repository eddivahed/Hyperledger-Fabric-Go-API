package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"go-api/config"
	"go-api/logging"
	"go-api/services"

	"github.com/dgrijalva/jwt-go"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the API!"))
}

func Register(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value("role").(string)
	if !HasPermission(role, "register") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var reqData services.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to decode request payload")
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = services.Register(reqData.Username, reqData.Password)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to register user")
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func Minter(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value("role").(string)
	if !HasPermission(role, "mint") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var reqData services.MintRequest
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to decode request payload")
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = services.Mint(reqData.Username, reqData.Value)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to mint tokens")
		http.Error(w, "Failed to mint tokens", http.StatusInternalServerError)
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
	role := r.Context().Value("role").(string)
	if !HasPermission(role, "balance") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var reqData services.BalanceRequest
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to decode request payload")
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	balance, err := services.ClientAccountBalance(reqData.Username)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to get account balance")
		http.Error(w, "Failed to get account balance", http.StatusInternalServerError)
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
	role := r.Context().Value("role").(string)
	if !HasPermission(role, "transfer") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var reqData services.TransferRequest
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to decode request payload")
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = services.Transfer(reqData.Username, reqData.Receiver, reqData.Value)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to transfer tokens")
		http.Error(w, "Failed to transfer tokens", http.StatusInternalServerError)
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
	role := r.Context().Value("role").(string)
	if !HasPermission(role, "accountID") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var reqData services.AccountIDRequest
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to decode request payload")
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	clientID, err := services.ClientAccountID(reqData.Username)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to get client account ID")
		http.Error(w, "Failed to get client account ID", http.StatusInternalServerError)
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
	role := r.Context().Value("role").(string)
	if !HasPermission(role, "initialize") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err := services.Initialize()
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to initialize service")
		http.Error(w, "Failed to initialize service", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Service initialized successfully"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Verify the user credentials (replace with your own authentication logic)
	if credentials.Username != "admin" || credentials.Password != "password" {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate a JWT token
	token, err := GenerateToken(credentials.Username, RoleAdmin)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	response := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	json.NewEncoder(w).Encode(response)
}

func GenerateToken(userID, role string) (string, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to load configuration")
		return "", err
	}

	claims := jwt.MapClaims{
		"userID": userID,
		"role":   role,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.JWT.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
