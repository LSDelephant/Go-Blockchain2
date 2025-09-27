package blockchain


import (
"crypto/sha256"
"encoding/hex"
"fmt"
"time"
)


// Block represents a single block in the blockchain
type Block struct {
Index int `json:"index"`
Timestamp int64 `json:"timestamp"`
PrevHash string `json:"prev_hash"`
Hash string `json:"hash"`
Nonce int64 `json:"nonce"`
Data string `json:"data"`
}


func (b *Block) calculateHash() string {
h := sha256.New()
h.Write([]byte(fmt.Sprintf("%d%d%s%d%s", b.Index, b.Timestamp, b.PrevHash, b.Nonce, b.Data)))
return hex.EncodeToString(h.Sum(nil))
}


func NewBlock(index int, prevHash, data string) *Block {
b := &Block{
Index: index,
Timestamp: time.Now().Unix(),
PrevHash: prevHash,
Data: data,
}
b.Hash = b.calculateHash()
return b
}
