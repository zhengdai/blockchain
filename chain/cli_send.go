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
	defer bc.Db.Close()

	tx := NewUTXOTransaction(from, to, amount, bc)
	bc.MineBlock([]*Transaction{tx})
	fmt.Println("send success!")
}
