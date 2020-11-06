package chain

import "fmt"

func (cli *CLI) reindexUTXO() {
	bc := NewBlockchain()
	UTXOSet := UTXOSet{bc}
	UTXOSet.Reindex()
	count := UTXOSet.CountTransactions()
	fmt.Printf("reindex success, threr ar %d transactions in the UTXO set.\n", count)
}
