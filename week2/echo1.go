package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	// inefficient
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println("Inefficient:", s)
	fmt.Println("Time taken:", time.Since(start))

	start = time.Now()
	// efficient
	s2 := strings.Join(os.Args[1:], " ")
	fmt.Println("Efficient:", s2)
	fmt.Println("Time taken:", time.Since(start))

}
