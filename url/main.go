package main

import (
	"fmt"
	"net/url"
)

func main() {
	myurl := "https://www.youtube.com/watch?v=vu6ZQ-t1sUk&list=PLzjZaW71kMwSEVpdbHPr0nPo5zdzbDulm&index=26"
	updateUrl, err :=  url.Parse(myurl)
	if err != nil{
		fmt.Println("Url parse time showing error ", err)
		return
	}
	fmt.Println("URL Schema is ", updateUrl.Scheme)
	fmt.Println("URL Host is ", updateUrl.Host)
	fmt.Println("URL Path is ", updateUrl.Path)
	fmt.Println("URL Rawquery is ", updateUrl.RawQuery)
	
	updateUrl.Host = "www.ami.com"
		fmt.Println("URL Host is ", updateUrl.Host)

	
}