package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Stack struct  {
	items []*TreeNode
	idx int
}

func NewStack(len int) *Stack {
	return &Stack {
		items: make([]*TreeNode, len),
		idx: 0,
	}
}

func (s *Stack) print() {
	fmt.Println(s.idx, len(s.items))
}

func (s *Stack) push(node *TreeNode) {
	if s.idx >= len(s.items) {
		s.items = append(s.items, node)
		s.idx = s.idx + 1
		return
	}
	s.items[s.idx] = node
	s.idx = s.idx + 1
}

func (s *Stack) pop() (*TreeNode, bool) {
	if s.idx <= 0 {
		return nil, false
	}

	s.idx = s.idx - 1
	return s.items[s.idx], true
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0, 100000)
	stack := NewStack(100000)
	stack.push(root)

	for node, exist := stack.pop(); exist; node, exist = stack.pop() {
		//stack.print()
		if node == nil {
			continue
		}
		res = append(res, node.Val)
		stack.push(node.Left)
		stack.push(node.Right)
	}
	return res
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: nil,
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	fmt.Println(preorderTraversal(root))
}
