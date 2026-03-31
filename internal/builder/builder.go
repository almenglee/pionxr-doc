package builder

import (
	"pionxr-doc/internal/diag"
	"pionxr-doc/internal/dto"
	"pionxr-doc/internal/parser"
	"strings"
)

type Builder struct {
	parser    *parser.Parser
	collector *diag.Collector
}

func NewBuilder(src string) *Builder {
	collector := diag.NewCollector()
	return &Builder{
		parser:    parser.NewParserFromString(src, collector),
		collector: collector,
	}

}

func (b *Builder) Build() dto.Result {
	result := dto.Result{}
	blocks := b.parser.ParseBlocks()

	result.Ok = !b.collector.HasErrors()

	for _, block := range blocks {
		blk, ok := b.buildBlock(block)
		if !ok {
			result.Ok = false
			continue
		}
		result.Blocks = append(result.Blocks, blk)
	}

	for _, d := range b.collector.List() {
		result.Errors = append(result.Errors, buildError(d))
	}

	return result
}

func (b *Builder) buildBlock(block *parser.Block) (*dto.Block, bool) {
	secMap, ok := b.buildSectionMap(block.Sections)
	if !ok {
		return nil, false
	}

	return &dto.Block{
		Kind:     block.Sig.Kind,
		ID:       block.Sig.ID,
		Beg:      block.Beg,
		End:      block.End,
		Sections: secMap,
	}, true
}

func (b *Builder) buildSectionMap(sections []*parser.Section) (secMap map[string]*dto.Section, ok bool) {
	secMap = make(map[string]*dto.Section)
	ok = true
	for _, s := range sections {
		if _, exists := secMap[s.Sig.ID]; exists {
			ok = false
			b.collector.Add(diag.Errorf("duplicate section id %q", s.Sig.ID))
			continue
		}
		section := b.buildSection(s)
		secMap[s.Sig.ID] = section
	}
	return secMap, ok
}

func (b *Builder) buildSection(sec *parser.Section) *dto.Section {
	lines := b.parser.Src[sec.Beg+1 : sec.End]
	return &dto.Section{
		Kind: sec.Sig.Kind,
		Raw:  strings.Join(lines, "\n"),
	}
}
func buildError(d *diag.Diagnostic) *dto.Error {
	err := &dto.Error{
		Message: d.Message,
		Level:   string(d.Severity),
	}
	if d.Pos != nil {
		err.Pos = d.Pos.Pos()
	}
	return err
}
