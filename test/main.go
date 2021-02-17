package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(revertNum(int32(120)))
}

func revertNum(x int32) int32 {
	if x < 10 {
		return x
	}
	str := ""
	for x > 0 {
		str += fmt.Sprintf("%d", x%10)
		x = x / 10
	}
	num, _ := strconv.ParseInt(str, 10, 32)
	return int32(num)
}
