package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

func OnConnectionAdd(conn ziface.IConnection) {

	fmt.Println("client success")
}

func OnConnectionLost(conn ziface.IConnection) {

	fmt.Println("client lost")
}

func main() {

	ser := znet.NewServer()

	//註冊客戶端連接和丟失函數
	ser.SetOnConnStart(OnConnectionAdd)
	ser.SetOnConnStop(OnConnectionLost)

	//註冊路由
	//ser.AddRouter()
	//ser.AddRouter()

	ser.Serve()

	fmt.Println("succees")
}
