
# Hyperledger Fabric Asset Management System

This repository contains the code for an asset management system for a financial institution, implemented with Hyperledger Fabric and Golang. The system allows secure, transparent, and immutable management and tracking of assets through smart contracts, enabling asset creation, updates, queries, and transaction history retrieval.

## Project Structure

- **Level 1:** Setup of the Hyperledger Fabric test network.
- **Level 2:** Development and testing of smart contracts to support asset management operations.
- **Level 3:** REST API development for invoking the smart contract, deployed on Hyperledger Fabric, and creating a Docker image for the REST API.

## Features

- **Create Asset:** Define a new asset with specific attributes (e.g., `DEALERID`, `MSISDN`, `MPIN`, etc.).
- **Update Asset:** Modify asset attributes.
- **Query World State:** Read asset details from the blockchain.
- **Transaction History:** Retrieve transaction history for assets.

## Prerequisites

- Golang
- Docker
- Hyperledger Fabric
- [Hyperledger Fabric Samples](https://github.com/hyperledger/fabric-samples)

## Getting Started

### 1. Hyperledger Fabric Network Setup

Set up the test network by following the official Hyperledger Fabric documentation:

```shell
cd fabric-samples/test-network
./network.sh up createChannel
```

### 2. Smart Contract Deployment

To develop and deploy the smart contract:

1. Write the chaincode in Go (Golang) in the `asset-transfer-basic/chaincode-go` directory.
2. Deploy the chaincode:

   ```shell
   ./network.sh deployCC -ccn asset-transfer -ccp ../asset-transfer-basic/chaincode-go -ccl go
   ```

3. Test the chaincode functionality to ensure assets are created, updated, and queried as expected.

### 3. REST API Development

Implement REST API endpoints to interact with the smart contract, using the Hyperledger Fabric gateway.

```shell
# Example: Start the REST API server
go run main.go
```

#### Endpoints:

- **POST /assets** - Create a new asset.
- **PUT /assets/{id}** - Update an asset.
- **GET /assets/{id}** - Retrieve an asset’s details.
- **GET /assets/{id}/history** - Get transaction history for an asset.

### 4. Docker Setup for REST API

Build and run the Docker container for the REST API:

```shell
docker build -t asset-management-api .
docker run -p 8080:8080 asset-management-api
```

## References

- [Hyperledger Fabric Documentation](https://hyperledger-fabric.readthedocs.io/)
- [Fabric Samples Repository](https://github.com/hyperledger/fabric-samples)

## Author

- **Tejasree**


## output
**chaincode invoke successfully**

![Screenshot (1636)](https://github.com/user-attachments/assets/43397d58-f8c8-42b9-8b01-dbfe02805b41)


command: go run.

![Screenshot (1638)](https://github.com/user-attachments/assets/7c888852-3420-4e62-a88e-050a73126cb6)

![Screenshot (1640)](https://github.com/user-attachments/assets/d8d2be7c-8e42-412b-8434-87c0afabfb54)
![Screenshot (1641)](https://github.com/user-attachments/assets/85f3281e-1816-4834-a697-980e514142b1)



