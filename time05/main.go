package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println("present time->", presentTime)

	// if you want to change format 01(month)-02(day)-2006(year) 15:04:05(time) Monday(day) is special
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// create your custom time
	createdDate := time.Date(1999, time.September, 24, 13, 45, 49, 0, time.UTC)

	fmt.Println("Custom date here is->", createdDate.Format("01-02-2006 15:04:05 Monday"))
}

/*
Memory management
1. new()
	i)  Allocate memory but not INIT
	ii) You will get the memory access
	iii) Zeroed storage

2. make()
	i)  Allocate memory and INIT
	ii) You will get a memory address
	iii) Non-Zeroed storage
*/
