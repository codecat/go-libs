package pacman

import (
	"encoding/binary"
	"errors"
	"io"
)

// Packer allows packing of binary data
type Packer interface {
	SetByteOrder(v binary.ByteOrder)

	Write(v interface{})

	WriteUint8(v interface{})
	WriteUint16(v interface{})
	WriteUint32(v interface{})
	WriteUint64(v interface{})

	WriteInt8(v interface{})
	WriteInt16(v interface{})
	WriteInt32(v interface{})
	WriteInt64(v interface{})

	WriteFloat32(v interface{})
	WriteFloat64(v interface{})

	WriteString(v string)
	WriteBlock(v Block)

	BeginBlock() Packer
	EndBlock()
}

type packer struct {
	writer       io.Writer
	byteOrder    binary.ByteOrder
	currentBlock Block
}

func (packer *packer) SetByteOrder(v binary.ByteOrder) {
	packer.byteOrder = v
}

func (packer *packer) Write(v interface{}) {
	binary.Write(packer.writer, packer.byteOrder, v)
}

func (packer *packer) WriteString(v string) {
	packer.WriteUint32((uint32)(len(v)))
	packer.Write([]byte(v))
}

func (packer *packer) WriteBlock(v Block) {
	buffer := (v.(*block)).data.Bytes()
	packer.WriteUint32((uint32)(len(buffer)))
	packer.Write(buffer)
}

func (packer *packer) BeginBlock() Packer {
	block, ret := NewBlock()
	packer.currentBlock = block
	return ret
}

func (packer *packer) EndBlock() {
	if packer.currentBlock == nil {
		panic("there is no current block")
	}
	buffer := (packer.currentBlock.(*block)).data.Bytes()
	packer.WriteUint32((uint32)(len(buffer)))
	packer.Write(buffer)
	packer.currentBlock = nil
}

// NewPacker creates a new packer on top of a writer
func NewPacker(writer io.Writer) (Packer, error) {
	if writer == nil {
		return nil, errors.New("passed writer is nil")
	}

	return &packer{
		writer:    writer,
		byteOrder: binary.LittleEndian,
	}, nil
}
