package main

import (
	"fmt"
	"log"
)

// func main() {
// 	path := "draft.md"
// 	builder, := NewDocumentBuilder(path)
// 	_ = path
// 	lines, _, err := readDraftLines()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	parser := NewParser(lines)

// 	blocks, err := parser.ParseBlocks()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, block := range blocks {
// 		parser.buildSections(block)
// 		fmt.Println(block)
// 	}
// 	for _, block := range blocks {
// 		fmt.Println(block.Serialize())
// 	}

// }

func main() {
	path := "draft.md"
	builder, err := NewDocumentBuilder(path)
	if err != nil {
		log.Fatal(err)
	}

	if err := builder.Build(); err != nil {
		log.Fatal(err)
	}

	for _, line := range builder.lines {
		fmt.Println(line)
	}

}
