# day1-http完成情况
1. 主要完成了对框架的封装，已经路由拦截解析功能
2. 完成了GET，SET等方法的封装，定义了HandlerFunc func(http.ResponseWriter, *http.Request) 
   这个接口提供给框架用户的使用者，主要因为SeverHttp函数的参数为以上两个
# day2-context完成情况
1. 完成了对context上下文的封装，因为每次发送请求的时候，都要发送请求头，以及状态码，如果不封装将会大大提高了框架的难度
   因为是上下文所以它应该存储路由整个的信息。
   在封装context的同时，也修改了gee.go，HandlerFunc func(c *Context) 
2. 对gee.go 进行了代码提取，把关于路由方面的提取到router.go ，两个功能一个addRoute添加路由功能，一个路由查找功能