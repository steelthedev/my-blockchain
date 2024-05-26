package main

import (
	"flag"
	"log"
)

func init() {
	log.SetPrefix("Blockchain Node. ")
}

func main() {
	port := flag.Uint("port", 3333, "TCP port number for Blockchain Node")
	flag.Parse()

	app := NewBlockChainNode(uint16(*port))
	log.Default().Println("starting node on port", *port)

	app.Run()
}
