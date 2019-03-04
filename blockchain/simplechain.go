package main

import (
	"crypto/sha256"
	"encoding/binary"
	"time"
	"bytes"
	"fmt"
)

type Block struct {
	previousHash []byte
	hash []byte
	height int
	timestamp int64
	data []byte
}

// 创建一个新区块
func New(height int, timestamp int64, data []byte, previousHash []byte) Block {
	return Block{
		previousHash,
		BlockHash256(previousHash, data, timestamp),
		height,
		timestamp,
		data,
	}

}

func (block *Block) String() string {
	return fmt.Sprintf("height: %d, data size: %d, time: %s", block.height, len(block.data), time.Unix(block.timestamp, 0).String())
}

// 将区块 取hash摘要
func BlockHash256(previousHash []byte, data []byte, timestamp int64) []byte {
	instance := sha256.New()
	instance.Write(previousHash)
	instance.Write(data)
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(timestamp))
	instance.Write(buf)
	return instance.Sum(nil)
}

// 使用链表链接区块
type BlockChain struct {
	block []Block
	height int
}

func (chain *BlockChain) getLastHash() Block {
	return chain.block[len(chain.block)]
}

// 初始化
func (chain *BlockChain) initChain() {
	chain.block = make([]Block, 32)
	chain.height = 0
	chain.block[0] = genGenesisBlock()
}

// 初始化
func (chain *BlockChain) addBlock(data []byte) {
	preBlock := chain.block[chain.height]
	chain.height ++
	chain.block[chain.height] = New(chain.height, time.Now().Unix(), data, preBlock.hash)
}

// 遍历区块链 并验证完整性 通过hash
func (chain *BlockChain) isValidChain() {
	for i, block := range chain.block {
		// 验证当前区块的hash
		tmpHash := BlockHash256(block.previousHash, block.data, block.timestamp)
		if !bytes.Equal(tmpHash, block.hash) {
			fmt.Println(fmt.Sprintf("block hash err, block=%s", block.String()))
			return
		}
		// 验证是否存有正确前置区块hash
		if i>0 && !bytes.Equal(chain.block[i-1].hash, block.previousHash){
			fmt.Println(fmt.Sprintf("block hash err, block=%s", chain.block[i-1].String()))
			return
		}
		fmt.Printf("correct block, %s\r\n", block.String())
		if i >= chain.height {
			break
		}
	}
}

// 生成创世区块，创世区块 没有指向前区块的hash，默认为0，在创世区块一般分配好初始余额 
func genGenesisBlock() Block {
	return New(0, time.Now().Unix(), []byte("hello world"), []byte{31})
}

func main() {
	var chain BlockChain
	chain.initChain()
	chain.addBlock([]byte("second"))
	chain.addBlock([]byte("third"))
	chain.isValidChain()

	//// 篡改区块内容
	//chain.block[1].data = []byte("fake")
	//chain.isValidChain()

	//// 篡改 重新计算hash
	//chain.block[1] = New(1, time.Now().Unix(), []byte("second2"), chain.block[0].hash)
	//chain.isValidChain()

	// 篡改 重新计算该区块及其后的所有hash
	chain.block[1] = New(1, time.Now().Unix(), []byte("second2"), chain.block[0].hash)
	chain.block[2].previousHash = chain.block[1].hash
	chain.block[2].hash = BlockHash256(chain.block[2].previousHash, chain.block[2].data, chain.block[2].timestamp)
	chain.isValidChain()
}
