package main

//Blockchain 是一个Block指针数组
type Blockchain struct {
	blocks []*Block
}

//创建一个有创世区块的链
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

//向联众加入一个新块
//data在实际中就是交易数据
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}
