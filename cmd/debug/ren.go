package main

import "pionxr-doc/internal/parser"

func _renderCollapsible(block *parser.Block) string {
	var body *parser.Section
	for _, sec := range block.Sections {
		switch sec.Sig.ID {
		case "body":
			body = sec
		}
	}
	if body == nil {
		return "collapsible block missing body section"
	}

	return "<details>\n" + body.Raw + "\n</details>"
}

func renderBlock(block *parser.Block) string {
	switch block.Sig.Kind {
	case "collapsible":
		return _renderCollapsible(block)
		// case "choice":
		// 	return renderChoice(block)
		// case "short_answer":
		// 	return renderShortAnswer(block)
	}

	return "unimplemented render for block kind: " + block.Sig.Kind
}
