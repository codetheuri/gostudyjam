package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time  study of golang")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday "))

	createdDate := time.Date(2024, time.June, 27,12,59,23,0,time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("2006-Jan-02 Mon"))
}
