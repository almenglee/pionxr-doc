package parser

import (
	"fmt"
	"strings"
)

type Signature struct {
	Kind string
	ID   string
}

func (b *Signature) String() string {
	if b.ID == "" {
		return fmt.Sprintf("BlockSignature{kind=%q}", b.Kind)
	}
	return fmt.Sprintf("BlockSignature{id=%q kind=%q}", b.ID, b.Kind)
}

type Block struct {
	Sig      *Signature
	Beg      int
	End      int
	Sections []*Section

	Marker *BlockDeclMarker
}

type Section struct {
	Sig *Signature
	Beg int
	End int
	Raw string

	Marker *SectionDeclMarker
}

func (b Block) String() string {
	return fmt.Sprintf("Block{"+b.Sig.String()+", "+b.Marker.String()+", %v, end=%d }", b.Sections, b.End)
}

func (s Section) String() string {
	if s.Raw != "" {
		return fmt.Sprintf("Section{"+s.Sig.String()+", "+s.Marker.String()+", end=%d, raw=%q}", s.End, s.Raw)
	}
	return fmt.Sprintf("Section{"+s.Sig.String()+", "+s.Marker.String()+", end=%d}", s.End)
}

func (b *Block) resolveSectionEnds() {
	for i := range b.Sections {
		if i+1 < len(b.Sections) {
			b.Sections[i].End = b.Sections[i+1].Marker.begin()
		} else {
			b.Sections[i].End = b.End
		}
	}
}

func parseMarkerHead(head string) (*Signature, error) {
	parts := strings.SplitN(head, ":", 2)

	kind := strings.TrimSpace(parts[0])
	if !isIdent(kind) {
		// TODO: rewind block end
		println("warning: invalid kind in marker head, treating it as no kind")
		return nil, syntaxError("invalid kind %q: kind name must be a single identifier", kind)
	}

	m := &Signature{Kind: kind}
	if len(parts) == 1 {
		return m, nil
	}

	id := strings.TrimSpace(parts[1])
	if !isIdent(id) {
		if id == "" {
			// TODO: rewind block end
			println("warning: missing id after : in marker head, treating it as no id")
			return nil, syntaxError("missing id after :")
		}

		return nil, syntaxError("invalid id %q: id must be a single identifier", id)
	}

	m.Kind = kind
	m.ID = id
	return m, nil
}
