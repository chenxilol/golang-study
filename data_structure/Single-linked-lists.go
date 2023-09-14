package main

import (
	"fmt"
)

type Node struct {
	Data     any
	NodeNext *Node
}

type List struct {
	length   int //储存链表的长度
	headNode *Node
}

// insert data

func NewLinkList() *List {
	node := new(Node)
	return &List{
		headNode: node,
	}
}

// InsertForwardElem 指定位置插入
func (l *List) InsertForwardElem(index int, elem any) {
	if l.length < index+1 || index <= 0 {
		fmt.Println("插入的index错误，请重新输入")
	}
	pre := l.headNode
	node := &Node{Data: elem}
	if index == 1 {
		node.NodeNext = pre
		l.headNode = node
		return
	}
	for count := 1; count < index; count++ {
		pre = pre.NodeNext
	}
	node.NodeNext = pre.NodeNext
	pre.NodeNext = node
	l.length++
}

// DeleteForwardElem 删除指定位置的元素
func (l *List) DeleteForwardElem(index int) {
	if l.length < index+1 || index <= 0 {
		fmt.Println("插入的index错误，请重新输入")
		return
	}
	pre := l.headNode
	if index == 1 {
		l.headNode = pre.NodeNext.NodeNext
	}
	for count := 1; count < index; count++ {
		pre = pre.NodeNext
	}
	pre.NodeNext = pre.NodeNext.NodeNext
	l.length--
	fmt.Println("success deleted")
}

// RemoveElem 删除节点中第一个指定元素
func (l *List) RemoveElem(data any) {
	pre := l.headNode
	index := 0
	if pre.Data == data {
		l.headNode = pre.NodeNext
		fmt.Println("ok")
	}
	for pre != nil {
		index++
		if pre.NodeNext.Data == data {
			pre.NodeNext = pre.NodeNext.NodeNext
			fmt.Println("ok")
			l.length--
			return
		}
		pre = pre.NodeNext
	}
	fmt.Println("fail")
}

// RemoveElemAll 删除所有元素
func (l *List) RemoveElemAll(data any) {
	pre := l.headNode
	if pre.Data == data {
		l.headNode = pre.NodeNext
		l.length--
		fmt.Println("ok")
	}
	for pre.NodeNext != nil {
		if pre.NodeNext.Data == data {
			pre.NodeNext = pre.NodeNext.NodeNext
			fmt.Println("ok")
			l.length--
			if pre.NodeNext != nil {
				for pre.NodeNext.Data == data {
					pre.NodeNext = pre.NodeNext.NodeNext
					fmt.Println("ok")
					l.length--

					if pre.Data != data {
						break
					}
					if pre == nil {
						return
					}
					pre = pre.NodeNext
				}
			}
		}
		pre = pre.NodeNext
		if pre == nil {
			return
		}
	}
}

// AddElemForward 前插元素
func (l *List) AddElemForward(data any) {
	node := &Node{Data: data}
	if l.IsNull() {
		l.headNode = node
		l.length++
		return
	}
	node.NodeNext = l.headNode
	l.headNode = node
	l.length++
	return
}

// AppendElemBack 后插元素
func (l *List) AppendElemBack(data any) {
	node := &Node{Data: data}
	if l.IsNull() {
		l.headNode = node
		l.length++
		return
	}
	cur := l.headNode
	for cur.NodeNext != nil {
		cur = cur.NodeNext
	}
	cur.NodeNext = node
	l.length++
}

func (n *Node) Search() {
	for {
		if n.NodeNext != nil {
			fmt.Println(n.Data)
		}
		break
	}
}
func (l *List) IsNull() bool {
	if l.length == 0 {
		return true
	} else {
		return false
	}
}
func (l *List) ShowList() {
	if !l.IsNull() {
		cur := l.headNode
		for {
			fmt.Printf("\t%v", cur.Data)
			if cur.NodeNext != nil {
				cur = cur.NodeNext
			} else {
				break
			}
		}
	}
}

// SearchElem 查找是否包含指定值
func (l *List) SearchElem(data any) bool {
	/*单链表的按值查找
	  1、用指针p指向首元结点
	  2、从首元结点开始以此顺着链域next向下查找，只要指向当前结点的指针p不为空，
	  并且p所指结点的数据域不等于给定值e，则执行以下操作：p指向下一个结点
	  3、返回p。若查找成功，p此时即为结点的地址值，若查找失败，p的值即为NULL。
	*/
	if l.IsNull() {
		fmt.Println("err")
	}
	pre := l.headNode
	for pre != nil {
		if pre.Data == data {
			return true
		}
		pre = pre.NodeNext
	}
	return false

}
func main() {
	Link := NewLinkList()
	Link.AppendElemBack(3)
	Link.AppendElemBack(4)
	Link.AppendElemBack("5")
	Link.AppendElemBack(5)
	Link.AppendElemBack(553)
	Link.AppendElemBack(5)
	Link.AppendElemBack(5)
	Link.AppendElemBack(553)
	Link.AppendElemBack(5)
	Link.RemoveElemAll(5)
	Link.ShowList()
	Link.InsertForwardElem(2, 8)
	Link.ShowList()
}
