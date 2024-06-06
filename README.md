# Simple Blockchain in Golang

## Project Description
This project involves creating a simple blockchain implementation using the Go programming language. The blockchain will support basic functionalities such as adding new blocks, listing all blocks, changing block data, and verifying the integrity of the blockchain.

## Features
- **Add New Block**: Create and add new blocks to the blockchain with a specified transaction, nonce, and previous hash.
- **List Blocks**: Print all blocks in a formatted manner displaying transaction data, nonce, previous hash, and current block hash.
- **Change Block**: Modify the transaction data of a specific block.
- **Verify Chain**: Verify the integrity of the blockchain to ensure no tampering has occurred.
- **Calculate Hash**: Calculate the hash of a block.

## Functions
The project includes the following public functions:

### `NewBlock(transaction string, nonce int, previousHash string) *block`
A method to add a new block. You can use any string as a transaction (e.g., "bob to alice") and any integer value as a nonce. The `CreateHash()` method will provide the block hash value.

### `ListBlocks()`
A method to print all the blocks in a nice format showing block data such as transaction, nonce, previous hash, and current block hash.

### `ChangeBlock()`
A function to change the transaction of a given block reference.

### `VerifyChain()`
A function to verify the blockchain in case any changes are made. This can be done in two different ways.

### `CalculateHash(stringToHash string)`
A function for calculating the hash of a block.

## Repository Setup
1. Clone this repository:
    ```sh
    git clone https://github.com/ameerahaider/Simple-Blockchain-in-Golang.git
    ```
2. Navigate to the project directory:
    ```sh
    cd assignment01bca
    ```
3. Ensure your Go environment is set up:
    ```sh
    go env
    ```
4. Run the project:
    ```sh
    go run main.go
    ```
