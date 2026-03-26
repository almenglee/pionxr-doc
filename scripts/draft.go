package main

import (
	"fmt"
	"log"
)

func main() {
	lines, sourcePath, err := readDraftLines()
	if err != nil {
		log.Fatal(err)
	}

	parser := NewParser(sourcePath, lines)

	blocks, err := parser.ParseBlocks()
	if err != nil {
		log.Fatal(err)
	}

	for _, block := range blocks {
		fmt.Println(block)
	}

	// for _, block := range blocks {
	// 	fmt.Println(block.Serialize())
	// }
}
