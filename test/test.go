package main

import (
	"fmt"
	"math/big"
)

const targetBits = 24

func main() {
	target := big.NewInt(12345789)
	fmt.Println("target 1:", target)
	target.Lsh(target, uint(3))
	fmt.Println("target 2:", target)
	i := 3
	x := i << 4
	fmt.Println("x", x)
}
