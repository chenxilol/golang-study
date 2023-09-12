package main

import "fmt"

// Tree 二叉树的实现
type Tree struct {
	Data      string
	LeftNode  *Tree
	RightNode *Tree
}

// PreOrder 创建先序遍历
func PreOrder(t *Tree) {
	// 先判断t是否为nil
	if t == nil {
		return
	}
	fmt.Println(t.Data, "  ")
	PreOrder(t.LeftNode)
	PreOrder(t.RightNode)

}

// MidOrder 中序遍历
func MidOrder(tree *Tree) {
	if tree == nil {
		return
	}
	MidOrder(tree.LeftNode)
	fmt.Println(tree.Data, " ")
	MidOrder(tree.RightNode)
}

// 在树中插入数据
func main() {
	t := new(Tree)
	t = &Tree{
		Data: "A",
	}
	t.LeftNode = &Tree{Data: "B"}
	t.RightNode = &Tree{Data: "C"}
	t.LeftNode.LeftNode = &Tree{Data: "D"}
	t.LeftNode.RightNode = &Tree{Data: "E"}
	t.RightNode.RightNode = &Tree{Data: "F"}
	PreOrder(t)
	MidOrder(t)
}
