package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 1. 使用 rpc.Dial 连接服务器
	conn, err1 := rpc.Dial("tcp", "127.0.0.1:8080")
	if err1 != nil {
		fmt.Println("连接失败:", err1)
	}
	// 2.退出时断开连接
	defer conn.Close()

	// 3. 调用远程函数的方法
	var reply string
	// 第一个参数 Hello.SayHello 是服务名.方法名
	// 第二个参数 我是客户端 是传入req的参数
	// 第三个参数 &reply 是返回的参数
	err2 := conn.Call("Hello.SayHello", "我是客户端", &reply)
	if err2 != nil {
		fmt.Println("调用失败:", err2)
	}
	// 4. 获取微服务返回的数据
	fmt.Println(reply)
}
