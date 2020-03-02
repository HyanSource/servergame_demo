package api

import (
	"fmt"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/golang/protobuf/proto"
	"github.com/hyansource/servergame_demo/core"
	"github.com/hyansource/servergame_demo/pb"
)

//normalgame 路由
type NormalGameApi struct {
	znet.BaseRouter
}

func (*NormalGameApi) Handle(request ziface.IRequest) {

	msg := &pb.Bet{}
	//將客戶端傳來的protobuf解碼
	err := proto.Unmarshal(request.GetData(), msg)

	if err != nil {
		fmt.Println("NormalGame Unmarshal err:", err)
		return
	}

	//得知當前消息是從哪個玩家傳送來的
	pid, err := request.GetConnection().GetProperty("id")

	if err != nil {
		fmt.Println("GetProperty err:", err)
		request.GetConnection().Stop()
		return
	}
	//取得pid
	player := core.PlayerManageInstance.GetPlayerData(pid.(int32))
	//讓player發起請求
	player.PlayNormalGame(msg.Bet)
}
