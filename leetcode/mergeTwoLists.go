package main

import (
	"fmt"
	"strconv"
)

// 21. 合并两个有序链表
// https://leetcode-cn.com/problems/merge-two-sorted-lists/

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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	head := &ListNode{}
	tmp := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		}else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}

	for l1 != nil {
		tmp.Next = l1
		tmp = tmp.Next
		l1 = l1.Next
	}
	for l2 != nil {
		tmp.Next = l2
		tmp = tmp.Next
		l2 = l2.Next
	}
	tmp.Next = nil

	return head.Next
}

func main() {
	case1_l1 := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				4,
				nil,
			},
		},
	}
	case1_l2 := &ListNode{
		1,
		&ListNode{
			3,
			&ListNode{
				4,
				nil,
			},
		},
	}
	lists := mergeTwoLists(case1_l1, case1_l2)
	fmt.Println(lists)
}
