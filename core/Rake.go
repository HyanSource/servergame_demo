package core

import (
	"sync"
)

type RakeObj struct {
	//水庫總金額
	Pool float32
	//莊家利潤
	Profit float32
	//全局鎖
	Rake_Mutex sync.Mutex
}

const (
	//抽水比例
	Rate float32 = 0.05
)

var (
	//全局變數去呼叫使用
	RakeInstance *RakeObj
)

func init() {
	//初始化全局 未來可加入資料庫
	RakeInstance = &RakeObj{
		Pool:   0,
		Profit: 0,
	}

}

//綁定抽水方法
func (t *RakeObj) IncreasePool(raise float32) {
	t.Rake_Mutex.Lock()
	defer t.Rake_Mutex.Unlock()

	t.Pool += raise * (1 - Rate)
	t.Profit += raise * Rate
}

//綁定賠玩家方法
func (t *RakeObj) DeIncreasePool(getmoney float32) {
	t.Rake_Mutex.Lock()
	defer t.Rake_Mutex.Unlock()

	t.Pool -= getmoney
}

//綁定是否夠賠方法 (T)夠賠 (F)不夠賠
func (t *RakeObj) CheckPool(getmoney float32) bool {
	t.Rake_Mutex.Lock()
	defer t.Rake_Mutex.Unlock()

	return getmoney <= t.Pool
}
