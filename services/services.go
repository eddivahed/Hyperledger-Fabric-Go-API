package services

import (
	"net/http"
	"os/exec"
	"strconv"

	"go-api/errors"
	"go-api/logging"
	"go-api/config"
	"go-api/repositories"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(username, password string) error {
	cfg, err := config.LoadConfig()
    if err != nil {
        logging.Logger.WithError(err).Error("Failed to load configuration")
        return errors.NewAPIError(http.StatusInternalServerError, "Failed to register user")
    }

    cmd := cfg.RegisterScript
	out, err := exec.Command(cmd, username, password).Output()
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to execute register script")
		return errors.NewAPIError(http.StatusInternalServerError, "Failed to register user")
	}

	err = repositories.SaveUser(username, password)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to save user")
		return errors.NewAPIError(http.StatusInternalServerError, "Failed to register user")
	}

	logging.Logger.Infof("User registered: %s", string(out))
	return nil
}

type MintRequest struct {
	Username string `json:"username"`
	Value    int    `json:"value"`
}

type MintResponse struct {
	Message  string `json:"message"`
	Username string `json:"username"`
	Value    string `json:"value"`
}

func Mint(username string, value int) error {
	contract, err := repositories.GetContract(username)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to get contract")
		return errors.NewAPIError(http.StatusInternalServerError, "Failed to mint tokens")
	}

	_, err = contract.SubmitTransaction("Mint", strconv.Itoa(value))
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to submit transaction")
		return errors.NewAPIError(http.StatusInternalServerError, "Failed to mint tokens")
	}

	return nil
}

type BalanceRequest struct {
	Username string `json:"username"`
}

type BalanceResponse struct {
	Message string `json:"message"`
	Value   string `json:"value"`
}

func ClientAccountBalance(username string) (string, error) {
	contract, err := repositories.GetContract(username)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to get contract")
		return "", errors.NewAPIError(http.StatusInternalServerError, "Failed to get account balance")
	}

	evaluateResult, err := contract.EvaluateTransaction("ClientAccountBalance")
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to evaluate transaction")
		return "", errors.NewAPIError(http.StatusInternalServerError, "Failed to get account balance")
	}

	return string(evaluateResult), nil
}

type TransferRequest struct {
	Username string `json:"username"`
	Receiver string `json:"receiver"`
	Value    int    `json:"value"`
}

type TransferResponse struct {
	Message  string `json:"message"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Value    string `json:"value"`
}

func Transfer(username, receiver string, value int) error {
	contract, err := repositories.GetContract(username)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to get contract")
		return errors.NewAPIError(http.StatusInternalServerError, "Failed to transfer tokens")
	}

	_, err = contract.SubmitTransaction("Transfer", receiver, strconv.Itoa(value))
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to submit transaction")
		return errors.NewAPIError(http.StatusInternalServerError, "Failed to transfer tokens")
	}

	return nil
}

type AccountIDRequest struct {
	Username string `json:"username"`
}

type AccountIDResponse struct {
	Message string `json:"message"`
	ID      string `json:"id"`
}

func ClientAccountID(username string) (string, error) {
	contract, err := repositories.GetContract(username)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to get contract")
		return "", errors.NewAPIError(http.StatusInternalServerError, "Failed to get client account ID")
	}

	evaluateResult, err := contract.EvaluateTransaction("ClientAccountID")
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to evaluate transaction")
		return "", errors.NewAPIError(http.StatusInternalServerError, "Failed to get client account ID")
	}

	return string(evaluateResult), nil
}

func Initialize() error {
	contract, err := repositories.GetAdminContract()
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to get admin contract")
		return errors.NewAPIError(http.StatusInternalServerError, "Failed to initialize service")
	}

	_, err = contract.SubmitTransaction("Initialize", "energycoin", "ec", "2")
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to submit transaction")
		return errors.NewAPIError(http.StatusInternalServerError, "Failed to initialize service")
	}

	return nil
}
 