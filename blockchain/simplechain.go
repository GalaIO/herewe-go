package main

import (
	"crypto/sha256"
	"encoding/binary"
	"time"
	"bytes"
	"fmt"
	"strconv"
	"encoding/base64"
)

type Block struct {
	previousHash []byte
	hash []byte
	height int
	timestamp int64
	data []byte
	// 增肌hash时 随机数用于pow
	nonce int
}

// 创建一个新区块
func New(height int, timestamp int64, data []byte, previousHash []byte) Block {
	// 初始随机值
	nonce := 0
	block := Block{
		previousHash,
		nil,
		height,
		timestamp,
		data,
		nonce,
	}
	// 挖矿，保证验证一致，直至等到网络出块
	nonce, hash := block.mineBlock(nonce, 6)
	block.nonce = nonce
	block.hash = hash
	return block

}

func (block *Block) String() string {
	return fmt.Sprintf("height: %d, data size: %d, time: %s, nonce: %d, hash: %s", block.height, len(block.data), time.Unix(block.timestamp, 0).String(),
		block.nonce, strconv.Itoa(int(block.hash[0])) + "||" + base64.StdEncoding.EncodeToString(block.hash))
}

// 将区块 取hash摘要
func BlockHash256(previousHash []byte, data []byte, nonce int, timestamp int64) []byte {
	instance := sha256.New()
	instance.Write(previousHash)
	instance.Write(data)
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(timestamp))
	instance.Write(buf)
	binary.BigEndian.PutUint64(buf, uint64(nonce))
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
		tmpHash := BlockHash256(block.previousHash, block.data, block.nonce, block.timestamp)
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

// 挖矿程序，pow机制，前n位为0，简单期间n必须整除4
func (block *Block) mineBlock(nonce int, difficulty int) (int, []byte) {
	hashed := BlockHash256(block.previousHash, block.data, nonce, block.timestamp)
	for !validMinedBlockHash(hashed, difficulty) {
		nonce ++
		hashed = BlockHash256(block.previousHash, block.data, nonce, block.timestamp)
	}
	return nonce, hashed
}

// 验证hash是否满足前difficulty位 都是0
var zeroBytes = [8]byte{
	0x00,
	0x7F,
	0x3F,
	0x1F,
	0x0F,
	0x07,
	0x03,
	0x01,
}
func validMinedBlockHash(hashed []byte, difficulty int) bool {
	left := difficulty % 8
	allZeroCount := difficulty / 8
	for i, cbyte := range hashed {
		if i >= allZeroCount {
			break
		}
		if cbyte != 0x00 {
			return false
		}
	}
	if left <= 0 {
		return true
	}
	return hashed[allZeroCount] & (^zeroBytes[left]) == 0x00
}

// 生成创世区块，创世区块 没有指向前区块的hash，默认为0，在创世区块一般分配好初始余额 
func genGenesisBlock() Block {
	return New(0, time.Now().Unix(), []byte("hello world"), []byte{31})
}

func main() {
	//fmt.Println(validMinedBlockHash([]byte{0x00, 0x31}, 10))
	sampleBlock()
}

func sampleBlock() {
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
	//chain.block[1] = New(1, time.Now().Unix(), []byte("second2"), chain.block[0].hash)
	//chain.block[2].previousHash = chain.block[1].hash
	//chain.block[2].hash = BlockHash256(chain.block[2].previousHash, chain.block[2].data, chain.block[2].nonce, chain.block[2].timestamp)
	//chain.isValidChain()
}
