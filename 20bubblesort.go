// https://www.hackerrank.com/challenges/ctci-bubble-sort/problem

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// CountSwaps implemements Bubblesort using Gayle McDowell's Java algorithm and adds a counter
func countSwaps(a []int32) {
	isSorted := false
	numSwaps := 0
	lastUnsorted := len(a) - 1
	for !isSorted {
		isSorted = true
		for i := 0; i < lastUnsorted; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				numSwaps++
				isSorted = false
			}
		}
		lastUnsorted--
	}

	fmt.Printf("Array is sorted in %d swaps.\n", numSwaps)
	fmt.Println("First Element:", a[0])

	fmt.Println("Last Element:", a[len(a)-1])
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(readLine(reader), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	countSwaps(a)
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
