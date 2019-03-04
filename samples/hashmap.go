package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// hash
// 不同长度变成相同长度，
// 快速生成 逆向困难 输入敏感 避免冲突

func main() {
	head := CreateHead("1", 1)
	head.AddNode("2", 2)
	head.AddNode("3", 3)
	head.AddNode("4", 4)
	fmt.Println(head)
	fmt.Println(head.FindNode("3"))
	fmt.Println(head.FindNode("5"))
	fmt.Println(head)
	fmt.Println(head.RemoveNode("2"))
	fmt.Println(head)
	var hmap = HashMap{}
	hmap.put("test", 2)
	fmt.Println(hmap)
	fmt.Println(hmap.get("test"))
}

type KV struct {
	key string
	value int
}

func (kv *KV) String() string {
	return kv.key + ": " + strconv.Itoa(kv.value)
}

type Node struct {
	KV
	NextNode *Node
}

func CreateHead(key string, val int) *Node {
	var head = &Node{KV{key, val}, nil}
	return head
}

func (node *Node)AddNode(key string, val int) bool {
	nNode := &Node{KV{key, val}, nil}
	for ; node.NextNode!= nil;  {
		node = node.NextNode
	}
	node.NextNode = nNode
	return true
}

func (node *Node) FindNode(key string) *Node {
	for ; node!= nil;  {
		if node.key == key {
			return node
		}
		node = node.NextNode
	}
	return nil
}

func (node *Node) RemoveNode(data string) bool {
	if node.key == data {
		node = node.NextNode
		return true
	}
	preNode := node
	for ; node!= nil;  {
		if node.key == data {
			preNode.NextNode = node.NextNode
			return true
		}
		preNode = node
		node = node.NextNode
	}
	return false
}

func (node *Node) String() string {
	buffer := bytes.Buffer{}
	for ; node != nil;  {
		buffer.WriteString(node.key + ": " + strconv.Itoa(node.value))
		buffer.WriteString(" -> ")
		node = node.NextNode
	}
	buffer.WriteString("nil\r\n")
	return buffer.String()
}

type HashMap struct {
	List [16]Node
}

func (hashMap *HashMap) put(key string, val int) bool {
	index := hashCode(key)
	return hashMap.List[index].AddNode(key, val)
}

func (hashMap *HashMap) get(key string) int {
	index := hashCode(key)
	node := hashMap.List[index].FindNode(key)
	if node == nil {
		panic("cannot find")
	}
	return node.value
}

func (hashMap *HashMap) String() string {
	buffer := bytes.Buffer{}
	for _, node := range hashMap.List {
		buffer.WriteString("| ")
		buffer.WriteString(node.String())
	}
	return buffer.String()
}

// 核心散列函数
func hashCode(key string) int {
	hash := 1
	for ch := range key {
		hash *= 1103515245 + ch
	}
	return ((hash) >> 27) & (15)
}