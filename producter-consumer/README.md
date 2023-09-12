# chan 底层原理
    qcount   uint           // chan 中存在元素的数量
	dataqsiz uint           // chan 中的元素容量
	buf      unsafe.Pointer // chan 中的元素队列（环形数组）
	elemsize uint16         // chan 元素类型大小
	closed   uint32         // 标识chan是否关闭
	elemtype *_type         // chan 元素类型
	sendx    uint           // 写入元素的 index
	recvx    uint           // 读取元素的 index
	recvq    waitq          // 阻塞的读协程队列
	sendq    waitq          // 锁