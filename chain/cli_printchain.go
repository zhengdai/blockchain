package chain

import (
	"fmt"
	"strconv"
)

func (cli *CLI) printChain() {
	bc := NewBlockchain()
	defer bc.Db.Close()
	bci := bc.Iterator()
	for bci.CurrentHash != nil {
		block := bci.Next()
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %x\n", block.Nonce)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		for _, tx := range block.Transactions {
			fmt.Println(tx)
		}
		fmt.Println()
	}
}
