package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Print("Enter your website url: ")
	var website string
	fmt.Scan(&website)

	req, err := http.Get(website)

	if err != nil {
		panic(err)
	} else {
		fmt.Println(req.StatusCode)
	}
}
