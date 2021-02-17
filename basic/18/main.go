package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the appendAndDelete function below.
func appendAndDelete(s string, t string, k int32) string {

	sLength := len(s)
	tLength := len(t)

	if sLength < tLength {
		if (tLength-sLength)%2 == 0 {
			return "Yes"
		} else {
			return "No"
		}
	}

	commonPrefix := 0

	minLen := 0
	if sLength > tLength {
		minLen = tLength
	} else {
		minLen = sLength
	}

	// find a largest common prefix between s and t
	for i := 0; i < minLen; i++ {
		if s[i] != t[i] {
			break
		}
		commonPrefix++
	}

	needToProcess := (tLength - commonPrefix) + (sLength - commonPrefix)
	if int32(needToProcess) <= k {
		return "Yes"
	} else {
		return "No"
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	t := readLine(reader)

	kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := appendAndDelete(s, t, k)

	fmt.Fprintf(writer, "%s\n", result)

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
