package main

import (
	"fmt"
	"log"

	"github.con/steelthedev/my-blockchain/wallet"
)

func init() {
	log.SetPrefix("Blockchain Node. ")
}

func main() {
	w := wallet.NewWallet()
	t := wallet.NewTransaction(w.PrivateKey(), w.PublicKey(), w.BlockChainAddress(), "Rick", 120)

	fmt.Printf("Signature %s", t.GenerateSignature())
}
