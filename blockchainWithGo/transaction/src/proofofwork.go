package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const targetBits = 24

var(
	maxNonce = math.MaxInt64
)

type ProofOfWork struct {
	block *Block 
	target *big.Int 
}

func NewProofOfWork(b *Block) *ProofOfWork{
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow 
}
