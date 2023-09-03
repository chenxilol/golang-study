package _case

import (
	"fmt"
	"os"
)

const sourceDir = "/Users/a1-6/GolandProjects/file-rw/soucre-file/"
const destDir = "/Users/a1-6/GolandProjects/file-rw/dest-file/"

func GetFileDir1(src string) []string {

	// 源文件目录
	dir, err := os.ReadDir(src)
	if err != nil {
		return nil
	}
	dirList := make([]string, 0)
	for _, entry := range dir {
		if entry.IsDir() {
			return nil
		}
		fileName := sourceDir + entry.Name()
		fmt.Println(fileName)
		dirList = append(dirList, fileName)
	}
	return dirList
}
