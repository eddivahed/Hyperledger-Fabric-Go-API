package services

import (
	"fmt"
	"os/exec"
	"strconv"


	"go-api/repositories"

)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(username, password string) error {
	// Execute the script.sh script
	cmd := "./script.sh"
	out, err := exec.Command(cmd, username, password).Output()
	if err != nil {
		return fmt.Errorf("failed to execute register script: %v", err)
	}

	// Save the user in the repository
	err = repositories.SaveUser(username, password)
	if err != nil {
		return fmt.Errorf("failed to save user: %v", err)
	}

	fmt.Printf("User registered: %s\n", string(out))
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
		return err
	}

	_, err = contract.SubmitTransaction("Mint", strconv.Itoa(value))
	if err != nil {
		return fmt.Errorf("failed to submit transaction: %v", err)
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
		return "", err
	}

	evaluateResult, err := contract.EvaluateTransaction("ClientAccountBalance")
	if err != nil {
		return "", fmt.Errorf("failed to evaluate transaction: %v", err)
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
		return err
	}

	_, err = contract.SubmitTransaction("Transfer", receiver, strconv.Itoa(value))
	if err != nil {
		return fmt.Errorf("failed to submit transaction: %v", err)
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
		return "", err
	}

	evaluateResult, err := contract.EvaluateTransaction("ClientAccountID")
	if err != nil {
		return "", fmt.Errorf("failed to evaluate transaction: %v", err)
	}

	return string(evaluateResult), nil
}

func Initialize() error {
	contract, err := repositories.GetAdminContract()
	if err != nil {
		return err
	}

	_, err = contract.SubmitTransaction("Initialize", "energycoin", "ec", "2")
	if err != nil {
		return fmt.Errorf("failed to submit transaction: %v", err)
	}

	return nil
}
