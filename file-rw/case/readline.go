package _case

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// 以README文件为读取文件

const README = "/Users/a1-6/GolandProjects/file-rw/README.md"

// 一次性读取文件
// 按行拆分并打印
// 适合小文件读取

func ReadLine() {
	fileHander, err := os.OpenFile(README, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err)
	}
	defer fileHander.Close()
	bytes, err := io.ReadAll(fileHander)
	if err != nil {
		log.Fatal(err)
	}
	fileHander.Read(bytes)
	list := strings.Split(string(bytes), "\n")
	for _, l := range list {
		fmt.Println(l)
	}
}

// 通过bufio按行读取
// bufio通过io模块的封装，提供了数据的缓冲功能，能够一定程度上减少数据读写带来的缺陷
// 当发起读写操作时，会尝试从缓冲区读取数据，缓冲区没有数据，才会从数据中读取
// 缓冲区大小默认为4k

func Readline2() {
	fileHandle, err := os.OpenFile(README, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandle.Close()
	reader := bufio.NewReader(fileHandle)

	// 只读一行
	//line, _, err := reader.ReadLine()
	//if err != nil {
	//	return
	//}
	//fmt.Println(string(line))

	for {
		// 会从输入中读取数据直到遇到换行符（\n），然后返回包含已读取数据的字符串
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}
}

// 通过scanner 按行读取
// 单行默认大小为64k
// 使用 bufio.NewScanner(fileHandle) 可以方便地遍历文件的每一行文本，而无需显式处理换行符 \n

func ReadLine3() {
	fileHandle, err := os.OpenFile(README, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandle.Close()
	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
