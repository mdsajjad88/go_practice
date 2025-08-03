package main

import (
	"fmt"
	"time"
)
func main(){
	var time = time.Now()
	fmt.Println("time is ", time)
	
	formatTime := time.Format("2006-01-02, Monday, 15:04:05")
	formatSecondTime := time.Format("02/01/2006, Monday, 3:04:05 PM")
	fmt.Println("Formatted time is ", formatTime)
	fmt.Println("Formatted second time is ", formatSecondTime)
}