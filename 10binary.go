package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int64(nTemp)
	binary := strconv.FormatInt(n, 2)

	fmt.Println(binary)

	onecount := 0
	runningmax := 1
	prevone := false

	for _, v := range binary {
		if string(v) == "1" {
			if prevone == true {
				onecount++
				if onecount > runningmax {
					runningmax = onecount
				}
			} else {
				prevone = true
				onecount = 1
			}
		} else {
			prevone = false
		}

	}

	fmt.Println(runningmax)

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
