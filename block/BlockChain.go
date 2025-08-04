package block

type BlockChain struct {
	Blocks []*Block
}

func CreateBlockChainWithGenesisBlock() *BlockChain {
	block := CreateGenesisBlock([]byte("Genesis Block"))
	return &BlockChain{[]*Block{block}}
}
func (bc *BlockChain) AddBlock(height int64, prevBlock []byte, data []byte) {
	newBlock := NewBlock(height, prevBlock, data)
	bc.Blocks = append(bc.Blocks, newBlock)
}
