package parser

import (
	"fmt"
	"pionxr-doc/internal/debug"
	"strings"
)

type Parser struct {
	src     []string
	curr    Marker
	markers []Marker
}

func NewParser(src []string) *Parser {
	return &Parser{
		src:  src,
		curr: nil,
	}
}

func (p *Parser) got(expected MarkerKind) bool {
	if p.curr == nil {
		return false
	}

	return p.curr.kind() == expected
}

func (p *Parser) advance() error {
	if p.curr == nil {
		return syntaxError("unexpected end of file")
	}

	next := p.curr.cursor() + 1
	if next >= len(p.markers) {
		return syntaxError("unexpected end of file")
	}

	p.curr = p.markers[next]
	return nil
}

func (p *Parser) locate(m Marker) {
	p.curr = m
}

func (p *Parser) getLine(m Marker) (string, error) {
	p.curr = m
	return p._getLine(m.begin())
}

func (p *Parser) _getLine(i int) (string, error) {
	if i < 0 || i >= len(p.src) {
		// TODO: make error fancier
		return "", fmt.Errorf("tbd")
	}
	return p.src[i], nil
}

func (p *Parser) ParseBlocks() ([]*Block, error) {
	p.markers = p.scanMarkers()
	blocks := make([]*Block, 0)
	for _, mrkr := range p.markers {
		if mrkr.kind() == BlockDecl {
			block_decl := mrkr.(*BlockDeclMarker)
			block, err := p.ParseBlockDecl(block_decl)
			if err != nil {
				return nil, debug.UnimplementedFeature("semantic error style is not implemented yet: " + err.Error())
			}
			blocks = append(blocks, block)
		}
	}

	return blocks, nil
}

func (p *Parser) ParseBlockDecl(block_decl *BlockDeclMarker) (block *Block, e error) {
	// TODO: warning: ParseBlockDecl is still a stub and currently returns unimplementedError.
	block = &Block{Marker: block_decl, Beg: block_decl.begin()}

	sig, err := p.parseBlockSignature(block_decl)
	if err != nil {
		return nil, err
	}

	block.Sig = sig

	err = p.ParseSectionList(block)
	if err != nil {
		debug.Warn(fmt.Sprintf("failed to parse section list for block %q: %v", block.Sig.ID, err))
	}
	block.End = block.Marker._upperBound
	block.resolveSectionEnds()

	return block, nil
}

func (p *Parser) parseBlockSignature(block_decl *BlockDeclMarker) (*Signature, error) {
	line, err := p.getLine(block_decl)
	if err != nil {
		return nil, err
	}

	if !strings.HasPrefix(line, "@[") {
		// it is expected that the kind of block_decl is legit at the time
		// so panic for now and implement internal error return later
		return nil, debug.UnimplementedFeature("parseBlockSignature is not implemented yet")
	}

	end := strings.Index(line, "]")
	if end < 0 {
		return nil, syntaxError("missing closing ]")
	}

	head := strings.TrimSpace(line[2:end])
	if head == "" {
		return nil, syntaxError("empty marker head")
	}

	sig, err := parseMarkerHead(head)
	return sig, err
}

func (p *Parser) ParseSectionList(parent *Block) (err error) {
	p.locate(parent.Marker)
	if err := p.advance(); err != nil {
		return debug.Bug(err.Error())
	}
	// case: update block termination point
	if !p.got(SectionDecl) {
		parent.Marker.setUpperBound(p.curr.begin())
		return syntaxError("expected section declaration after block declaration")
	}

	for p.got(SectionDecl) {
		var sec *Section
		sec, err = p.parseSectionDecl()
		if err != nil {
			parent.Marker.setUpperBound(p.curr.begin())
			break
		}
		parent.Sections = append(parent.Sections, sec)
		if err := p.advance(); err != nil {
			return debug.Bug(err.Error())
		}
	}

	if !p.got(BlockTerminator) {
		return syntaxError("block is not properly terminated with +++")
	}

	return err
}

func (p *Parser) parseSectionDecl() (*Section, error) {
	marker, ok := p.curr.(*SectionDeclMarker)
	if !ok {
		return nil, debug.Bug("current marker is not SectionDeclMarker")
	}

	line, err := p.getLine(marker)
	if err != nil {
		return nil, debug.Bug(err.Error())
	}

	head := strings.TrimSpace(line[3:])
	sig, err := parseMarkerHead(head)
	if err != nil {
		return nil, err
	}

	sec := &Section{
		Sig:    sig,
		Beg:    marker.begin(),
		Marker: marker,
	}
	return sec, nil
}

func (p *Parser) scanMarkers() []Marker {
	var markers []Marker
	var lastBlock *BlockDeclMarker
	cnt := 0
	var marker Marker
	for i, line := range p.src {
		line = strings.TrimSpace(line)

		switch getMarkerKind(line) {
		case SectionDecl:
			marker = NewSectionDeclMarker(i, cnt)

		case BlockDecl:
			block := NewBlockDeclMarker(i, cnt)
			lastBlock.setUpperBound(block.begin())
			lastBlock = block
			marker = block

		case BlockTerminator:
			marker = NewBlockTerminatorMarker(i, cnt)
			lastBlock.setUpperBound(marker.begin())
			lastBlock = nil

		default:
			continue
		}

		markers = append(markers, marker)
		cnt++

	}

	lastBlock.setUpperBound(len(p.src))

	markers = append(markers, NewEOFMarker(len(p.src), cnt))
	return markers
}
