package main

import (
	"flag"
	"log"
)

func Init() {
	log.SetPrefix("Wallet Server:")
}

func main() {
	port := flag.Uint("port", 8080, "TCP Port Number for Online Wallet")
	gateway := flag.String("gateway", "http://127.0.0.0.1:3332", "BlockChain gateway")
	flag.Parse()

	app := NewWalletServer(uint16(*port), *gateway)
	log.Println("Starting Wallet on port:", *port, "using blockchain node", *gateway, "as gateway")
	app.Run()
}
