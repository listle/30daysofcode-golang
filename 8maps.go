package main

import (
	"fmt"
	"strings"
)

func main() {
	var num int
	var input string
	var m map[string]string
	m = make(map[string]string)

	fmt.Scan(&num)

	for i := 0; i < num; i++ {
		fmt.Scan(&input)
		name := input

		fmt.Scan(&input)
		phone := strings.TrimSpace(input)
		m[name] = phone
	}

	numinputs, _ := fmt.Scan(&input)

	for numinputs > 0 {
		if m[input] == "" {
			fmt.Println("Not found")

		} else {
			fmt.Printf("%s=%s\n", input, m[input])
		}
		numinputs, _ = fmt.Scan(&input)
	}

}
