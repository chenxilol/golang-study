package main

import "fmt"

// Element 创建双链表的节点
type Node1 struct {
	next, prev *Node1
	Value      any
	Key        string
	list       *LRU
}
type LRU struct {
	Length   uint
	maxBytes uint
	Cache    map[string]*Node1
	root     Node1
}

func NewLRU() *LRU {
	return new(LRU).Init()
}
func (l *LRU) Init() *LRU {
	l.root.next = &l.root
	l.maxBytes = 5
	l.root.prev = &l.root
	l.Cache = make(map[string]*Node1)
	l.Length = 0
	return l
}
func (l *LRU) Append(key string, value any) {
	if nowNode, ok := l.Cache[key]; !ok {
		node := Node1{Value: value, Key: key}
		l.insert(&node, l.root.prev)
		l.Cache[key] = &node
		if l.Length > l.maxBytes {
			l.delete(l.root.prev)
			delete(l.Cache, l.root.prev.Key)
		}
	} else {
		l.FrontBack(key, value)
		l.delete(nowNode)
	}
}

// FrontBack 前插
func (l *LRU) FrontBack(k string, v any) {
	node := &Node1{Value: v, Key: k}
	l.Cache[k] = node
	l.insert(node, &l.root)
}
func (l *LRU) insert(node, root *Node1) {
	node.prev = root
	node.next = root.next
	node.prev.next = node
	node.next.prev = node
	l.Length++
	node.list = l
}
func (l *LRU) delete(node *Node1) {
	node.prev.next = node.next
	node.next.prev = node.prev
	l.Length--
}

// PrintlnDList 打印双链表的所有元素
func (l *LRU) PrintlnDList1() {
	if l.Length == 0 {
	}
	prev := l.root.next
	index := 0
	for prev.Value != nil {
		fmt.Printf("下标index: %d 元素 Element: %d\n", index, prev.Value)
		prev = prev.next
		index++
	}
	return
}
func main() {
	a := NewLRU()
	a.Append("1", 1)
	a.Append("2", 2)
	a.Append("3", 3)
	a.Append("4", 4)
	a.Append("5", 5)
	a.Append("5", 5)
	a.Append("5", 5)
	a.Append("5", 5)
	a.PrintlnDList1()
}
