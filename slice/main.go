package main
import "fmt"
func main(){
	number := make([]int, 5, 6)
	number = append(number, 4)
	number = append(number, 8)
	number = append(number, 6)
	number = append(number, 7)
	fmt.Println("Slice : ", number)
	fmt.Println("Slice : ", len(number))
	fmt.Println("Slice : ", cap(number))
		
}