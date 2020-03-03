package core

import (
	"fmt"
	"sync"
	"zinx/ziface"

	"github.com/golang/protobuf/proto"
	"github.com/hyansource/servergame_demo/pb" //pb
)

//玩家資訊的結構體
type Player struct {
	PlayerID    int32
	Conn        ziface.IConnection
	PlayerName  string
	PlayerMoney int32
}

//全局ID
var (
	IDGen      int32 = 1
	IDGen_Lock sync.Mutex
)

//創建一個玩家
func NewPlayer(conn ziface.IConnection) *Player {

	//生成ID

	IDGen_Lock.Lock()
	id := IDGen
	IDGen++
	defer IDGen_Lock.Unlock()

	p := &Player{
		PlayerID:    id,
		Conn:        conn,
		PlayerName:  "",
		PlayerMoney: 10000,
	}

	return p
}

//玩家上線事件
func (p *Player) OnLineEvent() {
	//proto數據
	data := &pb.PlayerData{
		PlayerID:    p.PlayerID,
		PlayerName:  p.PlayerName,
		PlayerMoney: p.PlayerMoney,
	}

	//給客戶端訊息
	p.SendMsg(1, data)
}

//玩家下線事件
func (p *Player) OffLineEvent() {

	//封裝消息
	msg := &pb.PlayerData{
		PlayerID:    p.PlayerID,
		PlayerName:  "",
		PlayerMoney: 0,
	}

	//給所有玩家傳達離線的訊息
	for _, player := range PlayerManageInstance.AllPlayer {
		player.SendMsg(201, msg)
	}

}

//傳送訊息
func (p *Player) SendMsg(msgID uint32, data proto.Message) {

	//將proto message結構體序列化

	marshalmsg, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	if p.Conn == nil {
		fmt.Println("player connection nil")
		return
	}

	if err := p.Conn.SendMsg(msgID, marshalmsg); err != nil {
		fmt.Println("Player SendMsg err:", err)
		return
	}

}

//遊玩一次normalgame
func (p *Player) PlayNormalGame(bet int32) {

	if (p.PlayerMoney - bet) < 0 {
		fmt.Println("playermoney<=0")
		return
	}

	if bet == 0 {
		fmt.Println("bet==0")
		return
	}
	//取得盤面
	t := NewGameResult(RandomGet())
	//計算金額(需要加上mutex)
	p.PlayerMoney -= bet
	p.PlayerMoney += bet * t.AllOdds
	//封裝訊息
	msg := &pb.ReturnGameResult{
		AllMoney:   p.PlayerMoney,
		GetMoney:   bet * t.AllOdds,
		GameResult: t,
	}

	p.SendMsg(100, msg)

}

//廣播給所有人
func (p *Player) AllBroadCast(talkmsg string) {

	msg := &pb.BroadCast{
		PlayerID: p.PlayerID,
		TalkMsg:  &pb.TalkMsg{Content: talkmsg},
	}

	GetPlayers := PlayerManageInstance.GetAllPlayerData()

	for _, v := range GetPlayers {
		v.SendMsg(150, msg)
	}

}
