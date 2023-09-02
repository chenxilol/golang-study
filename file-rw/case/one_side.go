package _case

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

// OneSideReadWriteToDest 边读边写
func OneSideReadWriteToDest() {
	dir := GetFiles(sourceDir)
	for _, s := range dir {
		_, dirName := path.Split(s)
		fileName := destDir + "one-side/" + dirName
		Open(s, fileName)
	}
}
func Open(src, dst string) {
	read, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer read.Close()
	file, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// 每次读取1024个字节
	bytes := make([]byte, 1024)
	for {
		n, err := read.Read(bytes)
		if err != nil {
			return
		}
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}
		file.Write(bytes[:n])
	}
}
