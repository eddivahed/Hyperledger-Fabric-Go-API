```markdown
# API Documentation

This document provides detailed information about the available endpoints, request/response formats, and authentication requirements for the Hyperledger Fabric ERC20 Token API.

## Authentication

The API endpoints are protected with JWT-based authentication. To access the protected endpoints, you need to include a valid JWT token in the `Authorization` header of your requests.

To obtain a JWT token, send a POST request to the `/login` endpoint with the following payload:
```json
{
  "username": "admin",
  "password": "password"
}
```

The response will include a JWT token that you can use for subsequent requests.

## Endpoints

### Register User

- URL: `/register`
- Method: POST
- Description: Registers a new user on the network.
- Request Body:
  ```json
  {
    "username": "consumer0001",
    "password": "consumer0001pw"
  }
  ```
- Response:
  ```json
  {
    "message": "User registered successfully"
  }
  ```

The `/register` endpoint allows you to register a new user on the Hyperledger Fabric network. By providing a unique username and password, you can create a new user account that can interact with the ERC20 token chaincode.

### Mint Tokens

- URL: `/mint`
- Method: POST
- Description: Mints new tokens for a registered user.
- Request Body:
  ```json
  {
    "username": "consumer0001",
    "value": 100
  }
  ```
- Response:
  ```json
  {
    "message": "ok",
    "username": "consumer0001",
    "value": "100"
  }
  ```

The `/mint` endpoint enables you to mint new tokens for a registered user. By specifying the username and the amount of tokens to mint, you can increase the token balance of the user.

### Transfer Tokens

- URL: `/transfer`
- Method: POST
- Description: Transfers tokens from one registered user to another.
- Request Body:
  ```json
  {
    "username": "consumer0001",
    "receiver": "eDUwOTo6Q049Y29uc3VtZXIwMDAyLE9VPWNsaWVudCxPPUh5cGVybGVkZ2VyLFNUPU5vcnRoIENhcm9saW5hLEM9VVM6OkNOPWNhLm9yZzEuZXhhbXBsZS5jb20sTz1vcmcxLmV4YW1wbGUuY29tLEw9RHVyaGFtLFNUPU5vcnRoIENhcm9saW5hLEM9VVM=",
    "value": 50
  }
  ```
- Response:
  ```json
  {
    "message": "ok",
    "sender": "consumer0001",
    "receiver": "eDUwOTo6Q049Y29uc3VtZXIwMDAyLE9VPWNsaWVudCxPPUh5cGVybGVkZ2VyLFNUPU5vcnRoIENhcm9saW5hLEM9VVM6OkNOPWNhLm9yZzEuZXhhbXBsZS5jb20sTz1vcmcxLmV4YW1wbGUuY29tLEw9RHVyaGFtLFNUPU5vcnRoIENhcm9saW5hLEM9VVM=",
    "value": "50"
  }
  ```

The `/transfer` endpoint allows you to transfer tokens from one registered user to another. By providing the username of the sender, the unique account ID (wallet ID) of the receiver, and the amount of tokens to transfer, you can move tokens between user accounts.

### Get Balance

- URL: `/balance`
- Method: POST
- Description: Retrieves the token balance of a registered user.
- Request Body:
  ```json
  {
    "username": "consumer0001"
  }
  ```
- Response:
  ```json
  {
    "message": "ok",
    "value": "100"
  }
  ```

The `/balance` endpoint enables you to retrieve the current token balance of a registered user. By providing the username, you can query the balance associated with that user's account.

### Get Client Account ID

- URL: `/accountid`
- Method: POST
- Description: Retrieves the unique account ID (wallet ID) associated with a registered user.
- Request Body:
  ```json
  {
    "username": "consumer0001"
  }
  ```
- Response:
  ```json
  {
    "message": "ok",
    "id": "eDUwOTo6Q049Y29uc3VtZXIwMDAxLE9VPWNsaWVudCxPPUh5cGVybGVkZ2VyLFNUPU5vcnRoIENhcm9saW5hLEM9VVM6OkNOPWNhLm9yZzEuZXhhbXBsZS5jb20sTz1vcmcxLmV4YW1wbGUuY29tLEw9RHVyaGFtLFNUPU5vcnRoIENhcm9saW5hLEM9VVM="
  }
  ```

The `/accountid` endpoint allows you to retrieve the unique account ID (wallet ID) associated with a registered user. By providing the username, you can obtain the account ID, which is necessary when transferring tokens to another user.

## Error Handling

In case of any errors, the API will respond with an appropriate HTTP status code and an error message in the response body. Common error scenarios include:

- 400 Bad Request: The request payload is invalid or missing required fields.
- 401 Unauthorized: The provided JWT token is missing, invalid, or expired.
- 403 Forbidden: The user does not have sufficient permissions to perform the requested action.
- 500 Internal Server Error: An unexpected error occurred on the server side.

## Postman Collection

For your convenience, a Postman collection is provided in the `Blockchain API V0001.postman_collection.json` file. You can import this collection into Postman to easily test and interact with the API endpoints.

Make sure to update the `{{token}}` variable in the Postman collection with a valid JWT token obtained from the `/login` endpoint.
