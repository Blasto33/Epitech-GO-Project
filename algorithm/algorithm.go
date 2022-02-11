package algorithm

type warehouseMap struct {
	X, Y   uint16
	NbIter uint32
}

type algorithm struct {
	Ware warehouseMap
}

func executeAlgorithm(warehouse warehouseMap) error {
	algo := algorithm{Ware: warehouse}
	for i := 0; i < int(algo.Ware.NbIter); i++ {
	}
	return nil
}
