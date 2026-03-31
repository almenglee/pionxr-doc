package parser

import (
	"fmt"
	"strings"
)

type Signature struct {
	Kind string
	ID   string
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

	Marker *SectionDeclMarker
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

func (p *Parser) parseMarkerHead(head string) (*Signature, error) {
	parts := strings.SplitN(head, ":", 2)

	kind := strings.TrimSpace(parts[0])
	if !isIdent(kind) {
		// TODO: rewind block end
		return nil, p.syntaxError(fmt.Sprintf("invalid kind %q: kind name must be a single identifier", kind))
	}

	m := &Signature{Kind: kind}
	if len(parts) == 1 {
		return m, nil
	}

	id := strings.TrimSpace(parts[1])
	if !isIdent(id) {
		if id == "" {
			// TODO: rewind block end
			return nil, p.syntaxError("missing id after :")
		}

		return nil, p.syntaxError(fmt.Sprintf("invalid id %q: id must be a single identifier", id))
	}

	m.Kind = kind
	m.ID = id
	return m, nil
}
