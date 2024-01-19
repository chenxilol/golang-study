package main

import (
	"fmt"
	"github.com/google/martian/log"
)

// Element 创建双链表的节点
type Element struct {
	next, prev *Element
	Value      any
	list       *DList
}
type DList struct {
	Length uint
	root   Element
}

func NewDList() *DList {
	return new(DList).Init()
}
func (l *DList) Init() *DList {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.Length = 0
	return l
}
func (l *DList) insert(e, root *Element) {
	// root 为per的地址，刚开始一直都是初始化的地址,我们要把最后一个的pre的地址加入到，新插入的pre中
	e.prev = root
	// 把刚开始的初始化的节点传给新节点
	e.next = root.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.Length++
}

// PushBack 后插
func (l *DList) PushBack(v any) {
	l.insert(&Element{Value: v}, l.root.prev)
}

// FrontBack 前插
func (l *DList) FrontBack(v any) {
	l.insert(&Element{Value: v}, &l.root)
}

// Front 返回双链表的第一个元素地址
func (l *DList) Front() *Element {
	if l.Length == 0 {
		return nil
	}
	return l.root.next
}

// Back 返回双链表的做后一个元素地址
func (l *DList) Back() *Element {
	if l.Length == 0 {
		return nil
	}
	return l.root.prev
}

// PrintlnDList 打印双链表的所有元素
func (l *DList) PrintlnDList() {
	if l.Length == 0 {
		log.Infof("双链表中没有元素")
	}
	prev := l.root.next
	index := 0
	for prev.Value != nil {
		fmt.Printf("下标index: %d 元素 Element: %s\n", index, prev.Value)
		prev = prev.next
		index++
	}
	return
}

// RemoveElement 根据元素删除
func (l *DList) RemoveElement(v any) bool {
	prev := l.root.next
	count := 0
	index := 0
	for prev.Value != nil {
		if prev.Value == v {
			prev.prev.next = prev.next
			prev.next.prev = prev.prev
			l.Length--
			count++
			fmt.Printf("删除的元素Info 下标index: %d 元素 Element: %s\n", index, prev.Value)
		}
		prev = prev.next
		index++
	}
	if count == 0 {
		fmt.Println("链表中没有该元素")
	}
	return true
}
func main() {
	ab := NewDList()
	ab.PushBack("sdfd")
	ab.PushBack("sd")
	ab.PushBack("sad")
	ab.PushBack("sdfd")
	ab.PushBack("sdfd")
	ab.PushBack("sdfd")
	ab.PrintlnDList()
	ab.RemoveElement("sdfd")
	ab.PushBack("sdfd")
	ab.PrintlnDList()
}
