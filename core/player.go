package core

import (
	"fmt"
	"sync"

	"github.com/aceld/zinx/ziface"
	"github.com/golang/protobuf/proto"
	"github.com/hyansource/servergame_demo/pb" //pb
)

//玩家資訊的結構體
type Player struct {
	PlayerID    int32              //玩家ID
	Conn        ziface.IConnection //連線
	PlayerName  string             //玩家名字
	PlayerMoney int32              //玩家金錢
	FreeCount   int32              //免費次數
	FreeRound   int32              //免費回合
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
		PlayerName:  "Guest",
		PlayerMoney: 10000,
		FreeCount:   0,
		FreeRound:   0,
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

	t := &pb.GameResult{}

	if p.FreeCount <= 0 { //普通模式
		if (p.PlayerMoney - bet) < 0 {
			fmt.Println("playermoney<=0")
			return
		}

		if bet == 0 {
			fmt.Println("bet==0")
			return
		}
		p.PlayerMoney -= bet

		//取得盤面
		t = NewGameResult(RandomGet(0), 0)
		//計算金額(需要加上mutex)
		p.PlayerMoney += bet * t.AllOdds

		//獲得免費遊戲
		if t.ScatterCount == 3 {
			p.FreeCount = 10
		}
	} else { //特殊模式
		//取得盤面
		t = NewGameResult(RandomGet(p.FreeRound), p.FreeRound)

		//計算金額(需要加上mutex)
		p.PlayerMoney += bet * t.AllOdds
		//獲得免費遊戲(需要經過單元測試(testing)驗證)
		if t.ScatterCount > 0 {
			p.FreeCount += t.ScatterCount
		}
	}

	//封裝訊息
	msg := &pb.ReturnGameResult{
		AllMoney:   p.PlayerMoney,
		GetMoney:   bet * t.AllOdds,
		FreeCount:  p.FreeCount,
		FreeRound:  p.FreeRound,
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
