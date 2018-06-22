package main

import (
	"yoo"
	"fmt"
)

func main() {
	yoo.T()
	if err := yoo.ExecuteFile("a.yoo"); err != nil {
		fmt.Println("Error:", err)
	}
}
