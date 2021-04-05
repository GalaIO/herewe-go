package main

import (
	"bytes"
	"fmt"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
	"github.com/tyler-smith/go-bip39/wordlists"
	"math"
	"reflect"
	"strings"
)

const (
	DollarSymbol = 36
)

type preTree struct {
	val  int32
	next map[int32]*preTree
}

func newPreTree(val int32) *preTree {
	return &preTree{
		val:  val,
		next: make(map[int32]*preTree, 16),
	}
}

func main() {
	testFindShort()
	//testGenAccount()
	shortest := ""
	length := math.MaxInt32
	for i := 0; i < 10000; i++ {
		s := testGenShortAccount()
		if length > len(s) {
			shortest = s
			length = len(s)
		}
	}
	fmt.Println("shortest is", shortest)

}

func testFindShort() {
	english := wordlists.English
	root := newPreTree(DollarSymbol)
	for _, letter := range english {
		cur := root
		for _, ch := range letter {
			if cur.next[ch] == nil {
				cur.next[ch] = newPreTree(ch)
			}
			cur = cur.next[ch]
		}
		// tail
		if cur.next[DollarSymbol] == nil {
			cur.next[DollarSymbol] = newPreTree(DollarSymbol)
		}
		cur = cur.next[DollarSymbol]
	}

	travelTree(root, make([]byte, 0, 2048))
	//fmt.Println("result", result)
	//fmt.Println("shortMap", shortMap)
	//fmt.Println("shortReserveMap", shortReserveMap)
	//fmt.Println("total", len(wordlists.English))
	fmt.Println("result", len(result))
	fmt.Println("shortMap", len(shortMap))
	fmt.Println("shortReserveMap", len(shortReserveMap))
}

var result = make([]string, 0, 2048)
var shortMap = make(map[string]string, 2048)
var shortReserveMap = make(map[string]string, 2048)

func travelTree(cur *preTree, pre []byte) {
	if cur == nil {
		return
	}
	tmp := append(pre, byte(cur.val))
	if checkOnePath(cur) {
		short := parseWord(tmp)
		result = append(result, short)
		for len(cur.next) > 0 {
			for _, sub := range cur.next {
				tmp = append(tmp, byte(sub.val))
				cur = sub
				if len(cur.next) > 1 {
					panic(fmt.Errorf("violation rule!!!!, pre %v, cur %v", string(tmp), cur))
				}
			}
		}
		long := parseWord(tmp)
		shortMap[short] = long
		shortReserveMap[long] = short
		return
	}
	for _, node := range cur.next {
		travelTree(node, tmp)
	}
}

func parseWord(tmp []byte) string {
	if tmp[0] == DollarSymbol {
		tmp = tmp[1:]
	}
	lastIndex := len(tmp) - 1
	if tmp[lastIndex] == 36 {
		return string(tmp[:lastIndex])
	}
	return string(tmp)
}

func checkOnePath(node *preTree) bool {
	if len(node.next) > 1 {
		return false
	}
	for _, t := range node.next {
		return checkOnePath(t)
	}

	return true
}

func testGenShortAccount() string {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		panic(err)
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		panic(err)
	}
	fmt.Println("mnemonic:", mnemonic)

	split := strings.Split(strings.TrimSpace(mnemonic), " ")
	buffer := bytes.NewBuffer(nil)
	for i, s := range split {
		if i > 0 {
			buffer.WriteString(" ")
		}
		ss, ok := shortReserveMap[s]
		if !ok {
			panic(fmt.Errorf("cannot find buffer, %v, %v, %v, %v, %v, len %v", s, ok, shortMap[s], ss, reflect.TypeOf(s), len(shortReserveMap)))
		}
		buffer.WriteString(ss)
	}
	s := buffer.String()
	fmt.Printf("buffer mnemonic: %v, len %v\n", s, len(s))
	return s
}

func testGenAccount() {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		panic(err)
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		panic(err)
	}
	fmt.Println("mnemonic:", mnemonic)

	split := strings.Split(strings.TrimSpace(mnemonic), " ")
	short := bytes.NewBuffer(nil)
	for i, s := range split {
		if i > 0 {
			short.WriteString(" ")
		}
		ss, ok := shortReserveMap[s]
		if !ok {
			panic(fmt.Errorf("cannot find short, %v, %v, %v, %v, %v, len %v", s, ok, shortMap[s], ss, reflect.TypeOf(s), len(shortReserveMap)))
		}
		short.WriteString(ss)
	}
	fmt.Println("short mnemonic: ", short.String())

	seed := bip39.NewSeed(mnemonic, "")
	fmt.Println("seed is equal mnemonic: ", bytes.Equal(seed, entropy))

	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		panic(err)
	}
	account, err := wallet.Derive(hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0"), false)
	if err != nil {
		panic(err)
	}

	privateKeyHex, _ := wallet.PrivateKeyHex(account)
	fmt.Println("prikey: " + privateKeyHex)
	fmt.Println("address: " + account.Address.Hex())
}
