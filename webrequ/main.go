package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://awc.careneterp.com:82/api/getCustomers")

	if err != nil {
		fmt.Println("Error getting in response ", err)
	}
	defer res.Body.Close()
	fmt.Printf("Type of response is %T\n ", res)
	data, err2 := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("getting error when read all data ", err2)
	}
	fmt.Println("Api Data is = ", string(data))
	
	
}