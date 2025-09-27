package blockchain


// Difficulty settings for PoW. Simple fixed difficulty.
const Difficulty = 4


func validHashDifficulty(hash string) bool {
// require hash to start with `Difficulty` zeros in hex representation
prefix := ""
for i := 0; i < Difficulty; i++ {
prefix += "0"
}
return len(hash) >= Difficulty && hash[:Difficulty] == prefix
}
