# Context 是什么

1. 译为上下文　　
2. 用于进程之间信息和信号的传递
3. 用于服务之间的信号和信息从传递
4. 
## Context 的功能
1. Context 可以用于不同的api或者进程之间传递(携带键值对传递)
2. 传递取消信号(主动取消，超时/时限取消)，因为Context是树结构，所以传递是单向传递的，只有父节点取消的时候，才会把取消的信号传递给父节点的衍生子节点
3. 
## 应用场景
1. 用于父子协程间取消信号传递
2. 用于客服端与服务器之间的信息传递
3. 用于设置请求超时时间等

## 数据结构

```
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key any) any
}
```

# 知识点
1. context.WithTimeout(ctx, time.Second*10)
   他用于创建一个具有指定超时时间的新上下文，上面就代表过10秒后，所有的操作都将会关闭
2. context.WithValue(ctx, "desc", "ContextCase")
   这个用于创建一个_