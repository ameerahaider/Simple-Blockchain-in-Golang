package assignment01bca

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

type Block struct {
	transaction string
	nonce       int
	prevHash    string
	blockHash   string
}

type Blockchain struct {
	list []*Block
}

func CalculateHash(b *Block) string {
	stringToHash := strconv.Itoa(b.nonce) + b.transaction + b.prevHash
	//fmt.Printf("String Received: %s\n", stringToHash)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))
}

func (bc *Blockchain) NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := new(Block)
	block.transaction = transaction
	block.nonce = nonce
	block.prevHash = previousHash
	block.blockHash = CalculateHash(block)
	bc.list = append(bc.list, block)
	return block
}

func PrintBlock(b *Block) {
	fmt.Printf("%s Block %s\n", strings.Repeat("=", 25), strings.Repeat("=", 25))
	fmt.Printf("Transaction: %s\n", b.transaction)
	fmt.Printf("Nonce: %d\n", b.nonce)
	fmt.Printf("Previous Hash: %s\n", b.prevHash)
	fmt.Printf("Block Hash: %s\n", b.blockHash)
}

func ListBlocks(bc *Blockchain) {
	println()
	for i := 0; i < len(bc.list); i++ {
		b := bc.list[i]
		PrintBlock(b)
	}
}

func ChangeBlock(b *Block, newTransaction string) {
	//function to change block transaction of the given block ref
	b.transaction = newTransaction
	b.blockHash = CalculateHash(b)
}

func VerifyChain(bc *Blockchain, LastHash string) {
	// Function to verify the integrity of the blockchain by checking the hashes.
	for i := len(bc.list) - 1; i > 0; i-- {
		b := bc.list[i-1]

		if CalculateHash(b) == LastHash {
			//fmt.Printf("Chain Valid at Block %d\n", i)
			LastHash = b.prevHash
			continue
		} else {
			fmt.Printf("\nChain Invalid at Block %d\n", i)
			return
		}
	}

	fmt.Println("\nChain Verified: All blocks are connected and valid.")
}

func main() {

	//Make Blockchain
	blockchain := new(Blockchain)
	b1 := blockchain.NewBlock("Transaction 1", 1, "")
	b2 := blockchain.NewBlock("Transaction 2", 2, CalculateHash(b1))
	b3 := blockchain.NewBlock("Transaction 3", 3, CalculateHash(b2))
	b4 := blockchain.NewBlock("Transaction 4", 4, CalculateHash(b3))
	b5 := blockchain.NewBlock("Transaction 5", 5, CalculateHash(b4))
	b6 := blockchain.NewBlock("Transaction 6", 6, CalculateHash(b5))
	b7 := blockchain.NewBlock("", 7, CalculateHash(b6))
	LastHash := b7.prevHash //Store Last Hash Only

	//Print Chain
	ListBlocks(blockchain)

	//Verify Chain
	VerifyChain(blockchain, LastHash)

	//Tamper Chain
	ChangeBlock(b2, "Transaction 2 Updated")

	//Print Tampered Chain
	ListBlocks(blockchain)

	//Verify Tampered Chain
	VerifyChain(blockchain, LastHash)
}
