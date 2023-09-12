package main

import (
	"log"
	"strings"
)

// Trie 前缀树结构
type Trie struct {
	next   map[string]*Trie
	isWord bool
}

func NewTrie() Trie {
	root := new(Trie)
	root.next = make(map[string]*Trie)
	root.isWord = false
	return *root
}

// Insert 插入数据， 路由根据 "/" 进行拆分
func (t *Trie) Insert(word string) {
	for _, v := range strings.Split(word, "/") {
		if t.next[v] == nil {
			node := new(Trie)
			node.next = make(map[string]*Trie)
			node.isWord = false
			t.next[v] = node
		}
		// * 匹配所有
		// {X}  匹配路由参数 X
		if v == "*" || strings.Index(v, "{") != -1 {
			t.isWord = true
		}
		t = t.next[v]
	}
	t.isWord = true
}

// Search 匹配路由
func (t *Trie) Search(word string) (isHave bool, arg map[string]string) {
	arg = make(map[string]string)
	isHave = false
	for _, v := range strings.Split(word, "/") {
		if t.isWord {
			for k, _ := range t.next {
				if strings.Index(k, "{") != -1 {
					key := strings.Replace(k, "{", "", -1)
					key = strings.Replace(key, "}", "", -1)
					arg[key] = v
				}
				v = k
			}
			//log.Println("v = ", v)
		}
		if t.next[v] == nil {
			log.Println("找不到了, 匹配不上")
			return
		}
		t = t.next[v]
	}
	//log.Println(t.next, len(t.next))
	// 必须匹配全  比如: /v1/{b}/{a}  /v1/123匹配不到， /v1/123/456才可匹配
	if len(t.next) == 0 {
		isHave = t.isWord
		return
	}
	return
}

// Show 打印树
func (t *Trie) Show() {
Loop:
	if t == nil {
		log.Println("end")
		return
	}
	log.Println(t)
	for _, v := range t.next {
		t = v
		goto Loop
	}

}

func main() {

	// 1. 使用前缀树 匹配路由
	// 2. 提取路由上的参数规则:  定义路由/v1/{a},其中a就是参数名,请求路由/v1/123,那么a=123
	// 3. 可以定义多个参数, 如: 定义路由/v1/{a}/{b}/{c},其中就有 a,b,c 3个路由参数
	// 4. 必须匹配完,如: 定义路由/v1/{a}/{b}则/v1/123/456能匹配上，/v1/123不能匹配上
	// 5. 不允许的规则,如:  /v1/*  /v1/{a}  两个都不能用，相互冲突

	t := NewTrie()
	t.Insert("/")
	t.Insert("H/e/c/{b}/{a}")
	t.Insert("a/e/c/l/o")
	t.Insert("e/c/c/*")       // 任意的 无法取 *的值
	t.Insert("e/c/c/{a}/{a}") // 只取后一个a
	t.Insert("/v1/*")
	t.Insert("/v1/{a}")
	t.Show()
	//isok, cs := t.Search("H/e/c/asdasd/asdasd")
	//isok, cs := t.Search("/v1/123")
	isok, cs := t.Search("/")
	log.Println(isok, cs)

	//fmt.Print(t.Search("H/e/l/l/aaa"),"\n")

}
