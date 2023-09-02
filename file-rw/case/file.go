package _case

import (
	"os"
)

// 源文件目录
const sourceDir = "/Users/a1-6/GolandProjects/file-rw/soucre-file/"
const destDir = "/Users/a1-6/GolandProjects/file-rw/dest-file/"

func GetFiles(dir string) []string {
	readDir, err := os.ReadDir(dir)
	if err != nil {
		panic("读取文件错误")
		return nil
	}
	list := make([]string, 0)
	for _, f := range readDir {
		if f.IsDir() {
			continue
		}
		// strings,trim() 在字符串的开头和结尾处，删除目标字符
		//s := strings.Trim(dir, "/") + "/" + f.Name()
		s := dir + "/" + f.Name()
		list = append(list, s)
	}
	return list
}
