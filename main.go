package main

import (
	"yoo"
	"fmt"
	"io/ioutil"
)

func main() {
	f, err := ioutil.ReadFile("a.yoo")
	if err != nil { return }
	err = yoo.Parse(f)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
