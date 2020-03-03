package api

import (
	"fmt"

	"github.com/hyansource/servergame_demo/core"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/golang/protobuf/proto"
	"github.com/hyansource/servergame_demo/pb"
)

type BroadCastApi struct {
	znet.BaseRouter
}

func (*BroadCastApi) Handle(request ziface.IRequest) {
	msg := &pb.BroadCast{}

	err := proto.Unmarshal(request.GetData(), msg)

	if err != nil {
		fmt.Println("BroadCast Unmarshal err:", err)
		return
	}

	pid, err := request.GetConnection().GetProperty("id")

	if err != nil {
		fmt.Println("GetProperty err:", err)
		request.GetConnection().Stop()
		return
	}

	player := core.PlayerManageInstance.GetPlayerData(pid.(int32))

	player.AllBroadCast(msg.TalkMsg.Content)
}
