package utils

import (
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"path"
	"time"

	"github.com/hyperledger/fabric-gateway/pkg/identity"
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
	certificate, err := LoadCertificate(TLSCertPath)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, GatewayPeer)

	connection, err := grpc.Dial(PeerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %v", err))
	}

	return connection
}

func NewIdentity(certPath string) *identity.X509Identity {
	certificate, err := LoadCertificate(certPath)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(MSPID, certificate)
	if err != nil {
		panic(err)
	}

	return id
}

func LoadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %v", err)
	}
	return identity.CertificateFromPEM(certificatePEM)
}

func NewSign(keyPath string) identity.Sign {
	files, err := ioutil.ReadDir(keyPath)
	if err != nil {
		panic(fmt.Errorf("failed to read private key directory: %v", err))
	}
	privateKeyPEM, err := ioutil.ReadFile(path.Join(keyPath, files[0].Name()))

	if err != nil {
		panic(fmt.Errorf("failed to read private key file: %v", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
} 
