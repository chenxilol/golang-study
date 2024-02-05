package main

import (
	"fmt"
)

// LruNode 创建双链表的节点
type LruNode struct {
	next, prev *LruNode
	Value      any
	Key        string
	list       *LRU
}
type LRU struct {
	Length   uint
	maxBytes uint
	Cache    map[string]*LruNode
	root     LruNode
}

func NewLRU() *LRU {
	return new(LRU).Init()
}
func (l *LRU) Init() *LRU {
	l.root.next = &l.root
	l.maxBytes = 5
	l.root.prev = &l.root
	l.Cache = make(map[string]*LruNode)
	l.Length = 0
	return l
}

// FrontBack 前插
func (l *LRU) FrontBack(k string, v any) {
	if node, ok := l.Cache[k]; !ok {
		nodeNow := &LruNode{Value: v, Key: k}
		l.Cache[k] = nodeNow
		l.insert(nodeNow, &l.root)
		if l.Length+1 > l.maxBytes {
			delete(l.Cache, l.root.prev.Key)
			l.delete(l.root.prev)
		}
	} else {
		delete(l.Cache, k)
		nodeNow := &LruNode{Value: v, Key: k}
		l.Cache[k] = node
		l.insert(nodeNow, &l.root)
	}

}

// GetValue 查找k,v
func (l *LRU) GetValue(k string) (val any, ok bool) {
	if node, ok := l.Cache[k]; ok {
		l.delete(node)
		l.insert(node, &l.root)
		return node.Value, ok
	}
	return nil, ok
}
func (l *LRU) insert(node, root *LruNode) {
	node.prev = root
	node.next = root.next
	node.prev.next = node
	node.next.prev = node
	l.Length++
	node.list = l
}
func (l *LRU) delete(node *LruNode) {
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
	a.FrontBack("1", 1)
	a.FrontBack("2", 2)
	a.FrontBack("3", 3)
	a.FrontBack("4", 4)
	a.FrontBack("5", 5)
	a.FrontBack("5", 5)
	a.FrontBack("6", 6)
	a.FrontBack("7", 7)
	a.GetValue("4")
	a.PrintlnDList1()
}
