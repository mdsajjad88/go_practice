package main
import "fmt"

func main(){
	num := 25
	ak := &num
	new := 4545
	fmt.Println("Hello pointer ", ak)
	fmt.Println("Pointer value is = ", *ak)
	
	var newPointer *int
	newPointer = &new
	if(newPointer == nil){
		fmt.Println("Pointer is nill")
	}else{
		fmt.Println("Hello new pointer ", newPointer)
	fmt.Println("New Pointer value is = ", *newPointer)
	}
	
	org := 50
	modifyValue(&org)
	fmt.Println("value in pointer is ", org)
	
	// a := 10
	// b := 20
	// c := a + b
	// updateValue(c)
	// fmt.Println("Update value in c variable is ", c)
}
// this call modify value by reference
func modifyValue(value *int){
	*value = *value * 10
}

// func updateValue(value int){
// 	value = value + 9
// }