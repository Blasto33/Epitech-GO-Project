package algorithm

import (
	"fmt"
)

type warehouseMap struct {
	NbIter uint32
	X, Y   uint16
}

type palette struct {
	Pack    *packet
	Name    string
	Command string
	X, Y    uint16
	Carry   bool
}

type packet struct {
	Name         string
	X, Y, Weight uint16
}

type truck struct {
	Name                                 string
	MaxContent, Content, Round, MaxRound uint32
	X, Y                                 uint16
}

type algorithm struct {
	Ware       warehouseMap
	Listpal    []palette
	Listpacket []packet
	Listtruck  []truck
	TwoDMap    [][]uint8
}

func remove(lis []packet, pack *packet) []packet {
	dest := make([]packet, 0)
	for i := 0; i < len(lis); i++ {
		if &lis[i] != pack {
			dest = append(dest, lis[i])
		}
	}
	return dest
}

// colorWeight get the color to print with the weight
func colorWeight(weight uint16) string {
	if weight == 500 {
		return "BLUE"
	} else if weight == 200 {
		return "GREEN"
	}
	return "YELLOW"
}

func (a algorithm) createMap() {
	for i := 0; i < int(a.Ware.X); i++ {
		for j := 0; j < int(a.Ware.Y); j++ {
			a.TwoDMap[i][j] = 0
		}
	}
	for _, pal := range a.Listpal {
		a.TwoDMap[pal.Y][pal.X] = 1
	}
	for _, pack := range a.Listpacket {
		a.TwoDMap[pack.X][pack.Y] = 2
	}
	for _, trk := range a.Listtruck {
		a.TwoDMap[trk.X][trk.Y] = 3
	}
}

func (a algorithm) findPacket(x uint16, y uint16) *packet {
	for _, pack := range a.Listpacket {
		if pack.X == x && pack.Y == y {
			return &pack
		}
	}
	return nil
}

func (a algorithm) gotopacket(palIndex int, packIndex int) {
	var ptr *packet
	if a.TwoDMap[a.Listpal[palIndex].Y][a.Listpal[palIndex].X+1] == 2 {
		ptr = a.findPacket(a.Listpal[palIndex].Y, a.Listpal[palIndex].X+1)
	} else if a.TwoDMap[a.Listpal[palIndex].Y+1][a.Listpal[palIndex].X] == 2 {
		ptr = a.findPacket(a.Listpal[palIndex].Y+1, a.Listpal[palIndex].X)
	} else if a.TwoDMap[a.Listpal[palIndex].Y][a.Listpal[palIndex].X-1] == 2 {
		ptr = a.findPacket(a.Listpal[palIndex].Y, a.Listpal[palIndex].X-1)
	} else if a.TwoDMap[a.Listpal[palIndex].Y-1][a.Listpal[palIndex].X] == 2 {
		ptr = a.findPacket(a.Listpal[palIndex].Y-1, a.Listpal[palIndex].X)
	}
	if ptr == nil {
		if (a.Listpal[palIndex].Y < a.Listtruck[packIndex].Y) && (a.TwoDMap[a.Listpal[palIndex].Y-1][a.Listpal[palIndex].X] == 0) {
			a.Listpal[palIndex].Y--
		} else if (a.Listpal[palIndex].X > a.Listtruck[packIndex].X) && (a.TwoDMap[a.Listpal[palIndex].Y][a.Listpal[palIndex].X-1] == 0) {
			a.Listpal[palIndex].X--
		} else if (a.Listpal[palIndex].Y < a.Listtruck[packIndex].Y) && (a.TwoDMap[a.Listpal[palIndex].Y+1][a.Listpal[palIndex].X] == 0) {
			a.Listpal[palIndex].Y++
		} else if (a.Listpal[palIndex].X < a.Listtruck[packIndex].X) && (a.TwoDMap[a.Listpal[palIndex].Y][a.Listpal[palIndex].X+1] == 0) {
			a.Listpal[palIndex].X++
		} else {
			a.Listpal[palIndex].Command = fmt.Sprintf("%s WAIT\n", a.Listpal[palIndex].Name)
			return
		}
		a.Listpal[palIndex].Command = fmt.Sprintf("%s GO [%d,%d]\n", a.Listpal[palIndex].Name, a.Listpal[palIndex].X, a.Listpal[palIndex].Y)
	} else {
		a.TwoDMap[ptr.X][ptr.Y] = 0
		remove(a.Listpacket, ptr)
		a.Listpal[palIndex].Carry = true
		a.Listpal[palIndex].Pack = ptr
		fmt.Printf("%s TAKE %s %s\n", a.Listpal[palIndex].Name, ptr.Name, colorWeight(ptr.Weight))
	}
}

