package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

const pathToTemplateDir = "templates"

type WalletServer struct {
	port    uint16
	gateway string
}

func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{port, gateway}
}

func (ws *WalletServer) Port() uint16 {
	return ws.port
}

func (ws *WalletServer) Gateway() string {
	return ws.gateway
}

func (ws *WalletServer) Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(path.Join(pathToTemplateDir, "index.html"))
	t.Execute(w, "")
}

func (ws *WalletServer) Run() {
	http.HandleFunc("GET /", ws.Index)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port())), nil))
}
