package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func min(x, y int32) int32 {
	if x < y {
		return x
	} else if y < x {
		return y
	} else {
		return x
	}
}

/*
 * Complete the pageCount function below.
 */
func pageCount(n int32, p int32) int32 {
	/*
	 * Write your code here.
	 */
	var turnFront, turnBack int32
	if p%2 == 1 {
		turnFront = (p - 1) / 2
	} else {
		turnFront = p / 2
	}

	if n%2 == 1 {
		turnBack = (n - p) / 2
	} else {
		n = n + 1
		turnBack = (n - p) / 2
	}

	return min(turnFront, turnBack)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	pTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	p := int32(pTemp)

	result := pageCount(n, p)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
