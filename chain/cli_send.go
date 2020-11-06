package chain

import (
	"fmt"
	"log"
)

func (cli *CLI) send(from, to string, amount int) {
	if !ValidateAddress(from) {
		log.Panic("ERROR: sender address is not valid")
	}
	if !ValidateAddress(to) {
		log.Panic("ERROR: recipient address is not valid")
	}
	bc := NewBlockchain()
	UTXOSet := UTXOSet{bc}
	defer bc.Db.Close()

	tx := NewUTXOTransaction(from, to, amount, &UTXOSet)
	cbTx := NewCoinbaseTX(from, "")
	txs := []*Transaction{cbTx, tx} //两笔交易
	newBlock := bc.MineBlock(txs)
	UTXOSet.Update(newBlock)
	fmt.Println("send success!")
}
