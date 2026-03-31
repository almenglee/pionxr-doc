package main

import (
	"bufio"
	"os"
	"pionxr-doc/internal/debug"
	"pionxr-doc/internal/parser"
)

type DocumentBuilder struct {
	doc    *os.File
	parser *parser.Parser
	lines  []string
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

func NewDocumentBuilderFromFile(file *os.File) (*DocumentBuilder, error) {
	lines, err := scanLines(file)
	if err != nil {
		return nil, debug.UndecidedError("failed to read lines from file: " + err.Error())
	}

	return &DocumentBuilder{
		doc:    file,
		lines:  lines,
		parser: parser.NewParser(lines),
	}, nil
}

func NewDocumentBuilder(path string) (*DocumentBuilder, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, debug.UndecidedError("failed to open file: " + err.Error())
	}

	return NewDocumentBuilderFromFile(file)
}

func (b *DocumentBuilder) Build() error {
	blocks, err := b.parser.ParseBlocks()
	if err != nil {
		return debug.UnimplementedFeature(err.Error())
	}

	for _, block := range blocks {
		b.parser.BuildSections(block)
	}

	// range blocks in reverse order
	for i := len(blocks) - 1; i >= 0; i-- {
		block := blocks[i]
		obj := renderBlock(block)
		b.lines = append(b.lines[:block.Beg+1], b.lines[block.End:]...)
		b.lines[block.Beg] = obj
	}

	return nil
}
