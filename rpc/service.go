package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// 定义远程调用方法
type Hello struct {
}

/*
* 1. 方法只能有两个可序列化的参数，其中第二个参数是指针类型
*   req: 表示获取客服端传过来的数据
*   res: 表示给客户端传递的数据
* 2. 方法要返回一个error类型，同时必须是公开的方法
* 3. req和res的类型不能是： channel（通道），func(函数) 均不能进行序列化
 */
func (h Hello) SayHello(req string, res *string) error {
	*res = "Hello11, " + req
	return nil
}

func main() {
	// 1. 注册服务
	err1 := rpc.RegisterName("Hello", new(Hello))
	if err1 != nil {
		fmt.Println("注册服务失败:", err1)
	}
	// 2. 监听端口
	listener, err2 := net.Listen("tcp", "127.0.0.1:8080")
	if err2 != nil {
		fmt.Println("监听端口失败:", err2)
	}
	//3. 应用退出时关闭监听
	defer listener.Close()

	for {
		fmt.Println("监听端口成功")
		conn, err3 := listener.Accept()
		if err3 != nil {
			fmt.Print(err3)
		}
		// 绑定
		rpc.ServeConn(conn)
	}

}
