package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

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

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}

func (bcn *BlockChainNode) Run() {
	http.HandleFunc("/", HelloWorld)

	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcn.port)), nil))
}
