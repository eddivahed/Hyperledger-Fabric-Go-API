# Hyperledger Fabric ERC20 Token API

This project provides a Go-based API for interacting with the ERC20 token chaincode sample in Hyperledger Fabric. It allows users to register, mint tokens, transfer tokens, check balances, and retrieve account IDs.

## Prerequisites

Before using the Go API, ensure that you have the following prerequisites installed:

- Go programming language (version 1.16 or later)
- Docker and Docker Compose
- Hyperledger Fabric binaries and Docker images

## Getting Started

Follow these steps to set up and run the Go API:

1. Clone the Hyperledger Fabric samples repository:
git clone https://github.com/hyperledger/fabric-samples.git

2. Navigate to the `fabric-samples/test-network` directory:
cd fabric-samples/test-network

3. Start the Hyperledger Fabric test network:
./network.sh up createChannel -ca

4. Deploy the ERC20 token chaincode:
./network.sh deployCC -ccn token_erc20 -ccp ../token-erc20/chaincode-go -ccl go

5. Set up the initial contract configuration:
- Open the `config.yaml` file in the `fabric-samples/token-erc20/chaincode-go` directory.
- Modify the configuration values as needed (e.g., token symbol, initial supply).

6. Clone this Go API project repository into the `fabric-samples/test-network` directory:
git colne https://github.com/eddivahed/Hyperledger-Fabric-Go-API.git

7. Navigate to the Go API project directory
8. Install the project dependencies:
go mod download
go mod tidy

9. Start the Go API server:
go run main.go

The Go API server should now be running and accessible at `http://localhost:8082`.

## Interacting with the API

Once the Go API server is running, you can interact with it using HTTP requests. Refer to the [API Documentation](api.md) for detailed information on the available endpoints, request/response formats, and authentication requirements.

You can use tools like cURL or Postman to send requests to the API endpoints.

## Troubleshooting

- If you encounter any issues related to the Hyperledger Fabric test network setup or chaincode installation, refer to the Hyperledger Fabric documentation for troubleshooting steps.
- Make sure that the Go API project is placed in the correct directory (`fabric-samples/test-network`) to have proper access to the network cryptographic materials.
- Verify that the `config.yaml` file in the `fabric-samples/token-erc20/chaincode-go` directory is properly configured with the desired initial contract settings.

## License

This project is licensed under the [MIT License](LICENSE).
