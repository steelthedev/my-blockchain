package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce         int
	previousHash  [32]byte
	timeStamp     int64
	transanctions []*Transaction
}

type BlockChain struct {
	transactionPool []*Transaction
	chain           []*Block
}

type Transaction struct {
	senderBlockChainAddress    string
	recepientBlockChainAddress string
	value                      float32
}

func NewTransaction(sender string, recepient string, value float32) *Transaction {
	return &Transaction{
		senderBlockChainAddress:    sender,
		recepientBlockChainAddress: recepient,
		value:                      value,
	}
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_blockchain_address\t%s\n", t.senderBlockChainAddress)
	fmt.Printf(" recepient_blockchain_address\t%s\n", t.recepientBlockChainAddress)
	fmt.Printf(" value\t\t\t%1f\n", t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recepient string  `json:"recepient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockChainAddress,
		Recepient: t.recepientBlockChainAddress,
		Value:     t.value,
	})
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Nonce         int            `json:"nonce"`
		PreviousHash  [32]byte       `json:"previousHash"`
		TimeStamp     int64          `json:"timestamp"`
		Transanctions []*Transaction `json:"transactions"`
	}{
		TimeStamp:     b.timeStamp,
		Nonce:         b.nonce,
		PreviousHash:  b.previousHash,
		Transanctions: b.transanctions,
	})
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	b := new(Block)
	b.timeStamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	b.transanctions = transactions

	return b
}

func NewBlockChain() *BlockChain {
	b := &Block{}
	bc := new(BlockChain)
	bc.CreateBlock(0, b.Hash())

	return bc
}

func (bc *BlockChain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
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
	fmt.Printf("previousHash\t%x\n", b.previousHash)
	for _, transaction := range b.transanctions {
		transaction.Print()
	}
}

func (bc *BlockChain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func init() {
	log.SetPrefix("Blockchain Node. ")
}

func main() {
	blockChain := NewBlockChain()
	blockChain.Print()
	blockChain.CreateBlock(23, blockChain.LastBlock().previousHash)
	blockChain.Print()
	blockChain.CreateBlock(42, blockChain.LastBlock().previousHash)
	blockChain.Print()

}
