package main

import (
	"fmt"
	"log"
)

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
