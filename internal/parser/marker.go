package parser

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

type SectionDeclMarker struct {
	_marker
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

func (m BlockDeclMarker) String() string {
	return fmt.Sprintf("BlockDeclMarker{begin=%d cursor=%d upper_bound=%d}", m._begin, m._cursor, m._upperBound)
}

func (m SectionDeclMarker) String() string {
	return fmt.Sprintf("SectionDeclMarker{begin=%d cursor=%d}", m._begin, m._cursor)
}

func (bdecl *BlockDeclMarker) setUpperBound(upper int) {
	if bdecl == nil {
		return
	}
	bdecl._upperBound = upper
}
