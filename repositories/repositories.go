package repositories

import (
	"fmt"
	"net/http"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"go-api/errors"
	"go-api/logging"
	"go-api/utils"
)

func SaveUser(username, password string) error {
	logging.Logger.Infof("Saving user: %s", username)
	// Simulate saving user to the database
	return nil
}

func GetContract(username string) (*client.Contract, error) {
	certPath := fmt.Sprintf(utils.CryptoPath+"/users/%s@org1.example.com/msp/signcerts/cert.pem", username)
	keyPath := fmt.Sprintf(utils.CryptoPath+"/users/%s@org1.example.com/msp/keystore/", username)

	clientConnection := utils.NewGrpcConnection()
	defer clientConnection.Close()

	id := utils.NewIdentity(certPath)
	sign := utils.NewSign(keyPath)

	gateway, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		client.WithEvaluateTimeout(utils.EvaluateTimeout),
		client.WithEndorseTimeout(utils.EndorseTimeout),
		client.WithSubmitTimeout(utils.SubmitTimeout),
		client.WithCommitStatusTimeout(utils.CommitStatusTimeout),
	)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to connect to gateway")
		return nil, errors.NewAPIError(http.StatusInternalServerError, "Failed to get contract")
	}

	network := gateway.GetNetwork(utils.ChannelName)
	contract := network.GetContract(utils.ChaincodeName)

	return contract, nil
}

func GetAdminContract() (*client.Contract, error) {
	certPath := utils.CryptoPath + "/users/Admin@org1.example.com/msp/signcerts/cert.pem"
	keyPath := utils.CryptoPath + "/users/Admin@org1.example.com/msp/keystore/"

	clientConnection := utils.NewGrpcConnection()
	defer clientConnection.Close()

	id := utils.NewIdentity(certPath)
	sign := utils.NewSign(keyPath)

	gateway, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		client.WithEvaluateTimeout(utils.EvaluateTimeout),
		client.WithEndorseTimeout(utils.EndorseTimeout),
		client.WithSubmitTimeout(utils.SubmitTimeout),
		client.WithCommitStatusTimeout(utils.CommitStatusTimeout),
	)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to connect to gateway")
		return nil, errors.NewAPIError(http.StatusInternalServerError, "Failed to get admin contract")
	}

	network := gateway.GetNetwork(utils.ChannelName)
	contract := network.GetContract(utils.ChaincodeName)

	return contract, nil
}