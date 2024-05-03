package repositories

import (
	"fmt"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"go-api/utils"
)

func SaveUser(username, password string) error {
	// Simulate saving user to the database
	fmt.Printf("Saving user: %s\n", username)
	return nil
}

func GetContract(username string) (*client.Contract, error) {
	certPath := fmt.Sprintf(utils.CryptoPath+"/users/%s@org1.example.com/msp/signcerts/cert.pem", username)
	keyPath := fmt.Sprintf(utils.CryptoPath+"/users/%s@org1.example.com/msp/keystore/", username)

	clientConnection := utils.NewGrpcConnection()

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
		return nil, err
	}

	network := gateway.GetNetwork(utils.ChannelName)
	contract := network.GetContract(utils.ChaincodeName)

	return contract, nil
}

func GetAdminContract() (*client.Contract, error) {
	certPath := utils.CryptoPath + "/users/User1@org1.example.com/msp/signcerts/cert.pem"
	keyPath := utils.CryptoPath + "/users/User1@org1.example.com/msp/keystore/"

	clientConnection := utils.NewGrpcConnection()

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
		return nil, err
	}

	network := gateway.GetNetwork(utils.ChannelName)
	contract := network.GetContract(utils.ChaincodeName)

	return contract, nil
}

// Add other repository functions...
