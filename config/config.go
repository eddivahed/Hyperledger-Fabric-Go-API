package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server         ServerConfig
	Fabric         FabricConfig
	Timeouts       TimeoutsConfig
	RegisterScript string
	JWT JWTConfig
}

type ServerConfig struct {
	Port int
}

type JWTConfig struct {
	SecretKey string
}

type FabricConfig struct {
	MSPID         string
	CryptoPath    string
	TLSCertPath   string
	PeerEndpoint  string
	GatewayPeer   string
	ChannelName   string
	ChaincodeName string
}

type TimeoutsConfig struct {
	Evaluate     time.Duration
	Endorse      time.Duration
	Submit       time.Duration
	CommitStatus time.Duration
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml")

	viper.SetDefault("server.port", 8082)
	viper.SetDefault("fabric.mspid", "Org1MSP")
	viper.SetDefault("fabric.cryptopath", "../organizations/peerOrganizations/org1.example.com")
	viper.SetDefault("fabric.tlscertpath", "/peers/peer0.org1.example.com/tls/ca.crt")
	viper.SetDefault("fabric.peerendpoint", "localhost:7051")
	viper.SetDefault("fabric.gatewaypeer", "peer0.org1.example.com")
	viper.SetDefault("fabric.channelname", "mychannel")
	viper.SetDefault("fabric.chaincodename", "token_erc20")
	viper.SetDefault("timeouts.evaluate", 5*time.Second)
	viper.SetDefault("timeouts.endorse", 15*time.Second)
	viper.SetDefault("timeouts.submit", 5*time.Second)
	viper.SetDefault("timeouts.commitstatus", 1*time.Minute)
	viper.SetDefault("registerScript", "./script.sh")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}