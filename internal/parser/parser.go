package parser

import (
	"fmt"
	"pionxr-doc/internal/diag"
	"strings"
)

type Parser struct {
	Src       []string
	curr      Marker
	markers   []Marker
	collector *diag.Collector
}

func NewParser(src []string, collector *diag.Collector) *Parser {
	return &Parser{
		Src:       src,
		curr:      nil,
		markers:   nil,
		collector: collector,
	}
}

func NewParserFromString(src string, collector *diag.Collector) *Parser {
	lines := strings.Split(src, "\n")
	return NewParser(lines, collector)
}

func (p *Parser) got(expected MarkerKind) bool {
	if p.curr == nil {
		return false
	}

	return p.curr.kind() == expected
}

func (p *Parser) advance() error {
	if p.curr == nil {
		return p.error("unexpected end of file")
	}

	next := p.curr.cursor() + 1
	if next >= len(p.markers) {
		return p.error("unexpected end of file")
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
	if i < 0 || i >= len(p.Src) {
		// TODO: make error fancier
		return "", p.error("line index out of bounds")
	}
	return p.Src[i], nil
}

func (p *Parser) ParseBlocks() []*Block {
	p.markers = p.scanMarkers()
	blocks := make([]*Block, 0)
	for _, mrkr := range p.markers {
		if mrkr.kind() == BlockDecl {
			block_decl := mrkr.(*BlockDeclMarker)
			block, err := p.ParseBlockDecl(block_decl)
			if err != nil {
				p.setPos(err, block_decl.pos())
			}
			blocks = append(blocks, block)
		}
	}

	return blocks
}

func (p *Parser) ParseBlockDecl(block_decl *BlockDeclMarker) (block *Block, e error) {
	// TODO: warning: ParseBlockDecl is still a stub and currently returns unimplementedError.
	block = &Block{Marker: block_decl, Beg: block_decl.begin()}

	sig, err := p.parseBlockSignature(block_decl)
	if err != nil {
		return nil, p.setPos(err, block_decl.pos())
	}

	block.Sig = sig

	err = p.ParseSectionList(block)
	if err != nil {
		p.warningAt(block.Marker.pos(), fmt.Sprintf("failed to parse section list for block %q: %v", block.Sig.ID, err))
	}
	block.End = block.Marker._upperBound
	block.resolveSectionEnds()

	return block, nil
}

func (p *Parser) parseBlockSignature(block_decl *BlockDeclMarker) (*Signature, error) {
	line, err := p.getLine(block_decl)
	pos := block_decl.pos()
	if err != nil {
		return nil, p.setPos(err, pos)
	}

	if !strings.HasPrefix(line, "@[") {
		return nil, p.internalFailureAt(pos, "blockDecl without @[ prefix")
	}

	end := strings.Index(line, "]")
	if end < 0 {
		return nil, p.syntaxErrorAt(pos, "missing closing ]")
	}

	head := strings.TrimSpace(line[2:end])
	if head == "" {
		return nil, p.syntaxErrorAt(pos, "empty marker head")
	}

	sig, err := p.parseMarkerHead(head)
	return sig, p.setPos(err, pos)
}

func (p *Parser) ParseSectionList(parent *Block) (err error) {
	p.locate(parent.Marker)
	if err := p.advance(); err != nil {
		return p.setPos(err, parent.Marker.pos())
	}
	// case: update block termination point
	if !p.got(SectionDecl) {
		parent.Marker.setUpperBound(p.curr.begin())
		return p.syntaxErrorAt(parent.Marker.pos(), "expected section declaration after block declaration")
	}

	for p.got(SectionDecl) {
		var sec *Section
		sdecl, ok := p.curr.(*SectionDeclMarker)
		if !ok {
			return p.internalFailureAt(sdecl.pos(), "unexpected marker type, expected SectionDeclMarker")
		}
		sec, err = p.parseSectionDecl(sdecl)
		if err != nil {
			parent.Marker.setUpperBound(p.curr.begin())
			break
		}
		parent.Sections = append(parent.Sections, sec)
		if err := p.advance(); err != nil {
			return p.internalFailureAt(Pos{line: sec.End}, err.Error())
		}
	}

	if !p.got(BlockTerminator) {
		return p.syntaxErrorAt(parent.Marker.pos(), "block is not properly terminated with +++")
	}

	return err
}

func (p *Parser) parseSectionDecl(marker *SectionDeclMarker) (*Section, error) {
	line, err := p.getLine(marker)
	if err != nil {
		return nil, p.setPos(err, marker.pos())
	}

	head := strings.TrimSpace(line[3:])
	sig, err := p.parseMarkerHead(head)
	if err != nil {
		return nil, p.setPos(err, marker.pos())
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
	for i, line := range p.Src {
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

	lastBlock.setUpperBound(len(p.Src))

	markers = append(markers, NewEOFMarker(len(p.Src), cnt))
	return markers
}
