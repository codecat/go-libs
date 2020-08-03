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

	WriteUInt8(v uint8)
	WriteUInt16(v uint16)
	WriteUInt32(v uint32)
	WriteUInt64(v uint64)

	// Writes as a 32 bit unsigned integer
	WriteUInt(v uint)

	WriteInt8(v int8)
	WriteInt16(v int16)
	WriteInt32(v int32)
	WriteInt64(v int64)

	// Writes as a 32 bit integer
	WriteInt(v int)

	WriteFloat32(v float32)
	WriteFloat64(v float64)

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

func (packer *packer) WriteUInt8(v uint8) {
	binary.Write(packer.writer, packer.byteOrder, v)
}
func (packer *packer) WriteUInt16(v uint16) {
	binary.Write(packer.writer, packer.byteOrder, v)
}
func (packer *packer) WriteUInt32(v uint32) {
	binary.Write(packer.writer, packer.byteOrder, v)
}
func (packer *packer) WriteUInt64(v uint64) {
	binary.Write(packer.writer, packer.byteOrder, v)
}

func (packer *packer) WriteUInt(v uint) {
	packer.WriteUInt32((uint32)(v))
}

func (packer *packer) WriteInt8(v int8) {
	binary.Write(packer.writer, packer.byteOrder, v)
}
func (packer *packer) WriteInt16(v int16) {
	binary.Write(packer.writer, packer.byteOrder, v)
}
func (packer *packer) WriteInt32(v int32) {
	binary.Write(packer.writer, packer.byteOrder, v)
}
func (packer *packer) WriteInt64(v int64) {
	binary.Write(packer.writer, packer.byteOrder, v)
}

func (packer *packer) WriteInt(v int) {
	packer.WriteInt32((int32)(v))
}

func (packer *packer) WriteFloat32(v float32) {
	binary.Write(packer.writer, packer.byteOrder, v)
}
func (packer *packer) WriteFloat64(v float64) {
	binary.Write(packer.writer, packer.byteOrder, v)
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
