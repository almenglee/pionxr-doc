package main

// func (b *DocumentBuilder) Build() error {
// 	blocks, err := b.parser.ParseBlocks()
// 	if err != nil {
// 		return debug.UnimplementedFeature(err.Error())
// 	}

// 	for _, block := range blocks {
// 		_ = block
// 		// b.BuildSections(block)
// 	}

// 	// range blocks in reverse order
// 	for i := len(blocks) - 1; i >= 0; i-- {
// 		block := blocks[i]
// 		obj := renderBlock(block)
// 		b.lines = append(b.lines[:block.Beg+1], b.lines[block.End:]...)
// 		b.lines[block.Beg] = obj
// 	}

// 	return nil
// }
