package main

import (
	"fmt"
	"github.com/zhao/goblockchain/block"
)

func main() {

	bc := block.CreateBlockChainWithGenesisBlock()
	fmt.Println(bc)
	fmt.Println(bc.Blocks[0])
	prev := bc.Blocks[len(bc.Blocks)-1]
	bc.AddBlock(prev.Height-1, prev.Hash, []byte("alice"))
	prev = bc.Blocks[len(bc.Blocks)-1]
	bc.AddBlock(prev.Height-1, prev.Hash, []byte("bob"))
	for _, item := range bc.Blocks {
		fmt.Printf("Prev Hash:%x, Current Hash:%x\n", item.PrevHash, item.Hash)
	}
}
