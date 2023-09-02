package _case

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

func CopyDirToDir() {
	list := GetFiles(sourceDir)
	for _, f := range list {
		// 分割处文件的名字
		_, name := path.Split(f)
		destFileName := destDir + "copy/" + name
		_, err := COPYFile(f, destFileName)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
			return
		}
	}
}
func COPYFile(srcName, destName string) (int64, error) {
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer src.Close()
	dst, err := os.OpenFile(destName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
