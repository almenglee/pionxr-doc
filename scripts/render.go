package main

import "github.com/pelletier/go-toml/v2"

// type Entity struct {
// }

// type EntitySection struct {
// }

// func parseEntityFromBlock(block *Block) (*Entity, error) {
// 	switch block.Sig.Kind {
// 	case "collapsible":
// 		return parseCollapsible(block)
// 	case "choice":
// 		return parseChoice(block)
// 	case "short_answer":
// 		return parseShortAnswer(block)
// 	}

// 	return nil, unimplementedError("parseEntityFromBlock: fail message not implemented")
// }

// func (ntt *Entity) render() string {
// 	return unimplementedError("Entity.render: fail message not implemented").Error()
// }

// collapsible
// note
// choice
// short_answer

func renderBlock(block *Block) string {
	switch block.Sig.Kind {
	case "collapsible":
		return renderCollapsible(block)
	case "choice":
		return renderChoice(block)
	case "short_answer":
		return renderShortAnswer(block)
	}

	return unimplementedError("renderBlock: fail message not implemented").Error()
}

func parseSectionFromBlock(block *Block) (*Section, error) {
	return nil, unimplementedError("parseSectionFromBlock: fail message not implemented")
}

func parseDataFromSection(sec *Section) (interface{}, error) {

	return nil, unimplementedError("parseDataFromSection: fail message not implemented")
}

func parseTOML(data string) (any, error) {
	var m map[string]any
	err := toml.Unmarshal([]byte(data), &m) // validate TOML format

	if err != nil {
		return nil, err
	}

	return m, nil
}

func renderCollapsible(block *Block) string {
	var data, body *Section
	// check section schema
	for _, sec := range block.Sections {
		switch sec.Sig.ID {
		case "data":
			data = sec
		case "body":
			body = sec
		}
	}
	switch {
	case data == nil:
		return unimplementedError("renderCollapsible: missing data section").Error()
	case body == nil:
		return unimplementedError("renderCollapsible: missing body section").Error()
	case data.Sig.Kind != "toml":
		return unimplementedError("renderCollapsible: data section must have kind=toml").Error()
	case body.Sig.Kind != "md":
		return unimplementedError("renderCollapsible: body section must have kind=md").Error()
	}

	dat, err := parseTOML(data.Raw)
	if err != nil {
		return unimplementedError("renderCollapsible: failed to parse TOML data: " + err.Error()).Error()
	}

	_ = dat // TODO: use data

	rtn := "<details>\n"
	rtn += body.Raw + "\n"
	rtn += "</details>"

	return rtn
}

func renderChoice(block *Block) string {
	return unimplementedError("renderChoice: fail message not implemented").Error()
}

func renderShortAnswer(block *Block) string {
	return unimplementedError("renderShortAnswer: fail message not implemented").Error()
}
