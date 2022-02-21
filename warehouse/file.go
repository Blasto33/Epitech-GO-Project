package warehouse

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) []string {
	s := []string{}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("err")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			file.Close()
			os.Exit(84)
		}
		array := []string{}
		array = append(array, scanner.Text())
		s = append(s, array...)
	}

	if scanner.Err() != nil {
		log.Fatal(err)
	}

	return s
}

// Map contains the warehouse map
type Map struct {
	NbIter uint32
	X, Y   uint16
}

// Packet contains the packets
type Packet struct {
	Name, Color string
	X, Y        uint16
}

// Palette contains the palettes
type Palette struct {
	Pack    *Packet
	Name    string
	Command string
	X, Y    uint16
	Carry   bool
}

// Truck contains the trucks elements
type Truck struct {
	Name                                 string
	MaxContent, Content, Round, MaxRound uint32
	X, Y                                 uint16
}

// GetMap parses the map from the file
func GetMap(file []string) (wrh *Map) {
	w := Map{}
	mapVar := strings.Fields(file[0])

	for i := 0; i < 3; i++ {
		value, err := strconv.Atoi(mapVar[i])
		if err != nil {
			log.Fatal(err)
		}

		switch i {
		case 0:
			w.X = uint16(value)
		case 1:
			w.Y = uint16(value)
		case 2:
			w.NbIter = uint32(value)
		}
	}

	return &w
}

// GetPackets parses the packages from the file
func GetPackets(file []string) (pck []Packet) {
	p := Packet{}
	var packet []Packet

	for i := 1; i < len(file); i++ {
		packetsVar := strings.Fields(file[i])
		if len(packetsVar) != 4 {
			break
		}

		p.Name = packetsVar[0]
		p.Color = packetsVar[3]
		for w := 1; w < 3; w++ {
			value, err := strconv.Atoi(packetsVar[w])
			if err != nil {
				log.Fatal(err)
			}

			switch w {
			case 2:
				p.X = uint16(value)
			case 3:
				p.Y = uint16(value)

			}
		}
		packet = append(packet, p)
	}

	return packet
}

func GetPalettes(file []string) (plt []Palette) {
	p := Palette{}
	var palettes []Palette

	for i := 1; i < len(file); i++ {
		palettesVar := strings.Fields(file[i])
		if len(palettesVar) == 3 {

			p.Name = palettesVar[0]
			for w := 1; w < 3; w++ {
				value, err := strconv.Atoi(palettesVar[w])
				if err != nil {
					log.Fatal(err)
				}

				switch w {
				case 2:
					p.X = uint16(value)
				case 3:
					p.Y = uint16(value)

				}
			}
			palettes = append(palettes, p)
		}
	}
	return palettes
}

func GetTrucks(file []string) (trk []Truck) {
	p := Truck{}
	var trucks []Truck

	for i := 1; i < len(file); i++ {
		trucksVar := strings.Fields(file[i])
		if len(trucksVar) == 5 {

			p.Name = trucksVar[0]
			for w := 1; w < 5; w++ {
				value, err := strconv.Atoi(trucksVar[w])
				if err != nil {
					log.Fatal(err)
				}

				switch w {
				case 1:
					p.X = uint16(value)
				case 2:
					p.Y = uint16(value)
				case 3:
					p.MaxContent = uint32(value)
				case 4:
					p.MaxRound = uint32(value)
				}
			}
			trucks = append(trucks, p)
		}
	}
	return trucks
}

// ParseFile parses the entire file and returns a slice of structs
func ParseFile(path string) {
	file := readFile(path)

	//fmt.Println(file)
	//warehouse := GetMap(file)
	//packets := GetPackets(file)
	//palette := GetPalettes(file)
	trucks := GetTrucks(file)

	fmt.Println(trucks)
	//fmt.Println(palette)
	//fmt.Println(warehouse)
	//fmt.Println(packets)
}
