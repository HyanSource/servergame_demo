package core

import (
	"fmt"
	"testing"
)

func TestReelsResult(t *testing.T) {

	a := NewGameResult(SpecificGet([5]int{4, 5, 6, 7, 8}))

	fmt.Println(a)
}

func TestRandomResult(t *testing.T) {
	a := NewGameResult(RandomGet(0))
	fmt.Println(a.Result)
}

func TestAllResult(t *testing.T) {

}
