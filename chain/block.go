package chain

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
	Height        int
}

func NewBlock(transactions []*Transaction, preBlockHash []byte, height int) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Transactions:  transactions,
		PrevBlockHash: preBlockHash,
		Hash:          []byte{},
		Nonce:         0,
		Height:        height,
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{}, 0)
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}

func (b *Block) HashTransactions() []byte {
	var transactions [][]byte
	for _, tx := range b.Transactions {
		transactions = append(transactions, tx.Serialize())
	}
	mTree := NewMerkleTree(transactions)
	return mTree.RootNode.Data
}
