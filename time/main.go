package main

import (
	"fmt"
	"time"
)
func main(){
	var times = time.Now()
	fmt.Println("time is ", times)
	
	formatTime := times.Format("2006-01-02, Monday, 15:04:05")
	formatSecondTime := times.Format("02/01/2006, Monday, 3:04:05 PM")
	fmt.Println("Formatted time is ", formatTime)
	fmt.Println("Formatted second time is ", formatSecondTime)
	
	strTime := "31-01-2023"
	frmStrTime, _ := time.Parse("02-01-2006", strTime)
	fmt.Println("Formatted String time is ", frmStrTime)	
}