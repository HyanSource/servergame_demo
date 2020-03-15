package core

//
type IMode interface {
	//判斷盤面有幾個Wild
	WildCount([15]int) int
}
