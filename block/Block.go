package block

import (
	"bytes"
	"crypto/sha256"
	"time"
)

type Block struct {
	TimeStamp int64  //block time
	Hash      []byte //current hash
	PrevHash  []byte //prev hash
	Height    int64  //height of the block
	Data      []byte //data of the block
}

func NewBlock(height int64, prevHash []byte, data []byte) *Block {
	block := Block{
		TimeStamp: time.Now().Unix(),
		Hash:      nil,
		PrevHash:  prevHash,
		Height:    height,
		Data:      data,
	}
	block.SetHash()
	return &block
}

func (block *Block) SetHash() {
	blockBytes := bytes.Join([][]byte{
		intToHex(block.Height),
		intToHex(block.TimeStamp),
		block.PrevHash,
		block.Data,
	}, []byte{})
	hash := sha256.Sum256(blockBytes)
	block.Hash = hash[:]
}

func CreateGenesisBlock(data []byte) *Block {
	return NewBlock(1, nil, data)
}
