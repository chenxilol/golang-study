# 应用场景
1. 上传下载文件
2. 大文件分片传输
3. 文件移动
4. 文件内容按行获取

# 文件读写
1. 文件复制
2. 一次性读取文件内容并写入到新文件
3. 分片读取文件内容分步写入新文件
4. 文件按行读取

# 常见的文件操作函数的区别

1. ReadFile：

   ReadFile 是 io 包提供的一个函数，用于读取整个文件的内容。
它的参数是文件路径，它会打开、读取、关闭文件，并将文件的内容以字节切片（[]byte）的形式返回。
ReadFile 是一个高级别的函数，对于简单的文件读取操作非常方便，但不够灵活。

2. OpenFile：

   OpenFile 是 os 包提供的函数，用于打开文件并返回一个文件对象。
它的参数包括文件路径、打开模式（例如，读取、写入、追加等）、权限等。
打开文件后，你可以使用返回的文件对象执行不同的文件操作，例如读取、写入、关闭等。
OpenFile 提供了更多的控制和灵活性，适用于更复杂的文件操作。
Read：

3. Read 用于从文件对象中读取数据的方法，通常与 os.File 类型的文件对象一起使用。
它需要一个字节切片作为参数，并将读取的数据填充到该字节切片中。
Read 方法通常用于多次读取文件的一部分内容，直到到达文件的末尾。
你需要在循环中多次调用 Read 来读取文件的全部内容。


