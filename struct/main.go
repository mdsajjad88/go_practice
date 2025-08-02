package main
import "fmt"

type Student struct{
	firstName string
	lastName string
	age int
	roll int
}
// func main(){
// 	var studentData Student
// 	studentData.firstName = "Ähmed"
// 	studentData.lastName = "Shahzad"
// 	studentData.age = 24
// 	studentData.roll = 10
// 	fmt.Println("Student Data is ", studentData) 
// 	student2 := Student{
// 		firstName: "Najmul",
// 		lastName: "Hossain",
// 		age: 21,
// 		roll: 11,
// 	}
// 	fmt.Println("Second student data is ", student2)
	
// 	student3 := new(Student)
// 	student3.firstName = "Shakib"
// 	student3.lastName = "Al Hasan"
// 	student3.age = 35
// 	student3.roll = 1
// 	fmt.Println("Shakib data is :", student3)
// }
type Contact struct{
	mobile int
	email string
}
type StudentInfo struct{
	personal_info Student
	contact_info Contact
}
func main(){
	var ourStudent  StudentInfo
	ourStudent.personal_info.age = 45
	ourStudent.personal_info.firstName = "Kashem"
	fmt.Println("Data is ", ourStudent)
}