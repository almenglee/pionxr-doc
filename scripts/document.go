package main

import "os"

type DocumentBuilder struct {
	doc    *os.File
	parser *Parser
	lines  []string
}

func NewDocumentBuilderFromFile(file *os.File) (*DocumentBuilder, error) {
	lines, err := scanLines(file)
	if err != nil {
		return nil, undecidedError("failed to read lines from file: " + err.Error())
	}

	return &DocumentBuilder{
		doc:    file,
		lines:  lines,
		parser: NewParser(lines),
	}, nil
}

func NewDocumentBuilder(path string) (*DocumentBuilder, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, undecidedError("failed to open file: " + err.Error())
	}

	return NewDocumentBuilderFromFile(file)
}

func (b *DocumentBuilder) Build() error {
	blocks, err := b.parser.ParseBlocks()
	if err != nil {
		return unimplementedFeature(err.Error())
	}

	for _, block := range blocks {
		b.parser.buildSections(block)
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
