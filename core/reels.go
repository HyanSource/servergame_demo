package core

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/hyansource/servergame_demo/pb"
)

//https://online808.com/other/1305

//存放符號和賠率
//(0)10號 (1)J (2)Q (3)K (4)A
//(5)硯台 (6)紙 (7)墨 (8)毛筆
//(9)百搭 (10)免費

//盤面排序方式
//0  1  2  3  4
//5  6  7  8  9
//10 11 12 13 14

//規則
//百搭 出現在 2 3 4 5滾輪
//Scatter 出現在 2 3 4滾輪

//再附上excel

const (
	//wild的號碼
	WildID = 9
	//scatter的號碼
	FreeID = 10
)

var (
	//賠率
	Odds [11][3]int = [11][3]int{{5, 20, 100}, {5, 25, 125}, {5, 25, 125}, {10, 25, 125}, {10, 25, 125}, {20, 100, 300}, {20, 100, 300}, {25, 150, 500}, {50, 250, 1000}, {0, 0, 0}, {0, 0, 0}}
	//連線位置
	WinLines [20][5]int = [20][5]int{{5, 6, 7, 8, 9}, {0, 1, 2, 3, 4}, {10, 11, 12, 13, 14}, {0, 6, 12, 8, 4}, {10, 6, 2, 8, 14}, {0, 1, 7, 13, 14}, {10, 11, 7, 3, 4}, {5, 1, 7, 13, 9}, {5, 11, 7, 3, 9}, {0, 6, 7, 8, 14}, {10, 6, 7, 8, 4}, {5, 1, 2, 8, 14}, {5, 11, 12, 8, 4}, {5, 6, 2, 8, 14}, {5, 6, 12, 8, 4}, {0, 1, 7, 13, 9}, {10, 11, 7, 3, 9}, {5, 1, 7, 13, 14}, {5, 11, 7, 3, 4}, {0, 1, 2, 8, 14}}
)

//取得隨機的盤面
func RandomGet(gameround int32) [15]int {

	ReturnLott := [15]int{}

	for i := 0; i < 5; i++ {

		reelmax := len(AllModeReels[gameround].Reels[i]) - 1
		r := rand.Intn(reelmax)

		ReturnLott[10+i] = AllModeReels[gameround].Reels[i][r]
		ReturnLott[5+i] = AllModeReels[gameround].Reels[i][GetReelIndex(r, reelmax)]
		ReturnLott[i] = AllModeReels[gameround].Reels[i][GetReelIndex(GetReelIndex(r, reelmax), reelmax)]

	}

	return ReturnLott
}

//取得特定盤面 1,10,15,50,50
func SpecificGet(ID [5]int) [15]int {

	ReturnLott := [15]int{}

	for i := 0; i < 5; i++ {
		if ID[i] >= 50 || ID[i] < 0 {
			panic("取得特定盤面錯誤:" + strconv.Itoa(ID[i]))
		}

		reelmax := len(AllModeReels[0].Reels[i]) - 1

		ReturnLott[10+i] = AllModeReels[0].Reels[i][ID[i]]
		ReturnLott[5+i] = AllModeReels[0].Reels[i][GetReelIndex(ID[i], reelmax)]
		ReturnLott[i] = AllModeReels[0].Reels[i][GetReelIndex(GetReelIndex(ID[i], reelmax), reelmax)]
	}

	return ReturnLott
}

//得到下一個滾輪號
func GetReelIndex(NowIndex int, reelmax int) int {
	if NowIndex >= reelmax {
		return 0
	}
	return NowIndex + 1
}

//新增一個遊戲結果的物件
func NewGameResult(Lott [15]int, gameround int32) *pb.GameResult {

	//儲存每個得獎
	alllineodds := make([]*pb.LineOdds, 0)
	scattercount := 0
	allodds := 0

	//計算連線和賠率
	for k, v := range WinLines {

		//連線位置第一元素
		firstsymbol := Lott[v[0]]

		if (firstsymbol == Lott[v[1]] || Lott[v[1]] == WildID) &&
			(firstsymbol == Lott[v[2]] || Lott[v[2]] == WildID) &&
			(firstsymbol == Lott[v[3]] || Lott[v[3]] == WildID) &&
			(firstsymbol == Lott[v[4]] || Lott[v[4]] == WildID) {

			alllineodds = append(alllineodds, &pb.LineOdds{
				GetID:    int32(k),
				GetCount: 5,
				GetOdds:  int32(Odds[firstsymbol][2]),
			})
			allodds += Odds[firstsymbol][2]

		} else if (firstsymbol == Lott[v[1]] || Lott[v[1]] == WildID) &&
			(firstsymbol == Lott[v[2]] || Lott[v[2]] == WildID) &&
			(firstsymbol == Lott[v[3]] || Lott[v[3]] == WildID) {

			alllineodds = append(alllineodds, &pb.LineOdds{
				GetID:    int32(k),
				GetCount: 4,
				GetOdds:  int32(Odds[firstsymbol][1]),
			})
			allodds += Odds[firstsymbol][1]

		} else if (firstsymbol == Lott[v[1]] || Lott[v[1]] == WildID) &&
			(firstsymbol == Lott[v[2]] || Lott[v[2]] == WildID) {

			alllineodds = append(alllineodds, &pb.LineOdds{
				GetID:    int32(k),
				GetCount: 3,
				GetOdds:  int32(Odds[firstsymbol][0]),
			})
			allodds += Odds[firstsymbol][0]
		}

	}

	//計算免費個數 這段應該可以拿掉

	for i := range Lott {
		if Lott[i] == FreeID {
			scattercount++
		}
	}

	return &pb.GameResult{
		Result:       ArrayToString(Lott, ","),
		LinesOdds:    alllineodds,
		ScatterCount: int32(scattercount),
		AllOdds:      int32(allodds),
	}
}

//拚接
func ArrayToString(t interface{}, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(t), " ", delim, -1), "[]")
}
