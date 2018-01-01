package main

import (
	"fmt"
)
import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

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

/*
	//Block 由区块头和交易两部分构成
	//Timestamp, PrevBlockHash, Hash 属于区块头(block header)
	//Timestamp, 当前时间戳，也就是区块创建的时间
	//PrevBlockHash, 前一个块的哈希
	//Hash, 当前块的hash
	//Data, 区块实际存储的信息，比特币中也就是交易
*/
type Block struct {
	Timestamp     int64
	PrevBlockHash []byte
	Hash          []byte
	Data          []byte
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Data:          []byte(data),
	}

	block.SetHash()

	return block
}

//设置当前块的hash
//Hash=sha256(PrevBlockHash + Data + Timestamp)
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

//生成创世区块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
		fmt.Printf("Prev hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
