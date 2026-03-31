package dto

import (
	"encoding/json"
	"io"
)

type Result struct {
	Ok     bool     `json:"ok"`
	Blocks []*Block `json:"blocks"`
	Errors []*Error `json:"errors"`
}

type Error struct {
	Message string `json:"message"`
	Level   string `json:"level"`
	Pos     string `json:"pos,omitempty"`
}

type Block struct {
	Kind     string              `json:"kind"`
	ID       string              `json:"id"`
	Beg      int                 `json:"beg"`
	End      int                 `json:"end"`
	Sections map[string]*Section `json:"sections"`
}

type Section struct {
	Kind string `json:"kind"`
	Raw  string `json:"raw"`
}

func (d *Result) EncodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	return enc.Encode(d)
}
