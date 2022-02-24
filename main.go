package main

import (
	"epitech_go_project/algorithm"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("😱 You need only 2 arguments but you enter " + strconv.Itoa(len(os.Args)))
		os.Exit(84)
	}
	cmdArgs := os.Args[1]
	if algorithm.ExecuteAlgorithm(cmdArgs) {
		fmt.Printf("😎\n")
	} else {
		fmt.Printf("🙂\n")
	}
}
