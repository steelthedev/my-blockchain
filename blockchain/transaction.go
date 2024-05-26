package blockchain

import (
	"encoding/json"
	"fmt"
	"strings"
)

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
