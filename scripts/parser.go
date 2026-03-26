package main

import (
	"fmt"
	"strings"
)

type MarkerKind uint8

const (
	EOF MarkerKind = iota
	BlockDecl
	SectionDecl
	BlockTerminator
	RawText
)

func (k MarkerKind) String() string {
	switch k {
	case EOF:
		return "EOF"
	case BlockDecl:
		return "BlockDecl"
	case SectionDecl:
		return "SectionDecl"
	case BlockTerminator:
		return "BlockTerminator"
	case RawText:
		return "MarkerKindNone"
	default:
		return "MarkerKind(" + fmt.Sprint(uint8(k)) + ")"
	}
}

type Marker interface {
	cursor() int
	begin() int
	kind() MarkerKind
}

type _marker struct {
	_cursor int
	_begin  int
	_kind   MarkerKind
}

func (m _marker) cursor() int      { return m._cursor }
func (m _marker) begin() int       { return m._begin }
func (m _marker) kind() MarkerKind { return m._kind }

func (m _marker) assertKind(expected MarkerKind) error {
	if m._kind != expected {
		return syntaxError("unexpected marker of kind %v, expected %v", m._kind, expected)
	}
	return nil
}

type BlockDeclMarker struct {
	_marker
	_upperBound int
}

func (m BlockDeclMarker) String() string {
	return fmt.Sprintf("BlockDeclMarker{begin=%d cursor=%d upper_bound=%d}", m._begin, m._cursor, m._upperBound)
}

type SectionDeclMarker struct {
	_marker
}

func (m SectionDeclMarker) String() string {
	return fmt.Sprintf("SectionDeclMarker{begin=%d cursor=%d}", m._begin, m._cursor)
}

type BlockTerminatorMarker struct {
	_marker
}

type EOFMarker struct {
	_marker
}

func NewBlockDeclMarker(begin, index int) *BlockDeclMarker {
	return &BlockDeclMarker{_marker: _marker{_begin: begin, _cursor: index, _kind: BlockDecl}}
}

func NewSectionDeclMarker(begin, index int) *SectionDeclMarker {
	return &SectionDeclMarker{_marker: _marker{_begin: begin, _cursor: index, _kind: SectionDecl}}
}

func NewBlockTerminatorMarker(begin, index int) *BlockTerminatorMarker {
	return &BlockTerminatorMarker{_marker: _marker{_begin: begin, _cursor: index, _kind: BlockTerminator}}
}

func NewEOFMarker(begin, index int) *EOFMarker {
	return &EOFMarker{_marker: _marker{_begin: begin, _cursor: index, _kind: EOF}}
}

func (bdecl *BlockDeclMarker) setUpperBound(upper int) {
	if bdecl == nil {
		return
	}
	bdecl._upperBound = upper
}

type BlockSignature struct {
	kind, id string
}

func (b *BlockSignature) String() string {
	if b.id == "" {
		return fmt.Sprintf("BlockSignature{kind=%q}", b.kind)
	}
	return fmt.Sprintf("BlockSignature{id=%q kind=%q}", b.id, b.kind)
}

type Block struct {
	sig      *BlockSignature
	marker   *BlockDeclMarker
	sections []section
}

func (b Block) String() string {
	return fmt.Sprintf("Block{"+b.sig.String()+", "+b.marker.String()+", %v}", b.sections)
}

type section struct {
	sig    *BlockSignature
	marker *SectionDeclMarker
}

func (s section) String() string {
	return fmt.Sprintf("Section{" + s.sig.String() + " " + s.marker.String() + "}")
}

type Parser struct {
	src     []string
	srcPath string
	curr    Marker
	markers []Marker
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
				return nil, unimplementedError("semantic error style is not implemented yet: " + err.Error())
			}
			blocks = append(blocks, block)
		}
	}

	return blocks, nil
}

func NewParser(sourcePath string, src []string) *Parser {
	return &Parser{
		src:     src,
		srcPath: sourcePath,
		curr:    nil,
	}
}

func (p *Parser) parseBlockSignature(block_decl *BlockDeclMarker) (*BlockSignature, error) {
	line, err := p.getLine(block_decl)
	if err != nil {
		return nil, err
	}

	if !strings.HasPrefix(line, "@[") {
		// it is expected that the kind of block_decl is legit at the time
		// so panic for now and implement internal error return later
		return nil, unimplementedError("parseBlockSignature is not implemented yet")
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

func (p *Parser) ParseBlockDecl(block_decl *BlockDeclMarker) (block *Block, e error) {
	// TODO: warning: ParseBlockDecl is still a stub and currently returns unimplementedError.
	block = &Block{marker: block_decl}

	sig, err := p.parseBlockSignature(block_decl)
	if err != nil {
		return nil, err
	}

	block.sig = sig

	err = p.ParseSectionList(block)
	if err != nil {
		warn(fmt.Sprintf("failed to parse section list for block %q: %v", block.sig.id, err))
	}

	return block, nil
}

func (p *Parser) ParseSectionList(parent *Block) (err error) {
	p.locate(parent.marker)
	if err := p.advance(); err != nil {
		return fatal(err.Error())
	}
	// case: update block termination point
	if !p.got(SectionDecl) {
		parent.marker.setUpperBound(p.curr.begin())
		return syntaxError("expected section declaration after block declaration")
	}

	for p.got(SectionDecl) {
		var sec *section
		sec, err = p.parseSectionDecl()
		if err != nil {
			parent.marker.setUpperBound(p.curr.begin())
			break
		}
		parent.sections = append(parent.sections, *sec)
		if err := p.advance(); err != nil {
			return fatal(err.Error())
		}
	}

	if !p.got(BlockTerminator) {
		return syntaxError("block is not properly terminated with +++")
	}

	return err
}

func (p *Parser) parseSectionDecl() (*section, error) {
	marker, ok := p.curr.(*SectionDeclMarker)
	if !ok {
		return nil, fatal("current marker is not SectionDeclMarker")
	}

	line, err := p.getLine(marker)
	if err != nil {
		return nil, fatal(err.Error())
	}

	head := strings.TrimSpace(line[3:])
	sig, err := parseMarkerHead(head)
	if err != nil {
		return nil, err
	}

	sec := &section{
		sig:    sig,
		marker: marker,
	}
	return sec, nil
}

func getMarkerKind(line string) (kind MarkerKind) {
	switch {
	case line == "+++":
		kind = BlockTerminator
	case strings.HasPrefix(line, "@["):
		kind = BlockDecl
	case strings.HasPrefix(line, "+++"):
		kind = SectionDecl
	default:
		kind = RawText
	}
	return
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

func parseMarkerHead(head string) (*BlockSignature, error) {
	parts := strings.SplitN(head, ":", 2)

	kind := strings.TrimSpace(parts[0])
	if !isIdent(kind) {
		return nil, syntaxError("invalid kind %q: kind name must be a single identifier", kind)
	}

	m := &BlockSignature{kind: kind}
	if len(parts) == 1 {
		return m, nil
	}

	id := strings.TrimSpace(parts[1])
	if !isIdent(id) {
		if id == "" {
			return nil, syntaxError("missing id after :")
		}

		return nil, syntaxError("invalid id %q: id must be a single identifier", id)
	}

	m.kind = kind
	m.id = id
	return m, nil
}
