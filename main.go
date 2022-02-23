package main

import (
	"epitech_go_project/algorithm"
	. "fmt"
)

func main() {
	if algorithm.ExecuteAlgorithm("file.txt") {
		Printf("yes")
	} else {
		Printf("no")
	}
}
