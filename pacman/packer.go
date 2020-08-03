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

	WriteUInt8(v interface{})
	WriteUInt16(v interface{})
	WriteUInt32(v interface{})
	WriteUInt64(v interface{})

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

func (packer *packer) WriteUInt8(v interface{}) {
	binary.Write(packer.writer, packer.byteOrder, uint8(v))
}
func (packer *packer) WriteUInt16(v interface{}) {
	binary.Write(packer.writer, packer.byteOrder, uint16(v))
}
func (packer *packer) WriteUInt32(v interface{}) {
	binary.Write(packer.writer, packer.byteOrder, uint32(v))
}
func (packer *packer) WriteUInt64(v interface{}) {
	binary.Write(packer.writer, packer.byteOrder, uint64(v))
}

func (packer *packer) WriteInt8(v interface{}) {
	binary.Write(packer.writer, packer.byteOrder, int8(v))
}
func (packer *packer) WriteInt16(v interface{}) {
	binary.Write(packer.writer, packer.byteOrder, int16(v))
}
func (packer *packer) WriteInt32(v interface{}) {
	binary.Write(packer.writer, packer.byteOrder, int32(v))
}
func (packer *packer) WriteInt64(v interface{}) {
	binary.Write(packer.writer, packer.byteOrder, int64(v))
}

func (packer *packer) WriteFloat32(v interface{}) {
	binary.Write(packer.writer, packer.byteOrder, float32(v))
}
func (packer *packer) WriteFloat64(v interface{}) {
	binary.Write(packer.writer, packer.byteOrder, float64(v))
}

func (packer *packer) WriteString(v string) {
	packer.WriteUInt32((uint32)(len(v)))
	packer.Write([]byte(v))
}

func (packer *packer) WriteBlock(v Block) {
	buffer := (v.(*block)).data.Bytes()
	packer.WriteUInt32((uint32)(len(buffer)))
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
	packer.WriteUInt32((uint32)(len(buffer)))
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
