package main
import "fmt"
func main(){
	var name[5]string
	name[3] = "Rokon"
	name[0] = "Kishore"
	
	fmt.Println("Name data is", name)
	var numbers = [8]int{1,2,3}
	var members= [5]string{"akash", "Sabbir"}
	fmt.Println(members)
	fmt.Println(numbers)
}