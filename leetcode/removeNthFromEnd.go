package main

import "fmt"

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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return  nil
	}
	preNode := head
	i := 1
	for fooNode:=head; fooNode!=nil; i++{
		if i > n+1 && preNode != nil{
			preNode = preNode.Next
		}
		fooNode = fooNode.Next
	}

	if preNode != nil && preNode.Next != nil && i>n+1 {
		preNode.Next = preNode.Next.Next
	}

	if preNode == head && i>=n+1 {
		return preNode.Next
	}

	return head
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
					&ListNode{
						5,
						nil,
					},
				},
			},
		},
	}
	end := removeNthFromEnd(case1, 2)
	fmt.Println(end)

	case2 := &ListNode{1, nil}
	end2 := removeNthFromEnd(case2, 1)
	fmt.Println(end2)

	case3 := &ListNode{1, &ListNode{2, nil}}
	end3 := removeNthFromEnd(case3, 2)
	fmt.Println(end3)
}