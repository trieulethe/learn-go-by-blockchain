package main

import "fmt"

func main() {
	bc := NewBlockchain()

	bc.AddBlock("block 1")
	bc.AddBlock("block 2")

	for _, block := range bc.blocks {
		fmt.Printf("Pre hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println("")
	}

}
