package main

import (
	"fmt"

	//"zinx/ziface"
	//"zinx/znet"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"

	"github.com/hyansource/servergame_demo/core"
)

func OnConnectionAdd(conn ziface.IConnection) {

	//創建玩家
	player := core.NewPlayer(conn)

	//玩家新增進去Manage中
	core.PlayerManageInstance.AddPlayer(player)

	//將該連接綁定屬性id
	conn.SetProperty("id", player.PlayerID)

	fmt.Println("client success")
}

func OnConnectionLost(conn ziface.IConnection) {

	//獲取當前連接的id屬性
	id, _ := conn.GetProperty("id")

	//根據當前連接獲取玩家對象
	player := core.PlayerManageInstance.GetPlayerData(id.(int32))

	if id != nil {
		player.OffLineEvent()
		core.PlayerManageInstance.RemovePlayer(id.(int32))
	}

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
