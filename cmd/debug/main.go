package main

import (
	"bufio"
	"fmt"
	"os"
	"pionxr-doc/internal/builder"
)

func main() {
	path := "draft.md"
	builder, err := NewBuilderFromPath(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result := builder.Build()
	err = result.EncodeJSON(os.Stdout)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func scanLines(file *os.File) ([]string, error) {
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func NewBuilderFromPath(path string) (*builder.Builder, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return builder.NewBuilder(string(content)), nil
}
