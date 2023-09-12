package main

import _case "demo/file-rw/case"

func main() {
	//_ = _case.GetFiles("/Users/a1-6/GolandProjects/file-rw")
	//_case.CopyDirToDir()
	//_case.ReadWriteFiles()
	// 边读边写
	//_case.OneSideReadWriteToDest()
	// 一次性读取， 按行打印, 无缓冲
	//_case.ReadLine()
	// 一次性读取， 按行打印, 有缓冲
	//_case.Readline2()
	// 按行读取，不用ReadString('\n')，换行符号"\n"
	_case.ReadLine3()
}
