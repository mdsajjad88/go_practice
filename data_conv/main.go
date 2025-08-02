package main

import (
	"fmt"
	"strconv"
)
func main(){
	var number = 42
	var conv float64 = float64(number)
	fmt.Printf("type is = %t\n", conv)
	
	numberToStr := strconv.Itoa(number)
	fmt.Printf("now number type is  = %T\n", numberToStr)
	
	strNum := "1542"
	strToNum, _ := strconv.Atoi(strNum)
	fmt.Printf("type of StrToNum is = %t\n", strToNum)
	
	strFloat := "3.1416"
	strToFloat, _ := strconv.ParseFloat(strFloat, 64)
	fmt.Printf("str to float conv type check %T\n", strToFloat) 
	strToFloat32, _ := strconv.ParseFloat(strFloat, 32)
	parse32 := float32(strToFloat32)
	fmt.Printf("str to float conv type check %T\n", parse32) 
}