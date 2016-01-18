package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

const maxVal = 65535
const steps = 255.0
const factor = steps / maxVal

const width = 1000

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run waveform.go {filename}")
		return
	}
	file, err := os.Open(os.Args[1])
	//data, err := ioutil.ReadFile("audio.data")
	if err != nil {
		log.Fatal(err)
	}
	var vals []float64

	for {
		var i uint16
		err := binary.Read(file, binary.LittleEndian, &i)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		vals = append(vals, float64(i)*factor)
	}

	valsCount := len(vals)

	groups := float64(valsCount) / width
	group := int(math.Floor(groups))

	var filteredVals []float64
	tempVal := 0.0
	for i := 0; i < len(vals); i++ {
		if i%group == group-1 {
			tempVal = tempVal / float64(group)
			filteredVals = append(filteredVals, tempVal)
			tempVal = 0.0
		}
		tempVal += vals[i]
	}

	outfile, err := os.Create("audio.pbm")
	if err != nil {
		log.Fatal(err)
	}
	outfile.Write([]byte("P1\n\n"))
	outfile.Write([]byte("1000 255\n"))
	for i := 0; i < steps; i++ {
		str := ""
		for j := 0; j < width; j++ {
			newVal := i - 127
			if newVal < 0 {
				newVal = -newVal
			}
			newVal *= 2
			if float64(newVal) < filteredVals[j] {
				str += "1 "
			} else {
				str += "0 "
			}
		}
		outfile.Write([]byte(str))
	}
	outfile.Close()

}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
