package repositories

import (
	"fmt"
	"net/http"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"go-api/errors"
	"go-api/logging"
	"go-api/utils"
	"go-api/config"
)

func SaveUser(username, password string) error {
	logging.Logger.Infof("Saving user: %s", username)
	// Simulate saving user to the database
	return nil
}

func GetContract(username string) (*client.Contract, error) {
	cfg, err := config.LoadConfig()
    if err != nil {
        logging.Logger.WithError(err).Error("Failed to load configuration")
        return nil, errors.NewAPIError(http.StatusInternalServerError, "Failed to get contract")
    }

    certPath := fmt.Sprintf(cfg.Fabric.CryptoPath+"/users/%s@org1.example.com/msp/signcerts/cert.pem", username)
    keyPath := fmt.Sprintf(cfg.Fabric.CryptoPath+"/users/%s@org1.example.com/msp/keystore/", username)

    clientConnection := utils.NewGrpcConnection()

    id := utils.NewIdentity(certPath)
    sign := utils.NewSign(keyPath)

    gateway, err := client.Connect(
        id,
        client.WithSign(sign),
        client.WithClientConnection(clientConnection),
        client.WithEvaluateTimeout(cfg.Timeouts.Evaluate),
        client.WithEndorseTimeout(cfg.Timeouts.Endorse),
        client.WithSubmitTimeout(cfg.Timeouts.Submit),
        client.WithCommitStatusTimeout(cfg.Timeouts.CommitStatus),
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
	cfg, err := config.LoadConfig()
    if err != nil {
        logging.Logger.WithError(err).Error("Failed to load configuration")
        return nil, errors.NewAPIError(http.StatusInternalServerError, "Failed to get admin contract")
    }

    certPath := cfg.Fabric.CryptoPath + "/users/Admin@org1.example.com/msp/signcerts/cert.pem"
    keyPath := cfg.Fabric.CryptoPath + "/users/Admin@org1.example.com/msp/keystore/"

    clientConnection := utils.NewGrpcConnection()

    id := utils.NewIdentity(certPath)
    sign := utils.NewSign(keyPath)

    gateway, err := client.Connect(
        id,
        client.WithSign(sign),
        client.WithClientConnection(clientConnection),
        client.WithEvaluateTimeout(cfg.Timeouts.Evaluate),
        client.WithEndorseTimeout(cfg.Timeouts.Endorse),
        client.WithSubmitTimeout(cfg.Timeouts.Submit),
        client.WithCommitStatusTimeout(cfg.Timeouts.CommitStatus),
    )
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to connect to gateway")
		return nil, errors.NewAPIError(http.StatusInternalServerError, "Failed to get admin contract")
	}

	network := gateway.GetNetwork(utils.ChannelName)
	contract := network.GetContract(utils.ChaincodeName)

	return contract, nil
}