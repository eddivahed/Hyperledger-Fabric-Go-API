package utils

import (
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"path"
	"time"
	"net/http"


	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"go-api/config"
	"go-api/errors"
	"go-api/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	MSPID               = "Org1MSP"
	CryptoPath          = "../organizations/peerOrganizations/org1.example.com"
	TLSCertPath         = CryptoPath + "/peers/peer0.org1.example.com/tls/ca.crt"
	PeerEndpoint        = "localhost:7051"
	GatewayPeer         = "peer0.org1.example.com"
	ChannelName         = "mychannel"
	ChaincodeName       = "token_erc20"
	EvaluateTimeout     = 5 * time.Second
	EndorseTimeout      = 15 * time.Second
	SubmitTimeout       = 5 * time.Second
	CommitStatusTimeout = 1 * time.Minute
)

func NewGrpcConnection() *grpc.ClientConn {
	cfg, err := config.LoadConfig()
	if err != nil {
		logging.Logger.WithError(err).Fatal("Failed to load configuration")
	}

	certificate, err := LoadCertificate(cfg.Fabric.CryptoPath + cfg.Fabric.TLSCertPath)
	if err != nil {
		logging.Logger.WithError(err).Fatal("Failed to load TLS certificate")
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, cfg.Fabric.GatewayPeer)

	connection, err := grpc.Dial(cfg.Fabric.PeerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		logging.Logger.WithError(err).Fatal("Failed to create gRPC connection")
	}

	return connection
}

func NewIdentity(certPath string) *identity.X509Identity {
	cfg, err := config.LoadConfig()
	if err != nil {
		logging.Logger.WithError(err).Fatal("Failed to load configuration")
	}

	certificate, err := LoadCertificate(certPath)
	if err != nil {
		logging.Logger.WithError(err).Fatal("Failed to load identity certificate")
	}

	id, err := identity.NewX509Identity(cfg.Fabric.MSPID, certificate)
	if err != nil {
		logging.Logger.WithError(err).Fatal("Failed to create X509 identity")
	}

	return id
}

func LoadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.NewAPIError(http.StatusInternalServerError, fmt.Sprintf("Failed to read certificate file: %v", err))
	}
	return identity.CertificateFromPEM(certificatePEM)
}

func NewSign(keyPath string) identity.Sign {
	files, err := ioutil.ReadDir(keyPath)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to read private key directory")
		panic(err)
	}
	privateKeyPEM, err := ioutil.ReadFile(path.Join(keyPath, files[0].Name()))
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to read private key file")
		panic(err)
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to parse private key")
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to create signer")
		panic(err)
	}

	return sign
}
