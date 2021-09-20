package main

import (
	"fmt"
	"log"
	"syscall"
)

func main() {
	workingDir, err := syscall.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(workingDir)
}
