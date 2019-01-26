package main

import "math/big"

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

const targetBits = 24

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}
