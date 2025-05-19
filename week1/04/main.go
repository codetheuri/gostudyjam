package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welocme to my application")
	fmt.Println("please rate my application between 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating my application,", input)

	numRate, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
		// panic(err)
	} else {
		fmt.Println("added 1 to your rating", numRate+1)
	}
}
