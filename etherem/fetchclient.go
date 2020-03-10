package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"context"
	"math/big"
	"fmt"
	"math"
	"github.com/ethereum/go-ethereum/crypto"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
)

func main() {
	//queryBalance()

	//generateEthAccount()

	//importPrikey()

	//queryBlockAndPrintHeader()

	parseEventByABI()
}

func parseEventByABI() {
	client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")
	if err != nil {
		panic(err)
	}
	contractAddress := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(2394201),
		ToBlock:   big.NewInt(2394201),
		Addresses: []common.Address{
			contractAddress,
		},
	}
	json := `[
		{
			"constant": true,
			"inputs": [
				{
					"name": "",
					"type": "bytes32"
				}
			],
			"name": "items",
			"outputs": [
				{
					"name": "",
					"type": "bytes32"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		},
		{
			"constant": true,
			"inputs": [],
			"name": "version",
			"outputs": [
				{
					"name": "",
					"type": "string"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		},
		{
			"constant": false,
			"inputs": [
				{
					"name": "key",
					"type": "bytes32"
				},
				{
					"name": "value",
					"type": "bytes32"
				}
			],
			"name": "setItem",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"name": "_version",
					"type": "string"
				}
			],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "constructor"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": false,
					"name": "key",
					"type": "bytes32"
				},
				{
					"indexed": false,
					"name": "value",
					"type": "bytes32"
				}
			],
			"name": "ItemSet",
			"type": "event"
		}
		]`
	contractAbi, err := abi.JSON(strings.NewReader(json))
	if err != nil {
		panic(err)
	}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		panic(err)
	}
	for _, vLog := range logs {
		fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
		fmt.Println(vLog.BlockNumber)     // 2394201
		fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6

		event := struct {
			Key   [32]byte
			Value [32]byte
		}{}

		err := contractAbi.Unpack(&event, "ItemSet", vLog.Data)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(event.Key[:]))   // foo
		fmt.Println(string(event.Value[:])) // bar

		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}

		fmt.Println(topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
	}
	eventSignature := []byte("ItemSet(bytes32,bytes32)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println(hash.Hex())
	// 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
}

func queryBlockAndPrintHeader() {
	// 从infura 链接全节点
	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		panic(err)
	}
	bestBlock, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	header := bestBlock.Header()
	jsonbs, _ := (*header).MarshalJSON()
	fmt.Println(string(jsonbs))
}

func importPrikey() {
	// 导入16进制私钥
	privateKey, err := crypto.HexToECDSA("0E75359F0F7463A1CBE9FF7982B8F8F71437F46823EA164A9F3C65B4DD167482")
	if err != nil {
		panic(err)
	}
	ecdsaPub := privateKey.Public().(*ecdsa.PublicKey)
	// 获取地址
	fmt.Println(crypto.PubkeyToAddress(*ecdsaPub).Hex())
}

func generateEthAccount() {
	// 生产ECDSA密钥对
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	// 获取私钥 并打印16进制
	privBytes := crypto.FromECDSA(privateKey)
	hexs := common.Bytes2Hex(privBytes)
	fmt.Println(hexs)
	// 获取公钥 和 地址
	ecdsaPub := privateKey.Public().(*ecdsa.PublicKey)
	// 地址
	fmt.Println(crypto.PubkeyToAddress(*ecdsaPub).Hex())
}

func queryBalance() {
	// 从infura 链接全节点
	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		panic(err)
	}
	// 解析hex格式 地址
	address := common.HexToAddress("0x96C8402cd4750272e2a001845893DEdA58456915")
	// 发起rpc查询
	balanceAt, err := client.BalanceAt(context.Background(), address, nil)
	fmt.Println(balanceAt.Int64())
	// 大数计算 余额
	balance := new(big.Float).SetInt64(balanceAt.Int64())
	fmt.Println(balance.Quo(balance, big.NewFloat(math.Pow10(18))))
}
