package main
import "fmt"
func main(){
	fmt.Println("this is line 1")
	defer fmt.Println("this is line 2")
	fmt.Println("this is line 3")
	fmt.Println("this is line 4")
	defer fmt.Println("this is line 5")
	defer fmt.Println("this is line 6")
}