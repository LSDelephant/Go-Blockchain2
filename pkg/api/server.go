package api


import (
"encoding/json"
"log"
"net/http"


"github.com/yourusername/go-blockchain/pkg/blockchain"
"github.com/yourusername/go-blockchain/pkg/p2p"
"github.com/yourusername/go-blockchain/pkg/wallet"
)


type Server struct {
chain *blockchain.Blockchain
net *p2p.Network
mux *http.ServeMux
}


func NewServer(chain *blockchain.Blockchain, net *p2p.Network) *Server {
s := &Server{chain: chain, net: net, mux: http.NewServeMux()}
// chain endpoints
s.mux.HandleFunc("/blocks", s.handleGetBlocks)
s.mux.HandleFunc("/mine", s.handleMine)
// wallet
s.mux.HandleFunc("/wallet/new", s.handleNewWallet)
// peers (proxy to p2p)
s.mux.HandleFunc("/peers", net.peersHandler)
// mount p2p handlers under /p2p/
s.mux.Handle("/p2p/", http.StripPrefix("/p2p", net.Handler()))
return s
}


func (s *Server) Router() http.Handler {
return s.mux
}


func (s *Server) handleGetBlocks(w http.ResponseWr
