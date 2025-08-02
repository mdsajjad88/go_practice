package main

import (
	"fmt"
	"strings"
)

func main(){
	data := "Sajjad, Sabbir Sumon Sabbir Sabbir, Hossan, Akash, Khan, Sham, Sabbir, Sumon, Sabbir"
	parts := strings.Split(data, ",")
	fmt.Println("Part data is ", parts)
	count := strings.Count(data, "Sabbir");
	fmt.Println("Count of same vlaue is ", count)
	
	name := "         Mohammad                  Sajjad Hossan             "
	nameIs := strings.TrimSpace(name)
	fmt.Println(" The name is ", nameIs)
	
	first_name := "Shamim"
	last_name := "Reza"
	fullname := strings.Join([]string{first_name, last_name}, " ")
	fmt.Println("full name is ", fullname)
}