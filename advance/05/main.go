package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the morganAndString function below.
func morganAndString(a string, b string) string {
	var i, j, k int
	var pickA bool
	var whenSmaller byte
	res := make([]byte, len(a)+len(b))
	for i < len(a) && j < len(b) {
		var pick byte
		if a[i] < b[j] {
			pick = a[i]
			i++
		} else if a[i] > b[j] {
			pick = b[j]
			j++
		} else {
			if a[i] >= whenSmaller {
				ii, jj := i, j
				for ii < len(a) && jj < len(b) && a[ii] == b[jj] {
					ii++
					jj++
				}
				if ii == len(a) && jj == len(b) {
					pickA = true
					whenSmaller = a[ii-1]
				} else if ii == len(a) {
					pickA = false
					whenSmaller = b[jj]
				} else if jj == len(b) {
					pickA = true
					whenSmaller = a[ii]
				} else {
					pickA = a[ii] < b[jj]
					if pickA {
						whenSmaller = a[ii]
					} else {
						whenSmaller = b[jj]
					}
				}
			}
			if pickA {
				pick = a[i]
				i++
			} else {
				pick = b[j]
				j++
			}
			if pick >= whenSmaller {
				whenSmaller = 'A' - 1
			}
		}
		res[k] = pick
		k++
		if i < len(a) {
			copy(res[k:], a[i:])
		} else if j < len(b) {
			copy(res[k:], b[j:])
		}
	}
	return string(res)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		a := readLine(reader)

		b := readLine(reader)

		result := morganAndString(a, b)

		fmt.Fprintf(writer, "%s\n", result)
	}

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
