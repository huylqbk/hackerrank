package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func multiples(n int64) {
	var three, five, fifteen int64
	three = (n - 1) / 3
	five = (n - 1) / 5
	fifteen = (n - 1) / 15
	fmt.Println(3*(three*(three+1)/2) + 5*(five*(five+1)/2) - 15*(fifteen*(fifteen+1)/2))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	st := strings.Split(readLine(reader), " ")
	sTemp, err := strconv.ParseInt(st[0], 10, 64)
	checkError(err)
	s := int32(sTemp)
	for i := 0; i < int(s); i++ {
		ns := strings.Split(readLine(reader), " ")
		nTemp, err := strconv.ParseInt(ns[0], 10, 64)
		checkError(err)
		multiples(nTemp)
	}
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
