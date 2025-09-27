package blockchain


import (
"sync"
"time"
)


// Blockchain represents the chain and mutex for concurrent access
type Blockchain struct {
blocks []*Block
mu sync.RWMutex
}


func NewBlockchain() *Blockchain {
b := &Blockchain{blocks: make([]*Block, 0)}
// genesis
gen := &Block{
Index: 0,
Timestamp: time.Now().Unix(),
PrevHash: "0",
Data: "genesis",
}
gen.Hash = gen.calculateHash()
b.blocks = append(b.blocks, gen)
return b
}


func (bc *Blockchain) GetBlocks() []*Block {
bc.mu.RLock()
defer bc.mu.RUnlock()
copyBlocks := make([]*Block, len(bc.blocks))
copy(copyBlocks, bc.blocks)
return copyBlocks
}


func (bc *Blockchain) LastBlock() *Block {
bc.mu.RLock()
defer bc.mu.RUnlock()
return bc.blocks[len(bc.blocks)-1]
}


func (bc *Blockchain) AddBlock(data string) *Block {
bc.mu.Lock()
defer bc.mu.Unlock()
last := bc.blocks[len(bc.blocks)-1]
new := &Block{
Index: last.Index + 1,
Timestamp: time.Now().Unix(),
PrevHash: last.Hash,
Data: data,
}
// simple PoW
nonce := int64(0)
for {
new.Nonce = nonce
new.Hash = new.calculateHash()
if validHashDifficulty(new.Hash) {
break
}
nonce++
}
bc.blocks = append(bc.blocks, new)
return new
}


func (bc *Blockchain) IsValidChain(blocks []*Block) bool {
if len(blocks) == 0 {
return false
}
}
