package main

import (
	"github.com/ameerahaider/assignment01bca"
)

func main() {
	// Create a new instance of the Blockchain from the package
	blockchain := assignment01bca.NewBlockchain()

	// Create blocks and add them to the blockchain
	b1 := blockchain.NewBlock("Transaction 1", 1, "")
	b2 := blockchain.NewBlock("Transaction 2", 2, assignment01bca.CalculateHash(b1))
	b3 := blockchain.NewBlock("Transaction 3", 3, assignment01bca.CalculateHash(b2))
	b4 := blockchain.NewBlock("Transaction 4", 4, assignment01bca.CalculateHash(b3))
	b5 := blockchain.NewBlock("Transaction 5", 5, assignment01bca.CalculateHash(b4))
	b6 := blockchain.NewBlock("Transaction 6", 6, assignment01bca.CalculateHash(b5))
	b7 := blockchain.NewBlock("", 7, assignment01bca.CalculateHash(b6))
	LastHash := b7.PrevHash // Store the last hash only

	// Print the chain
	assignment01bca.ListBlocks(blockchain)

	// Verify the chain
	assignment01bca.VerifyChain(blockchain, LastHash)

	// Tamper with a block
	assignment01bca.ChangeBlock(b2, "Transaction 2 Updated")

	// Print the tampered chain
	assignment01bca.ListBlocks(blockchain)

	// Verify the tampered chain
	assignment01bca.VerifyChain(blockchain, LastHash)
}
