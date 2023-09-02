package _case

import (
	"log"
	"os"
)

const sourceDir = "/Users/a1-6/GolandProjects/file-rw/soucre-file/"
const destDir = "/Users/a1-6/GolandProjects/file-rw/dest-file/"

func GetFileDir(src string) []string {

	// 源文件目录

	dir, err := os.ReadDir(src)
	if err != nil {
		log.Fatal("读取文件失败", err)
		return nil
	}
	list := make([]string, 0)
	for _, f := range dir {
		if f.IsDir() {
			continue
		}
		dir := src + "/" + f.Name()
		list = append(list, dir)
	}
	return list
}
