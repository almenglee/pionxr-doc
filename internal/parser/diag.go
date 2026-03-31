package parser

import (
	"fmt"
	"pionxr-doc/internal/diag"
)

type Pos struct {
	line int
}

func (p Pos) Pos() string {
	return fmt.Sprintf("%d", p.line)
}

func (p Pos) String() string {
	return p.Pos()
}

// prevent orphaned diagnostics by attaching position to them
func (p *Parser) setPos(err error, pos diag.Position) error {
	d, ok := err.(*diag.Diagnostic)
	if !ok {
		return err
	}

	if d.Pos == nil {
		d.Pos = pos
	}
	return d
}

func (p *Parser) error(msg string) error {
	return p.errorAt(nil, msg)
}

func (p *Parser) errorAt(pos diag.Position, msg string) error {
	return p.diag(pos, msg, diag.Error)
}

func (p *Parser) internalFailureAt(pos diag.Position, msg string) error {
	return p.errorAt(pos, "internal parser invariant violated: "+msg)
}

func (p *Parser) internalFailure(msg string) error {
	return p.internalFailureAt(nil, msg)
}

func (p *Parser) syntaxErrorAt(pos diag.Position, msg string) error {
	return p.errorAt(pos, "syntax error: "+msg)
}

func (p *Parser) syntaxError(msg string) error {
	return p.syntaxErrorAt(nil, msg)
}

func (p *Parser) warningAt(pos diag.Position, msg string) error {
	return p.diag(pos, msg, diag.Warning)
}

func (p *Parser) diag(pos diag.Position, msg string, severity diag.Severity) error {
	d := &diag.Diagnostic{
		Severity: severity,
		Message:  msg,
		Pos:      pos,
	}
	p.collector.Add(d)
	return d
}
