package main

import (
	"epitech_go_project/gamemap"
)

/*func check(e error) {
	if e != nil {
		panic(e)
	}
} */

func main() {
	gamemap.GetMap("file.txt")
	//dat, err := os.ReadFile("./file.txt")
	//check(err)
	//fmt.Print(string(dat))
}
