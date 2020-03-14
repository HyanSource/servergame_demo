package core

import (
	"fmt"
	"testing"
)

//測試水庫初始化
func TestRakeInit(t *testing.T) {

	fmt.Println("RakeInstance.Pool:", RakeInstance.Pool)
	fmt.Println("RakeInstance.Profit:", RakeInstance.Profit)
	fmt.Println("OK")
}

func TestRakePool(t *testing.T) {

	RakeInstance.Pool = 1000

	RakeInstance.IncreasePool(200)

	fmt.Println("Pool:", RakeInstance.Pool, " Profit:", RakeInstance.Profit)

	fmt.Println("500:", RakeInstance.CheckPool(500))
	fmt.Println("1500:", RakeInstance.CheckPool(1500))

	RakeInstance.DeIncreasePool(1000)

	fmt.Println("Pool:", RakeInstance.Pool, " Profit:", RakeInstance.Profit)

}
