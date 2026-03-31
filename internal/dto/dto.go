package dto

import "pionxr-doc/internal/parser"

type Block struct {
	Kind     string             `json:"kind"`
	ID       string             `json:"id"`
	Beg      int                `json:"beg"`
	End      int                `json:"end"`
	Sections map[string]Section `json:"sections"`
}

type Section struct {
	Kind string `json:"kind"`
	Raw  string `json:"raw"`
}

func FromParserBlock(pb *parser.Block) Block {
	sections := make(map[string]Section)
	for _, sec := range pb.Sections {
		sections[sec.Sig.ID] = Section{
			Kind: sec.Sig.Kind,
			Raw:  sec.Raw,
		}
	}

	return Block{
		Kind:     pb.Sig.Kind,
		ID:       pb.Sig.ID,
		Beg:      pb.Beg,
		End:      pb.End,
		Sections: sections,
	}
}
