package main

import (
	"fmt"
	"io"
	"os"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout

func main() {
	var n, x int
	fmt.Fscan(stdin, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(stdin, &x)
		arr[i] = x
	}
	fmt.Fprintln(stdout, n, arr)
}
