package wallet


import (
"crypto/ecdsa"
"crypto/elliptic"
"crypto/rand"
"crypto/sha256"
"encoding/hex"
"fmt"
)


// Простий гаманець на ECDSA


func NewKeyPair() (*ecdsa.PrivateKey, []byte, error) {
priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
if err != nil {
return nil, nil, err
}
pub := append(priv.PublicKey.X.Bytes(), priv.PublicKey.Y.Bytes()...)
addr := addressFromPubKey(pub)
return priv, addr, nil
}


func addressFromPubKey(pub []byte) []byte {
h := sha256.Sum256(pub)
return h[:20] // take first 20 bytes as "address"
}


func AddrToString(addr []byte) string {
return hex.EncodeToString(addr)
}


func Sign(priv *ecdsa.PrivateKey, data []byte) ([]byte, error) {
h := sha256.Sum256(data)
r, s, err := ecdsa.Sign(rand.Reader, priv, h[:])
if err != nil {
return nil, err
}
// simple serialization r||s
rb := r.Bytes()
sb := s.Bytes()
return append(rb, sb...), nil
}


func Verify(pub *ecdsa.PublicKey, data, sig []byte) bool {
if len(sig)%2 != 0 {
return false
}
h := sha256.Sum256(data)
hlen := len(sig) / 2
r := new(big.Int).SetBytes(sig[:hlen])
s := new(big.Int).SetBytes(sig[hlen:])
return ecdsa.Verify(pub, h[:], r, s)
}
