package main


import (
"log"
"net/http"
"os"
"strconv"


"github.com/yourusername/go-blockchain/pkg/api"
"github.com/yourusername/go-blockchain/pkg/blockchain"
"github.com/yourusername/go-blockchain/pkg/p2p"
)


func main() {
port := 8000
if p := os.Getenv("PORT"); p != "" {
if v, err := strconv.Atoi(p); err == nil {
port = v
}
}


chain := blockchain.NewBlockchain()
net := p2p.NewNetwork(chain)
server := api.NewServer(chain, net)


addr := ":" + strconv.Itoa(port)
log.Printf("Starting node on %s...", addr)
log.Fatal(http.ListenAndServe(addr, server.Router()))
}
