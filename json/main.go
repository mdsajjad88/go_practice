package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string `json:"name"`
	Mobile int    `json:"mobile"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func main() {
	user := Person{Name: "Rokon", Mobile: 1548758965, Age: 24, Gender: "Male"}
	fmt.Println("user data is : ", user)

	jsonUser, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Something is wrong ", err)
	}
	fmt.Println("Json user data is :", string(jsonUser))

	var personData Person

	errs := json.Unmarshal(jsonUser, &personData)
	if errs != nil {
		fmt.Println("error in unmarshal is = ", errs)
		return
	}
	fmt.Println("Hello person data is = ", personData)
}
