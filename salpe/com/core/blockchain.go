package core

import (
	"crypto/sha256"
	"encoding/hex"
)

// Block is model for block
type Block struct {
	Hash          string
	PrevBlockHash string
	Data          string
}

// BlockChain is collection of blocks
type BlockChain struct {
	Blocks []*Block
}

func (b *Block) setHash() {
	hash := sha256.Sum256([]byte(b.PrevBlockHash + b.Data))
	b.Hash = hex.EncodeToString(hash[:])

}

// NewBlock creates new block in blockchain
func NewBlock(data string, PrevBlockHash string) *Block {
	block := &Block{
		Data: data, PrevBlockHash: PrevBlockHash}
	return block
}

func (bc *BlockChain) AddBlock(data string) *Block {

	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
	return newBlock

}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

func NewGenesisBlock() *Block {
	return NewBlock("genesis", "")
}
