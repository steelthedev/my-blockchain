package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

const (
	MINING_DIFFICULTY = 3
	MINING_SENDER     = "BLOCKCHAIN REWARD SYSTEM"
	MINING_REWARD     = 1.0
)

type Block struct {
	timeStamp     int64
	nonce         int
	previousHash  [32]byte
	transanctions []*Transaction
}

type BlockChain struct {
	transactionPool   []*Transaction
	chain             []*Block
	blockChainAddress string
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
	fmt.Printf(" value\t\t\t\t%1f\n", t.value)
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

func NewBlockChain(blockChainAddress string) *BlockChain {
	b := &Block{}
	bc := new(BlockChain)
	bc.blockChainAddress = blockChainAddress
	bc.CreateBlock(0, b.Hash())

	return bc
}

func (bc *BlockChain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
	return b
}

func (bc *BlockChain) AddTransaction(sender string, recipient string, value float32) error {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
	return nil
}

func (bc *BlockChain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, transaction := range bc.transactionPool {
		transactions = append(transactions, NewTransaction(transaction.senderBlockChainAddress,
			transaction.recepientBlockChainAddress,
			transaction.value))
	}

	return transactions
}

func (bc *BlockChain) ValidProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{0, nonce, previousHash, transactions}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeros
}

func (bc *BlockChain) ProofOfWork() int {
	transactions := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()
	nonce := 0

	for !bc.ValidProof(nonce, previousHash, transactions, MINING_DIFFICULTY) {
		nonce++
	}
	return nonce
}

func (bc *BlockChain) Mining() bool {
	bc.AddTransaction(MINING_SENDER, bc.blockChainAddress, MINING_REWARD)
	nonce := bc.ProofOfWork()
	previousHash := bc.LastBlock().Hash()
	bc.CreateBlock(nonce, previousHash)
	log.Println("action=mining, status=success")
	return true
}

func (bc *BlockChain) CalculateTotalAmount(blockChainAddress string) float32 {
	var totalAmount float32 = 0
	for _, b := range bc.chain {
		for _, t := range b.transanctions {
			value := t.value
			if blockChainAddress == t.recepientBlockChainAddress {
				totalAmount += value
			}

			if blockChainAddress == t.senderBlockChainAddress {
				totalAmount -= value
			}
		}
	}
	return totalAmount
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
	minerBlockChainAddress := "Miner's blockchain address"
	blockChain := NewBlockChain(minerBlockChainAddress)

	blockChain.AddTransaction("Rick", "Morty", 127)
	blockChain.Mining()

	blockChain.AddTransaction("Rick", "Taju", 300)
	blockChain.AddTransaction("Taju", "Seun", 350)
	blockChain.AddTransaction("Taju", "Rick", 500)
	blockChain.Mining()

	blockChain.Print()

	fmt.Printf("Miner %.1f\n", blockChain.CalculateTotalAmount(minerBlockChainAddress))

	fmt.Printf("Rick %.1f\n", blockChain.CalculateTotalAmount("Rick"))

	fmt.Printf("Morty %.1f\n", blockChain.CalculateTotalAmount("Morty"))

}
