package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type Block struct {
	ID        string   `json:"id"`
	Class     string   `json:"class,omitempty"`
	Hints     []string `json:"hints,omitempty"`
	LineIndex int      `json:"line_index"`
	LineCount int      `json:"line_count"`
	RawText   string   `json:"raw_text"`
}

func (b Block) String() string {

	return fmt.Sprintf(
		"Block{id=%q class=%v hints=%v line_index=%d line_count=%d}",
		b.ID,
		b.Class,
		b.Hints,
		b.LineIndex,
		b.LineCount,
	)
}

func (b Block) Serialize() string {

	data, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return fmt.Sprintf(`{"type":"block","error":"serialize failed: %s"}`, err)
	}

	return string(data)
}

func main() {
	lines, sourcePath, err := readDraftLines()
	if err != nil {
		log.Fatal(err)
	}

	blocks, err := findBlocks(sourcePath, lines)
	if err != nil {
		log.Fatal(err)
	}

	for _, block := range blocks {
		fmt.Println(block)
	}

	for _, block := range blocks {
		fmt.Println(block.Serialize())
	}
}

func readDraftLines() ([]string, string, error) {
	paths := []string{"draft.md", "../draft.md"}

	var file *os.File
	var err error

	for _, path := range paths {
		file, err = os.Open(path)
		if err == nil {
			defer file.Close()
			lines, scanErr := scanLines(file)
			if scanErr != nil {
				return nil, "", scanErr
			}
			return lines, path, nil
		}
	}

	return nil, "", fmt.Errorf("could not open draft.md from known paths: %v", paths)
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

func findBlocks(sourcePath string, lines []string) ([]*Block, error) {
	var blocks []*Block
	var errs []error
	seen := make(map[string]int)

	for i, line := range lines {
		block, err := parseBlockMarker(line)
		if err != nil {
			errs = append(errs, fmt.Errorf("%s:%d: %w", sourcePath, i+1, err))
			continue
		}
		if block == nil {
			continue
		}
		if firstLine, ok := seen[block.ID]; ok {
			errs = append(errs, fmt.Errorf("%s:%d: %w", sourcePath, i+1, syntaxError("duplicate id %q; previous declaration at %s:%d", block.ID, sourcePath, firstLine)))
			continue
		}

		block.LineIndex = i
		seen[block.ID] = i + 1
		blocks = append(blocks, block)
	}

	finalizeBlocks(blocks, lines)

	return blocks, joinErrors(errs)
}

type MarkerKind = uint8

const (
	BlockMarker MarkerKind = iota
	SectionHeader
	BlockTerminator
	MarkerKindNone
)

type Marker struct {
	index int
	end   int
	kind  MarkerKind
}

func scanMarkers(lines []string) []*Marker {
	var markers []*Marker
	lastBlock := &Marker{index: -1, end: -1, kind: MarkerKindNone}
	for i, line := range lines {
		line = strings.TrimSpace(line)
		kind := MarkerKindNone
		switch {
		case line == "+++":
			kind = BlockTerminator
		case strings.HasPrefix(line, "@["):
			kind = BlockMarker
		case strings.HasPrefix(line, "+++"):
			kind = SectionHeader
		}

		if kind == MarkerKindNone {
			continue
		}

		marker := &Marker{index: i, kind: kind}
		markers = append(markers, marker)

		if marker.kind == BlockMarker {
			lastBlock.end = marker.index
			lastBlock = marker
		}

		if marker.kind == BlockTerminator {
			lastBlock.end = marker.index
		}

	}
	return markers
}

func finalizeBlocks(blocks []*Block, lines []string) {
	totalLines := len(lines)

	for i, block := range blocks {
		end := totalLines
		if i+1 < len(blocks) {
			end = blocks[i+1].LineIndex
		}

		block.LineCount = end - block.LineIndex
		block.RawText = joinBlockRawText(lines, block.LineIndex+1, end)
	}
}

func joinBlockRawText(lines []string, start, end int) string {
	if start >= end {
		return ""
	}

	return strings.Join(lines[start:end], "\n")
}

func parseBlockMarker(line string) (*Block, error) {
	if !strings.HasPrefix(line, "@[") {
		return nil, nil
	}

	end := strings.Index(line, "]")
	if end < 0 {
		return nil, syntaxError("missing closing ]")
	}

	head := strings.TrimSpace(line[2:end])
	if head == "" {
		return nil, syntaxError("empty marker head")
	}

	block, err := parseMarkerHead(head)
	if err != nil {
		return nil, err
	}

	hintText := strings.TrimSpace(line[end+1:])
	if hintText != "" {
		block.Hints = strings.Fields(hintText)
		for _, h := range block.Hints {
			if !isIdent(h) {
				return nil, syntaxError("invalid hint %q: hints must be identifiers", h)
			}
		}
	}

	return block, nil
}

func parseMarkerHead(head string) (*Block, error) {
	parts := strings.SplitN(head, ":", 2)

	id := strings.TrimSpace(parts[0])
	if !isIdent(id) {
		return nil, syntaxError("invalid id %q: identifiers must start with a letter or underscore and contain only letters, digits, or underscores", id)
	}

	m := &Block{ID: id}
	if len(parts) == 1 {
		return m, nil
	}

	class := strings.TrimSpace(parts[1])
	if !isIdent(class) {
		if class == "" {
			return nil, syntaxError("missing class after :")
		}

		return nil, syntaxError("invalid class %q: class name must be a single identifier", class)
	}

	m.Class = class
	return m, nil
}

func syntaxError(format string, args ...any) error {
	return fmt.Errorf(format, args...)
}

func joinErrors(errs []error) error {
	if len(errs) == 0 {
		return nil
	}

	return errors.Join(errs...)
}

func isLetter(ch rune) bool  { return 'a' <= lower(ch) && lower(ch) <= 'z' || ch == '_' }
func lower(ch rune) rune     { return ('a' - 'A') | ch }
func isDecimal(ch rune) bool { return '0' <= ch && ch <= '9' }
func isIdent(str string) bool {
	if str == "" {
		return false
	}
	if !isLetter(rune(str[0])) {
		return false
	}

	for _, ch := range str {
		if !isLetter(rune(ch)) && !isDecimal(rune(ch)) {
			return false
		}
	}

	return true
}
