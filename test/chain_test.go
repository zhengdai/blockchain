package test

import (
	"block-chain/chain"
	"crypto/sha256"
	"fmt"
	"math/big"
	"testing"
)

func TestChain(t *testing.T) {
	bc := chain.NewBlockchain()
	defer bc.Db.Close()
	cli := chain.CLI{}
	cli.Run()
}

func TestHash(t *testing.T) {
	data1 := []byte("I like donuts")
	data2 := []byte("I like donutsca07ca")
	targetBits := 24
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	fmt.Printf("%x\n", sha256.Sum256(data1))
	fmt.Printf("%64x\n", target)
	fmt.Printf("%x\n", sha256.Sum256(data2))
}
