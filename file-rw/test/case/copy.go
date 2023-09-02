package _case

import (
	"fmt"
	"io"
	"os"
	"path"
)

func CopyDirToDir() {
	dir := GetFileDir(sourceDir)
	for _, s := range dir {
		_, filelName := path.Split(s)
		distDirs := destDir + "copy/" + filelName
		CopyFiles(s, distDirs)

	}
}
func CopyFiles(src, dist string) (int64, error) {
	scrFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	}
	defer func(scrFile *os.File) {
		err := scrFile.Close()
		if err != nil {

		}
	}(scrFile)
	// 打开文件os.O_CREATE 如果目录没有要打开的文件，那么他会自动创建文件，
	// os.O_WRONLY 这个是允许写的的权限，他和open区别是openFile可以设置权限
	dst, err := os.OpenFile(dist, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
		}
	}(dst)
	return io.Copy(dst, scrFile)

}
