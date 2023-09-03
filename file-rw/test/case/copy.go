package _case

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

func CopyDirToDir() {
	dir := GetFileDir1(sourceDir)

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

func CopyDirToDir1() {
	list := GetFileDir1(sourceDir)
	for _, filename := range list {
		_, file := path.Split(filename)
		destName := destDir + "copy/" + file
		CopyFiles2(filename, destName)
	}
}

// CopyFiles1 用readFile he writeFile 进行一次性读写
func CopyFiles1(src, destName string) {
	file, err := os.ReadFile(src)
	if err != nil {
		return
	}
	err = os.WriteFile(destName, file, 0644)
	if err != nil {
		return
	}

}

// CopyFiles2 提供了更多的控制和灵活性，适用于更复杂的文件操作，他可以灵活的对文件的权限进行操作

func CopyFiles2(src, destName string) {
	open, err := os.Open(src)
	if err != nil {
		return
	}
	defer func(open *os.File) {
		err := open.Close()
		if err != nil {
		}
	}(open)
	file, err := os.OpenFile(destName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	// 这里直接全部复制io流，也可以编写边读
	//io.Copy(file, open)
	data := make([]byte, 1024)
	for {
		n, err := open.Read(data)
		if err != nil && err != io.EOF {
			log.Fatal(err)
			return
		}
		if n == 0 {
			break
		}
		file.Write(data[:n])

	}
}
