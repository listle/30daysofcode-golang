package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	// Declare second integer, double, and String variables.
	var intinput uint64
	var floatinput float64
	var stringinput string

	scanner := bufio.NewScanner(os.Stdin)

	// Read and save an integer, double, and String to your variables.

	fmt.Scan(&intinput)
	fmt.Scan(&floatinput)
	_ = scanner.Scan() //advances scanner to next token, content is available through Text method
	stringinput = scanner.Text()

	// Print the sum of both integer variables on a new line.
	fmt.Println(i + intinput)

	// Print the sum of the double variables on a new line.
	sumdouble := d + floatinput
	fmt.Printf("%.1f \n", sumdouble)
	// Concatenate and print the String variables on a new line
	fmt.Println(s + stringinput)
	// The 's' variable above should be printed first.

}
