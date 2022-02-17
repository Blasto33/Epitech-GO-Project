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

// Palette contains the transpalette
type Palette struct {
	Pack  *Packet
	Name  string
	X, Y  uint16
	Carry bool
}

// Packet contains the packages
type Packet struct {
	Name, Color string
	X, Y        uint16
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
	var packets []Packet

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
		packets = append(packets, p)
	}

	return packets
}

// ParseFile parses the entire file and returns a slice of structs
func ParseFile(path string) {
	file := readFile(path)

	fmt.Println(file)
	//warehouse := GetMap(file)
	//packets := GetPackets(file)

	//fmt.Println(warehouse)
	//fmt.Println(packets)
}