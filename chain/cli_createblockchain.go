package chain

import (
	"fmt"
	"log"
)

func (cli *CLI) createBlockchain(address, nodeID string) {
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := CreateBlockchain(address, nodeID)
	defer bc.Db.Close()
	UTXOSet := UTXOSet{bc}
	UTXOSet.Reindex()
	fmt.Println("Create block chain Done!")
}
