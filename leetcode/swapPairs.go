package main

import (
	"strconv"
	"fmt"
)

// 24. 两两交换链表中的节点
// https://leetcode-cn.com/problems/swap-nodes-in-pairs/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func (l *ListNode) String() string {
	res := ""
	for l!=nil {
		res += strconv.Itoa(l.Val) + "->"
		l = l.Next
	}
	res += "nil"
	return res
}
func swapPairs(head *ListNode) *ListNode {
	// 头结点
	preHead := &ListNode{Next:head}
	preNode := preHead

	for preNode.Next != nil && preNode.Next.Next != nil {
		first := preNode.Next
		second := first.Next
		temp := second.Next
		preNode.Next = second
		second.Next = first
		first.Next = temp
		preNode = first
	}

	return preHead.Next
}

func main() {
	case1 := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4,
					nil,
				},
			},
		},
	}
	fmt.Println(swapPairs(case1))
	case2 := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				nil,
			},
		},
	}
	fmt.Println(swapPairs(case2))
}
