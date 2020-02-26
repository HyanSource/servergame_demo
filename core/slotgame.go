package core

import (
	"fmt"
)

//賠率物件
type Symbol struct {
	Id   int    //物件號碼
	Name string //物件名稱
	Odds int    //賠率
}

//得獎物件
type Lott struct {
	TableSymbol [9]int      //盤面元素ID
	Lines       map[int]int //key:中獎線,value:賠率
	GetOdds     int         //總賠率
}

var (
	AllSymbol []Symbol //全局賠率用物件

)

func init() {
	//清空元素
	AllSymbol = AllSymbol[:0]
	//放入物件
	AllSymbol = append(AllSymbol, Symbol{0, "櫻桃", 5})
	AllSymbol = append(AllSymbol, Symbol{1, "橘子", 10})
	AllSymbol = append(AllSymbol, Symbol{2, "葡萄", 15})
	AllSymbol = append(AllSymbol, Symbol{3, "西瓜", 20})
	AllSymbol = append(AllSymbol, Symbol{4, "銅鐘", 25})
	AllSymbol = append(AllSymbol, Symbol{5, "星星", 40})
	AllSymbol = append(AllSymbol, Symbol{6, "BAR", 50})
	AllSymbol = append(AllSymbol, Symbol{7, "雙色7", 60})
	AllSymbol = append(AllSymbol, Symbol{8, "紅色7", 80})
	AllSymbol = append(AllSymbol, Symbol{9, "藍色7", 100})

	fmt.Println("載入成功")
}

func GetLott() *Lott {
	//第二種想法是有全盤

	//或是可以拿之前slot_rtp的滾輪條直接用 再附上excel

	//直接改成15輪 (文房四寶)

	return &Lott{}
}
