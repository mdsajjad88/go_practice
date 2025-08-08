package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// file, err := os.Create("test5.txt")
	// if err != nil {
	// 	fmt.Println("Some thing went wrong", err)
	// }
	// defer file.Close()
	// content := "Hello World by Back Bencher!"
	// byte, errors :=	io.WriteString(file, content+"\n"+"How are you?")
	// if errors != nil{
	// 	fmt.Println("Something went wrong to file write time ", errors)
	// }
	// fmt.Println("Total content byte = ", byte)
	// fmt.Println("Successfully create file")
	
	open, err1 := os.Open("test5.txt")
	if err1 != nil{
		fmt.Println("Something went wrong to open file ", err1)
	}
	defer open.Close()
	buffer := make([]byte, 1024)
	for{
		number, err2 := open.Read(buffer)
		if err2 == io.EOF{
			break
		}
		if err2 != nil{
			fmt.Println("Error while reading the file ", err2)
		}
		
		fmt.Println(string(buffer[:number]))
	}
}