package main

import (
	"fmt"
	"github.com/zhao/goblockchain/block"
)

func main() {

	block := block.NewBlock(1, nil, []byte("this first block"))
	fmt.Println(block)
}
