package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
)

func main() {
	inputFile := flag.String("input", "", "Input file path")
	outputFile := flag.String("output", "", "Output file path")
	rate := flag.Float64("rate", 10.0, "Chance percentage of hex swapping")
	startPercent := flag.Int("startp", 0, "Start percentage of the range")
	endPercent := flag.Int("endp", 100, "End percentage of the range")
	chunk := flag.Int("chunk", 10, "Number of hex values to swap each time")
	mode := flag.String("mode", "swap", "Mode of operation: swap, copyswap, blackout, purerandom")

	flag.Parse()

	if *inputFile == "" || *outputFile == "" {
		fmt.Println("Please provide input and output file paths.")
		return
	}

	if *startPercent < 0 || *endPercent > 100 || *startPercent >= *endPercent {
		fmt.Println("Invalid range percentages.")
		return
	}

	if *rate < 0 || *rate > 100 {
		fmt.Println("Invalid rate percentage.")
		return
	}

	inputBytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	fileSize := len(inputBytes)
	startIdx := (*startPercent * fileSize) / 100
	endIdx := (*endPercent * fileSize) / 100

	if endIdx > fileSize {
		endIdx = fileSize
	}

	swaps := 0
	for i := startIdx; i < endIdx; i += *chunk {
		if rand.Float64()*100 < *rate {
			swaps++
			switch *mode {
			case "swap":
				for j := 0; j < *chunk; j++ {
					if i+j < endIdx {
						inputBytes[i+j], inputBytes[i+j-*chunk] = inputBytes[i+j-*chunk], inputBytes[i+j]
					}
				}
			case "copyswap":
				copy(inputBytes[i:i+*chunk], inputBytes[i-*chunk:i])
			case "blackout":
				for j := 0; j < *chunk; j++ {
					if i+j < endIdx {
						inputBytes[i+j] = 0x00
					}
				}
			case "purerandom":
				for j := 0; j < *chunk; j++ {
					if i+j < endIdx {
						inputBytes[i+j] = byte(rand.Intn(256))
					}
				}
			default:
				fmt.Println("Invalid mode specified.")
				return
			}
		}
	}

	err = ioutil.WriteFile(*outputFile, inputBytes, 0644)
	if err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
		return
	}

	fmt.Printf("File copy with random hex manipulation completed successfully. Operations performed: %d\n", swaps)
}

