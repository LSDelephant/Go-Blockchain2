package p2p


import (
"encoding/json"
"io"
"log"
"net/http"
"sync"


"github.com/yourusername/go-blockchain/pkg/blockchain"
)


type Network struct {
peers map[string]struct{}
mu sync.RWMutex
chain *blockchain.Blockchain
}


func NewNetwork(chain *blockchain.Blockchain) *Network {
n := &Network{peers: make(map[string]struct{}), chain: chain}
return n
}


func (n *Network) AddPeer(addr string) {
n.mu.Lock()
defer n.mu.Unlock()
n.peers[addr] = struct{}{}
}


func (n *Network) Peers() []string {
n.mu.RLock()
defer n.mu.RUnlock()
out := make([]string, 0, len(n.peers))
for p := range n.peers {
out = append(out, p)
}
return out
}


// BroadcastBlocks posts local chain to peers (simple naive approach)
func (n *Network) BroadcastChain() {
peers := n.Peers()
blocks := n.chain.GetBlocks()
bts, _ := json.Marshal(blocks)
for _, p := range peers {
resp, err := http.Post(p+"/receive-chain", "application/json", io.NopCloser(bytes.NewReader(bts)))
if err != nil {
log.Printf("failed to send chain to %s: %v", p, err)
continue
}
resp.Body.Close()
}
}


// HTTP handlers that a node exposes for P2P
func (n *Network) Handler() http.Handler {
mux := http.NewServeMux()
mux.HandleFunc("/receive-chain", n.receiveChain)
mux.HandleFunc("/blocks", n.blocksHandler)
mux.HandleFunc("/peers", n.peersHandler)
return mux
}


func (n *Network) blocksHandler(w http.ResponseWriter, r *http.Request) {
if r.Method != http.MethodGet {
w.WriteHeader(http.StatusMethodNotAllowed)
return
}
json.NewEncoder(w).Encode(n.chain.GetBlocks())
}


func (n *Network) peersHandler(w http.ResponseWriter, r *http.Request) {
switch r.Method {
case http.MethodGet:
}
