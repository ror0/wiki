package blockchain

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) AddBlock(data []byte) {
	prev_block := bc.Blocks[len(bc.Blocks) - 1]
	new_block := NewBlock(data, prev_block.CalcHeaderHash())
	bc.Blocks = append(bc.Blocks, new_block)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

