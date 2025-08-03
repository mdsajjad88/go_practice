package main

import (
	"fmt"
	"time"
)
func main(){
	var time = time.Now()
	fmt.Println("time is ", time)
	
	formatTime := time.Format("2006-01-02, Monday, 15:04:05")
	fmt.Println("Formatted time is ", formatTime)
}