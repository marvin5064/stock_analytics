package main

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

func main() {
	resp, body, errs := gorequest.New().Get("http://www.google.com").End()
	if errs != nil {
		fmt.Println(errs)
	}
	fmt.Println(resp, body)
	fmt.Println("init")
}
