package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type __block_deprecated struct {
	ID        string   `json:"id"`
	Class     string   `json:"class,omitempty"`
	Hints     []string `json:"hints,omitempty"`
	LineIndex int      `json:"line_index"`
	LineCount int      `json:"line_count"`
	RawText   string   `json:"raw_text"`
}

func (b __block_deprecated) String() string {

	return fmt.Sprintf(
		"Block{id=%q class=%v hints=%v line_index=%d line_count=%d}",
		b.ID,
		b.Class,
		b.Hints,
		b.LineIndex,
		b.LineCount,
	)
}

func (b __block_deprecated) Serialize() string {

	data, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return fmt.Sprintf(`{"type":"block","error":"serialize failed: %s"}`, err)
	}

	return string(data)
}

func parseBlockMarker(line string) (*__block_deprecated, error) {
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

	block, err := parseMarkerHead_deprecated(head)
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

func parseMarkerHead_deprecated(head string) (*__block_deprecated, error) {
	parts := strings.SplitN(head, ":", 2)

	id := strings.TrimSpace(parts[0])
	if !isIdent(id) {
		return nil, syntaxError("invalid id %q: identifiers must start with a letter or underscore and contain only letters, digits, or underscores", id)
	}

	m := &__block_deprecated{ID: id}
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
func findBlocks(sourcePath string, lines []string) ([]*__block_deprecated, error) {
	var blocks []*__block_deprecated
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

func finalizeBlocks(blocks []*__block_deprecated, lines []string) {
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
