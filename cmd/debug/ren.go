package main

import (
	"pionxr-doc/internal/dto"
)

func _renderCollapsible(block *dto.Block) string {
	var body dto.Section
	for id, sec := range block.Sections {
		switch id {
		case "body":
			body = *sec
		}
	}
	if body == (dto.Section{}) {
		return "collapsible block missing body section"
	}

	return "<details>\n" + body.Raw + "\n</details>"
}

func renderBlock(block *dto.Block) string {
	switch block.Kind {
	case "collapsible":
		return _renderCollapsible(block)
		// case "choice":
		// 	return renderChoice(block)
		// case "short_answer":
		// 	return renderShortAnswer(block)
	}

	return "unimplemented render for block kind: " + block.Kind
}
