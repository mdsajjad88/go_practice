package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("test5.txt")
	if err != nil {
		fmt.Println("Some thing went wrong", err)
	}
	defer file.Close()
	content := "Hello World by Back Bencher!"
	byte, errors :=	io.WriteString(file, content+"\n"+"How are you?")
	if errors != nil{
		fmt.Println("Something went wrong to file write time ", errors)
	}
	fmt.Println("Total content byte = ", byte)
	fmt.Println("Successfully create file")
}