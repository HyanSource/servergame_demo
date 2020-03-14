package core

import (
	"fmt"
	"testing"
)

func TestReelsResult(t *testing.T) {

	a := NewGameResult(SpecificGet([5]int{4, 5, 6, 7, 8}), 0)

	fmt.Println(a)
}

func TestRandomResult(t *testing.T) {
	a := NewGameResult(RandomGet(0), 0)
	fmt.Println(a.Result)
}

//專門對每個回合做測試
func TestFreeCount(t *testing.T) {
	for {
		a := NewGameResult(RandomGet(0), 0)
		fmt.Println(a.Result)
		fmt.Println(a.ScatterCount)
		if a.ScatterCount > 0 {
			break
		}
	}
}
