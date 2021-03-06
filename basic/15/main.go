package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

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

func minus(x, y int32) int32 {
	if x > y {
		return x - y
	} else {
		return y - x
	}
}

// Complete the beautifulDays function below.
func beautifulDays(i int32, j int32, k int32) int32 {
	count := int32(0)
	for a := i; a <= j; a++ {
		if minus(a, revertNum(a))%k == 0 {
			count++
		}
	}
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	ijk := strings.Split(readLine(reader), " ")

	iTemp, err := strconv.ParseInt(ijk[0], 10, 64)
	checkError(err)
	i := int32(iTemp)

	jTemp, err := strconv.ParseInt(ijk[1], 10, 64)
	checkError(err)
	j := int32(jTemp)

	kTemp, err := strconv.ParseInt(ijk[2], 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := beautifulDays(i, j, k)

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
