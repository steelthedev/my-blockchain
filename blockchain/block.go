package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"time"
)

type Block struct {
	timeStamp     int64
	nonce         int
	previousHash  [32]byte
	transanctions []*Transaction
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
