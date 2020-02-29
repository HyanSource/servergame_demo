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
	a := NewGameResult(RandomGet())
	fmt.Println(a)
}

func TestAllResult(t *testing.T) {
	// for i := 0; i < 50; i++ {
	// 	for j := 0; j < 50; j++ {
	// 		for k := 0; k < 50; k++ {
	// 			for l := 0; l < 50; l++ {
	// 				for m := 0; m < 50; m++ {

	// 					a := NewGameResult(SpecificGet([5]int{i, j, k, l, m}))
	// 					fmt.Println(a)
	// 				}
	// 			}
	// 		}
	// 	}
	// }
}
