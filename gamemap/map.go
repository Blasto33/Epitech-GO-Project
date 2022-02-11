package gamemap

import "fmt"

type warehouseMap struct {
	X, Y   uint16
	NbIter uint32
}

func readFile(file string) {

	return
}

func newWarehouseMap(file string) *warehouseMap {
	w := warehouseMap{X: 4, Y: 4, NbIter: 1000}
	return &w
}

// GetMap retrieves every data from the file passed as parameter
func GetMap(file string) {
	fmt.Println("I got map")
	fmt.Println(newWarehouseMap(file))
}
