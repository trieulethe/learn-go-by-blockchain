package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

type MerkleTree struct {
	root   *Node
	leaves [][]byte
}

type Node struct {
	data  []byte
	left  *Node
	right *Node
}

func calculateHash(leaves [][]byte) {
	h := sha256.New()
	length := len(leaves)
	if length%2 != 0 {
		leaves = append(leaves, leaves[length-1])
	}
	for i := 0; i < length; i += 2 {
		chash := append(leaves[i], leaves[i+1])
		h.Write(chash)
	}
}

func Greeting(prefix string, who ...string) {
	fmt.Println("prefix", prefix)
	fmt.Println("who: ", who)
}

func main() {
	h := sha256.New()
	h.Write([]byte("hello world"))
	fmt.Printf("%x", h.Sum(nil))
	fmt.Println()
	fmt.Println("hash:", h.Sum(nil))
	sha := base64.URLEncoding.EncodeToString(h.Sum(nil))
	fmt.Println("sha:", sha)
	hexcode := hex.EncodeToString(h.Sum(nil))
	fmt.Println("hexcode:", hexcode)
	s := []string{"James", "Jasmine"}
	Greeting("goodbye:", s...)
}
