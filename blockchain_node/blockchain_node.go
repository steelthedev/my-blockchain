package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.con/steelthedev/my-blockchain/blockchain"
	"github.con/steelthedev/my-blockchain/wallet"
)

var cache map[string]*blockchain.BlockChain = make(map[string]*blockchain.BlockChain)

type BlockChainNode struct {
	port uint16
}

func NewBlockChainNode(port uint16) *BlockChainNode {
	return &BlockChainNode{
		port,
	}
}

func (bcn *BlockChainNode) Port() uint16 {
	return bcn.port
}

func (bcn *BlockChainNode) GetBlockChain() *blockchain.BlockChain {
	bc, ok := cache["blockchain"]
	if !ok {
		minerWallet := wallet.NewWallet()
		bc = blockchain.NewBlockChain(minerWallet.BlockChainAddress(), bcn.Port())
		cache["blockchain"] = bc
	}
	return bc
}

func (bcn *BlockChainNode) GetChain(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	bc := bcn.GetBlockChain()
	m, _ := bc.MarshalJSON()
	io.WriteString(w, string(m[:]))
}

func (bcn *BlockChainNode) Run() {
	http.HandleFunc("GET /", bcn.GetChain)

	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcn.port)), nil))
}
