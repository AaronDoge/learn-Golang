package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

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
