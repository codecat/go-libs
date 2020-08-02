package pacman

import "bytes"

// Block is a container for binary data
type Block interface {
	Len() int
}

type block struct {
	data *bytes.Buffer
}

func (block *block) Len() int {
	return block.data.Len()
}

// NewBlock creates a new block and returns it, as well as a packer that writes to it
func NewBlock() (Block, Packer) {
	ret := &block{
		data: new(bytes.Buffer),
	}
	packer, _ := NewPacker(ret.data)
	return ret, packer
}
