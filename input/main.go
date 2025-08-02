package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("What's Your name ?")
	pk := bufio.NewReader(os.Stdin)
	name, _ := pk.ReadString('\n');
	
	fmt.Println("Hello Mr, ", name)
}