// gototruck make the palette move to the truck
func (a algorithm) gototruck(palIndex int, truckIndex int) {
	if Abs(int(a.Listpal[palIndex].X)-int(a.Listtruck[truckIndex].X))+(Abs(int(a.Listpal[palIndex].Y)-int(a.Listtruck[truckIndex].Y))) == 1 {
		if uint32(a.Listpal[palIndex].Pack.Weight) > a.Listtruck[truckIndex].MaxContent-a.Listtruck[truckIndex].Content {
			a.Listpal[palIndex].Command = fmt.Sprintf("%s LEAVE %s %s\n", a.Listpal[palIndex].Name, a.Listpal[palIndex].Pack.Name, colorWeight(a.Listpal[palIndex].Pack.Weight))
			a.Listpal[palIndex].Pack = nil
			a.Listpal[palIndex].Carry = false
		} else {
			a.Listpal[palIndex].Command = fmt.Sprintf("%s WAIT\n", a.Listpal[palIndex].Name)
		}
	} else {
		if (a.Listpal[palIndex].X < a.Listtruck[truckIndex].X) && (a.TwoDMap[a.Listpal[palIndex].Y][a.Listpal[palIndex].X+1] == 0) {
			a.Listpal[palIndex].X++
		} else if (a.Listpal[palIndex].Y < a.Listtruck[truckIndex].Y) && (a.TwoDMap[a.Listpal[palIndex].Y+1][a.Listpal[palIndex].X] == 0) {
			a.Listpal[palIndex].Y++
		} else if (a.Listpal[palIndex].X > a.Listtruck[truckIndex].X) && (a.TwoDMap[a.Listpal[palIndex].Y][a.Listpal[palIndex].X-1] == 0) {
			a.Listpal[palIndex].X--
		} else if (a.Listpal[palIndex].Y < a.Listtruck[truckIndex].Y) && (a.TwoDMap[a.Listpal[palIndex].Y-1][a.Listpal[palIndex].X] == 0) {
			a.Listpal[palIndex].Y--
		} else {
			a.Listpal[palIndex].Command = fmt.Sprintf("%s WAIT\n", a.Listpal[palIndex].Name)
			return
		}
		a.TwoDMap[a.Listpal[palIndex].X][a.Listpal[palIndex].Y] = 1
		a.Listpal[palIndex].Command = fmt.Sprintf("%s GO [%d,%d]\n", a.Listpal[palIndex].Name, a.Listpal[palIndex].X, a.Listpal[palIndex].Y)
	}
}

// isEmpty know if there is no more packet left
func (a algorithm) isEmpty() bool {
	for _, i := range a.TwoDMap {
		for _, j := range i {
			if j == 2 {
				return false
			}
		}
	}
	for _, pal := range a.Listpal {
		if pal.Carry {
			return false
		}
	}
	return true
}

// Abs get the absolute value of a int
func Abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

// getTruck get the closest truck of a palette
func getTruck(distance int, pal palette, truckList []truck) int {
	ret := -1
	for i := 0; i < len(truckList); i++ {
		tmp := (Abs(int(pal.X) - int(truckList[i].X))) + (Abs(int(pal.Y) - int(truckList[i].Y)))
		if distance > tmp {
			ret = i
			distance = tmp
		}
	}
	return ret
}

// getPal get the closest palette of a packet
func getPal(distance int, pack packet, listpal []palette) int {
	ret := -1
	for i := 0; i < len(listpal); i++ {
		if listpal[i].Carry {
			continue
		}
		tmp := (Abs(int(pack.X) - int(listpal[i].X))) + (Abs(int(pack.Y) - int(listpal[i].Y)))
		if distance > tmp {
			ret = i
			distance = tmp
		}
	}
	return ret
}

// Find find if a value exist in a slice of int
func Find(slice []int, index int) bool {
	for _, i := range slice {
		if i == index {
			return true
		}
	}
	return false
}

// executeAlgorithm Execute all the algorithm and write the output of the program
func executeAlgorithm(warehouse warehouseMap) bool {
	algo := algorithm{Ware: warehouse, TwoDMap: make([][]uint8, warehouse.Y, warehouse.X)}
	for i := 0; i < int(algo.Ware.NbIter); i++ {
		algo.createMap()
		move := make([]int, 0)
		fmt.Printf("tour %d\n", i+1)
		fmt.Printf("\n")
		for packindex := 0; packindex < len(algo.Listpacket); packindex++ {
			distance := algo.Ware.X * algo.Ware.Y
			tmp := getPal(int(distance), algo.Listpacket[packindex], algo.Listpal)
			if tmp != -1 {
				move = append(move, tmp)
				algo.gotopacket(tmp, packindex)
			}
		}
		for palindex := 0; palindex < len(algo.Listpal); palindex++ {
			if Find(move, palindex) || !algo.Listpal[palindex].Carry {
				continue
			}
			distance := algo.Ware.X * algo.Ware.Y
			tmp := getTruck(int(distance), algo.Listpal[palindex], algo.Listtruck)
			if tmp != -1 {
				move = append(move, tmp)
				algo.gototruck(palindex, tmp)
			}
		}
		for palindex := 0; palindex < len(algo.Listpal); palindex++ {
			if !Find(move, palindex) {
				algo.Listpal[palindex].Command = fmt.Sprintf("%s WAIT\n", algo.Listpal[palindex].Name)
			}
		}
		for _, pal := range algo.Listpal {
			fmt.Printf(pal.Command)
		}
		for _, truck := range algo.Listtruck {
			if truck.Round > 0 {
				truck.Round--
				fmt.Printf("%s GONE %d/%d", truck.Name, truck.Content, truck.MaxContent)
				if truck.Round == 0 {
					truck.Content = 0
				}
			} else if (truck.MaxContent-truck.Content < 500 && truck.MaxContent > 500) || algo.isEmpty() {
				truck.Round = truck.MaxRound
				fmt.Printf("%s GONE %d/%d", truck.Name, truck.Content, truck.MaxContent)
			} else {
				fmt.Printf("%s WAITING %d/%d", truck.Name, truck.Content, truck.MaxContent)
			}
		}
	}
	return algo.isEmpty()
}
