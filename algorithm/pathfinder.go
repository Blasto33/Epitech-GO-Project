package algorithm

import (
	"epitech_go_project/warehouse"
	"fmt"
)

type algorithm struct {
	Ware       *warehouse.Map
	Listpal    []warehouse.Palette
	ListPacket []warehouse.Packet
	ListTruck  []warehouse.Truck
	TwoDMap    [][]uint16
}

// colorWeight get the color to print with the weight
func colorWeight(weight string) uint16 {
	if weight == "BLUE" {
		return 500
	} else if weight == "GREEN" {
		return 200
	}
	return 100
}

// isEmpty know if there is no more warehouse.Packet left
func (a *algorithm) isEmpty() bool {
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

func (a *algorithm) createMap() {
	for i := 0; i < int(a.Ware.Y); i++ {
		for j := 0; j < int(a.Ware.X); j++ {
			a.TwoDMap[i][j] = 0
		}
	}
	for _, pal := range a.Listpal {
		a.TwoDMap[pal.Y][pal.X] = 1
	}
	for _, pack := range a.ListPacket {
		a.TwoDMap[pack.Y][pack.X] = 2
	}
	for _, trk := range a.ListTruck {
		a.TwoDMap[trk.Y][trk.X] = 3
	}
}

func (a *algorithm) findPacket(y uint16, x uint16) *warehouse.Packet {
	for _, pack := range a.ListPacket {
		if pack.X == x && pack.Y == y {
			return &pack
		}
	}
	return nil
}

func (a *algorithm) makePalMove(palIndex int, destX uint16, destY uint16) bool {
	tmpX := a.Listpal[palIndex].X
	tmpY := a.Listpal[palIndex].Y
	if (a.Listpal[palIndex].X < destX) && (a.TwoDMap[a.Listpal[palIndex].Y][a.Listpal[palIndex].X+1] == 0) {
		a.Listpal[palIndex].X++
		return false
	} else if (a.Listpal[palIndex].Y < destY) && (a.TwoDMap[a.Listpal[palIndex].Y+1][a.Listpal[palIndex].X] == 0) {
		a.Listpal[palIndex].Y++
		return false
	}
	if (a.Listpal[palIndex].X > destX) && (a.TwoDMap[a.Listpal[palIndex].Y][a.Listpal[palIndex].X-1] == 0) {
		a.Listpal[palIndex].X--
	} else if (a.Listpal[palIndex].Y > destY) && (a.TwoDMap[a.Listpal[palIndex].Y-1][a.Listpal[palIndex].X] == 0) {
		a.Listpal[palIndex].Y--
	}
	return tmpX == a.Listpal[palIndex].X && tmpY == a.Listpal[palIndex].Y
}

func indexFromPtr(lis []warehouse.Packet, ptr *warehouse.Packet) int {
	for i := 0; i < len(lis); i++ {
		if lis[i].X == ptr.X && lis[i].Y == ptr.Y {
			return i
		}
	}
	return 0
}

func (a *algorithm) commandPalette(palIndex int, packIndex int) {
	if a.makePalMove(palIndex, a.ListPacket[packIndex].X, a.ListPacket[packIndex].Y) {
		a.Listpal[palIndex].Command = fmt.Sprintf("%s WAIT\n", a.Listpal[palIndex].Name)
	} else {
		a.TwoDMap[a.Listpal[palIndex].X][a.Listpal[palIndex].Y] = 1
		a.Listpal[palIndex].Command = fmt.Sprintf("%s GO [%d,%d]\n", a.Listpal[palIndex].Name, a.Listpal[palIndex].X, a.Listpal[palIndex].Y)
	}
}

func (a *algorithm) removePacketIndex(i int) {
	if len(a.ListPacket) == 1 {
		a.ListPacket = make([]warehouse.Packet, 0)
	} else {
		a.ListPacket = append(a.ListPacket[:i], a.ListPacket[i+1:]...)
	}
}

func (a *algorithm) gotoPacket(palIndex int, packIndex int) {
	var ptr *warehouse.Packet
	if int(a.Listpal[palIndex].X+1) < len(a.TwoDMap[0]) && a.TwoDMap[a.Listpal[palIndex].Y][a.Listpal[palIndex].X+1] == 2 {
		ptr = a.findPacket(a.Listpal[palIndex].Y, a.Listpal[palIndex].X+1)
	} else if int(a.Listpal[palIndex].Y+1) < len(a.TwoDMap) && a.TwoDMap[a.Listpal[palIndex].Y+1][a.Listpal[palIndex].X] == 2 {
		ptr = a.findPacket(a.Listpal[palIndex].Y+1, a.Listpal[palIndex].X)
	}
	if a.Listpal[palIndex].X > 0 && a.TwoDMap[a.Listpal[palIndex].Y][a.Listpal[palIndex].X-1] == 2 {
		ptr = a.findPacket(a.Listpal[palIndex].Y, a.Listpal[palIndex].X-1)
	} else if a.Listpal[palIndex].Y > 0 && a.TwoDMap[a.Listpal[palIndex].Y-1][a.Listpal[palIndex].X] == 2 {
		ptr = a.findPacket(a.Listpal[palIndex].Y-1, a.Listpal[palIndex].X)
	}
	if ptr == nil {
		a.commandPalette(palIndex, packIndex)
	} else {
		a.TwoDMap[ptr.X][ptr.Y] = 0
		a.removePacketIndex(indexFromPtr(a.ListPacket, ptr))
		a.Listpal[palIndex].Carry = true
		a.Listpal[palIndex].Pack = ptr
		fmt.Printf("%s TAKE %s %s\n", a.Listpal[palIndex].Name, ptr.Name, ptr.Color)
	}
}

// gotoTruck make the Palette move to the Truck
func (a *algorithm) gotoTruck(palIndex int, truckIndex int) {
	if Abs(int(a.Listpal[palIndex].X)-int(a.ListTruck[truckIndex].X))+(Abs(int(a.Listpal[palIndex].Y)-int(a.ListTruck[truckIndex].Y))) == 1 {
		if uint32(colorWeight(a.Listpal[palIndex].Pack.Color)) < a.ListTruck[truckIndex].MaxContent-a.ListTruck[truckIndex].Content {
			a.ListTruck[truckIndex].Content += uint32(colorWeight(a.Listpal[palIndex].Pack.Color))
			a.Listpal[palIndex].Command = fmt.Sprintf("%s LEAVE %s %s\n", a.Listpal[palIndex].Name, a.Listpal[palIndex].Pack.Name, a.Listpal[palIndex].Pack.Color)
			a.Listpal[palIndex].Pack = nil
			a.Listpal[palIndex].Carry = false
		} else {
			a.Listpal[palIndex].Command = fmt.Sprintf("%s WAIT\n", a.Listpal[palIndex].Name)
		}
	} else {
		if a.makePalMove(palIndex, a.ListTruck[truckIndex].X, a.ListTruck[truckIndex].Y) {
			a.Listpal[palIndex].Command = fmt.Sprintf("%s WAIT\n", a.Listpal[palIndex].Name)
		} else {
			a.TwoDMap[a.Listpal[palIndex].X][a.Listpal[palIndex].Y] = 1
			a.Listpal[palIndex].Command = fmt.Sprintf("%s GO [%d,%d]\n", a.Listpal[palIndex].Name, a.Listpal[palIndex].X, a.Listpal[palIndex].Y)
		}
	}
}

func (a *algorithm) printTruck() {
	for _, Truck := range a.ListTruck {
		if Truck.Round > 0 {
			Truck.Round--
			fmt.Printf("%s GONE %d/%d", Truck.Name, Truck.Content, Truck.MaxContent)
			if Truck.Round == 0 {
				Truck.Content = 0
			}
			continue
		}
		if (Truck.MaxContent-Truck.Content < 500 && Truck.MaxContent > 500) || a.isEmpty() {
			Truck.Round = Truck.MaxRound
			fmt.Printf("%s GONE %d/%d\n", Truck.Name, Truck.Content, Truck.MaxContent)
		} else {
			fmt.Printf("%s WAITING %d/%d\n", Truck.Name, Truck.Content, Truck.MaxContent)
		}
	}
}

func (a *algorithm) printPal() {
	move := make([]int, 0)
	for packindex := 0; packindex < len(a.ListPacket); packindex++ {
		distance := a.Ware.X * a.Ware.Y
		tmp := getPal(int(distance), a.ListPacket[packindex], a.Listpal, move)
		if tmp != -1 {
			move = append(move, tmp)
			a.gotoPacket(tmp, packindex)
		}
	}
	for palindex := 0; palindex < len(a.Listpal); palindex++ {
		if Find(move, palindex) || !a.Listpal[palindex].Carry {
			continue
		}
		distance := a.Ware.X * a.Ware.Y
		tmp := getTruck(int(distance), a.Listpal[palindex], a.ListTruck)
		if tmp != -1 {
			move = append(move, tmp)
			a.gotoTruck(palindex, tmp)
		}
	}
	for palindex := 0; palindex < len(a.Listpal); palindex++ {
		if !Find(move, palindex) {
			a.Listpal[palindex].Command = fmt.Sprintf("%s WAIT\n", a.Listpal[palindex].Name)
		}
	}
	for _, pal := range a.Listpal {
		fmt.Printf(pal.Command)
	}
}

// Abs get the absolute value of a int
func Abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

// getTruck get the closest Truck of a Palette
func getTruck(distance int, pal warehouse.Palette, truckList []warehouse.Truck) int {
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

// getPal get the closest Palette of a Packet
func getPal(distance int, pack warehouse.Packet, listpal []warehouse.Palette, move []int) int {
	ret := -1
	for i := 0; i < len(listpal); i++ {
		if Find(move, i) || listpal[i].Carry {
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

func initMap(x uint16, y uint16) [][]uint16 {
	ret := make([][]uint16, y)
	for i := 0; i < int(y); i++ {
		ret[i] = make([]uint16, x)
	}
	return ret
}

// ExecuteAlgorithm Execute all the algorithm and write the output of the program
func ExecuteAlgorithm(path string) bool {
	content := warehouse.ParseFile(path)
	tmp := warehouse.GetMap(content)
	algo := algorithm{Ware: tmp, Listpal: warehouse.GetPalettes(content), ListPacket: warehouse.GetPackets(content), ListTruck: warehouse.GetTrucks(content), TwoDMap: initMap(tmp.X, tmp.Y)}
	for i := 0; i < int(algo.Ware.NbIter); i++ {
		algo.createMap()
		fmt.Printf("tour %d\n", i+1)
		algo.printPal()
		algo.printTruck()
		fmt.Printf("\n")
	}

	return algo.isEmpty()
}
