package _case

import (
	"fmt"
	"os"
	"path"
)

func ReadWriteFiles() {
	list := GetFiles(sourceDir)
	for _, f := range list {
		bytes, err := os.ReadFile(f)
		if err != nil {
			return
		}
		_, fileName := path.Split(f)
		normalFile := destDir + "normal/" + fileName
		// 如果文件不存在就会重新创建新的文件
		err = os.WriteFile(normalFile, bytes, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
