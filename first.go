package main

import (
	"fmt"
	"sajjad.com/myapp/latestUtil"
	"sajjad.com/myapp/newUtil"
)

func main() {
	fmt.Println("Hello Bangladesh")
	newUtil.CustomFunc("Kmn acho priya")
	latestutil.LatestFunc("Hey Sinamika")
	var name = "Sajjad"
	var age = 25
	gender := "Male"
	height := 54.848748
	height = 55.5552
	const pi = 3.1416
	fmt.Printf("pi value is ", pi)
	fmt.Printf("Heigh in %.2f\n", height); fmt.Printf("Type of name is %T\n", name);
	fmt.Printf("Type of age is %T\n", age)
	fmt.Printf("Type of gender is %T\n", gender)
	fmt.Printf("The name is %s\n", name)
	
	fmt.Printf("The name is %s, Type of age is %T, Type of height is %T ", name, age, height)
}
