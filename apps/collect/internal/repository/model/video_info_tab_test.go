package model

import (
	"fmt"
	"testing"
)

func TestVideoInfoTab(t *testing.T) {
	row := VideoInfoTab{
		Bvid:  "asfdsdf",
		Title: "xxxx"}
	res := fmt.Sprintf("%+v", row)
	fmt.Println(res)

	totalCount := 0
	defer func() {
		fmt.Println(totalCount)
	}()

	totalCount, x := 123, 234
	fmt.Println("x=", x)

}
