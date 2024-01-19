package main

import (
	"fmt"
	"sync"
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

// SingleNode 并发安全单链表
type SingleNode struct {
	Data any
	Next *SingleNode
}
type SingleList struct {
	mutex *sync.RWMutex
	Head  *SingleNode
	Tail  *SingleNode
	Size  uint
}

// Append 后插元素
func (List *SingleList) Append(node *SingleNode) bool {
	if node == nil {
		return false
	}
	List.mutex.Lock()
	defer List.mutex.Unlock()
	if List.Size == 0 {
		List.Head = node
		List.Tail = node
		List.Size = 1
		return true
	}
	tail := List.Tail
	tail.Next = node
	List.Tail = node
	List.Size += 1
	return true
}

// AddElem 前插元素
func (List *SingleList) AddElem(node *SingleNode) bool {
	if node == nil {
		return false
	}
	if List.Size == 0 {
		List.Head = node
		List.Tail = node
		List.Size = 1
		return true
	}
	node.Next = List.Head
	List.Head = node
	List.Size++
	return true
}

// Insert 指定位置插入
func (List *SingleList) Insert(index uint, node *SingleNode) bool {
	if node == nil || index > List.Size {
		return false
	}
	List.mutex.Lock()
	defer List.mutex.Unlock()
	if index == 0 {
		node.Next = List.Head
		List.Head = node
		List.Size++
		return true
	}
	pre := List.Head
	var count uint
	for count = 1; count < index; count++ {
		pre = pre.Next
	}
	next := pre.Next
	pre.Next = node
	node.Next = next
	List.Size++
	return true
}
func NewList() *SingleList {
	return &SingleList{
		Size:  0,
		Head:  nil,
		Tail:  nil,
		mutex: new(sync.RWMutex),
	}
}

// Display 输出链表
func (list *SingleList) Display() {
	if list == nil {
		fmt.Println("this single list is nil")
		return
	}
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	fmt.Printf("this single list size is %d \n", list.Size)
	ptr := list.Head
	var i uint
	for i = 0; i < list.Size; i++ {
		fmt.Printf("No%3d data is %v\n", i+1, ptr.Data)
		ptr = ptr.Next
	}
}

// SingleNodeS 使用哨兵实现单链表
type SingleNodeS struct {
	Data      any
	next, pre *SingleNodeS
	ListS     *SingleListS
}
type SingleListS struct {
	mutex *sync.RWMutex
	root  SingleNodeS
	Size  uint
}

func (l *SingleListS) Init() *SingleListS {
	l.root.next = &l.root
	l.root.pre = &l.root
	l.Size = 0
	return l
}
func (l *SingleListS) inset(e, root *SingleNodeS) {
	e.next = root.next
	root.next = e
	e.next.pre = e
	e.ListS = l
	l.Size++
}
func NewSingleNodes() *SingleListS {
	return new(SingleListS).Init()
}
func (l *SingleListS) PushBack(v any) {
	l.inset(&SingleNodeS{Data: v}, l.root.pre)
}
func (l *SingleListS) FrontBack(v any) {
	l.inset(&SingleNodeS{Data: v}, &l.root)
}
func main() {
	link := NewSingleNodes()
	link.PushBack("3")
	link.PushBack("32")
	link.FrontBack("front")
	link.FrontBack("frontfsdf")
	link.PushBack("33")
	link.PushBack("355")
}
