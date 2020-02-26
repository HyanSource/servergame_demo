package core

import "sync"

//
var (
	PlayerManageInstance *PlayerManage
)

//管理總玩家的結構體
type PlayerManage struct {
	AllPlayer   map[int32]*Player //當前玩家的集合
	PlayersLock sync.Mutex        //互斥鎖
}

//初始化的方法
func init() {
	PlayerManageInstance = &PlayerManage{
		AllPlayer: make(map[int32]*Player),
	}
}

//將玩家加入進管理的集合
func (pm *PlayerManage) AddPlayer(player *Player) {

	pm.PlayersLock.Lock()
	defer pm.PlayersLock.Unlock()

	pm.AllPlayer[player.PlayerID] = player
}

//刪除玩家資料
func (pm *PlayerManage) RemovePlayer(playerid int32) {
	pm.PlayersLock.Lock()
	defer pm.PlayersLock.Unlock()

	delete(pm.AllPlayer, playerid)
}

//取得玩家資料
func (pm *PlayerManage) GetPlayerData(playerid int32) *Player {
	pm.PlayersLock.Lock()
	defer pm.PlayersLock.Unlock()

	if playerdata, ok := pm.AllPlayer[playerid]; ok {
		return playerdata
	}

	return nil
}

//取得所有玩家資料
func (pm *PlayerManage) GetAllPlayerData() []*Player {

	pm.PlayersLock.Lock()
	defer pm.PlayersLock.Unlock()

	getallplayer := make([]*Player, 0)

	for _, v := range pm.AllPlayer {
		getallplayer = append(getallplayer, v)
	}

	return getallplayer
}
