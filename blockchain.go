package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce         int
	previousHash  string
	timeStamp     int64
	transanctions []string
}

type BlockChain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.timeStamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash

	return b
}

func NewBlockChain() *BlockChain {
	bc := new(BlockChain)

	bc.CreateBlock(0, "hash #0 genesis block")

	return bc
}

func (bc *BlockChain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)

	return b
}

func (bc *BlockChain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Block %d %s\n", strings.Repeat("=", 10), i, strings.Repeat("=", 10))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("#", 27))
}

func (b *Block) Print() {
	fmt.Printf("timestamp\t%d\n", b.timeStamp)
	fmt.Printf("nonce\t\t%d\n", b.nonce)
	fmt.Printf("previousHash\t%s\n", b.previousHash)
	fmt.Printf("transactions\t%s\n", b.transanctions)
}

func init() {
	log.SetPrefix("Blockchain Node. ")
}

func main() {
	blockChain := NewBlockChain()
	blockChain.Print()
	blockChain.CreateBlock(23, "Hash #1")
	blockChain.Print()
	blockChain.CreateBlock(42, "Hash #2")
	blockChain.Print()
}
