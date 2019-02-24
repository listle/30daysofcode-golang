/**
Objective
Today, we're building on our knowledge of Arrays by adding another dimension. Check out the Tutorial tab for learning materials and an instructional video!

Context
Given a 2D Array,

1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 0 0 0 0 0
0 0 0 0 0 0
0 0 0 0 0 0

We define an hourglass to be a subset of values with indices falling in this pattern's graphical representation:

a b c
  d
e f g

There are 16 hourglasses, and an hourglass sum is the sum of an hourglass' values.

Task
Calculate the hourglass sum for every hourglass, then print the maximum hourglass sum.

**/

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

	// build the 6x6 array
	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)

	}
	countHourglass(arr)

}

func countHourglass(arr [][]int32) {
	var maxSum, sum int32
	maxSum, sum = -63, -63

	// -63 is 7 * -9, the smallest value of an hourglass

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum = arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			if sum > maxSum {
				maxSum = sum
			}
		}

	}

	fmt.Println(maxSum)

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
