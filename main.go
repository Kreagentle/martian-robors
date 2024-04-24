package main

import (
	"fmt"
	"martian-robots/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Println("Cant execute the program")
		return
	}
}
