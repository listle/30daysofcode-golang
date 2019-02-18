package main
import "fmt"

func main() {
	var testcases int
    fmt.Scan(&testcases)


	for i := 0; i < testcases; i++ {
        var teststring string
		fmt.Scan(&teststring)
        testbytes := []byte(teststring)
        var evenbytes []byte
		var oddbytes []byte
				      
		for j := 0; j < len(teststring); j++ {
			if j % 2 == 0 {
				evenbytes = append(evenbytes,testbytes[j])
			} else {
				oddbytes = append(oddbytes, testbytes[j])
			}
		}

		fmt.Println(string(evenbytes), string(oddbytes))
		
	}
}