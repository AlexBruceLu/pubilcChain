package blc

type BlockChain struct {
	Blocks []*Block
}

// 创建带有创世区块的BlockChain
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

// 新增区块
func (blockChain *BlockChain) AddBlock(data string) {
	preBlock := blockChain.Blocks[len(blockChain.Blocks)-1]
	newBlock := NewBlock(data,preBlock.Hash)
	blockChain.Blocks = append(blockChain.Blocks,newBlock)
}