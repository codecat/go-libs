package pacman

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

// Unpacker allows unpacking of binary data
type Unpacker interface {
	SetByteOrder(v binary.ByteOrder)

	ReadUInt8() uint8
	ReadUInt16() uint16
	ReadUInt32() uint32
	ReadUInt64() uint64
	ReadUIntFrom8() uint
	ReadUIntFrom16() uint
	ReadUIntFrom32() uint
	ReadUIntFrom64() uint

	ReadInt8() int8
	ReadInt16() int16
	ReadInt32() int32
	ReadInt64() int64
	ReadIntFrom8() int
	ReadIntFrom16() int
	ReadIntFrom32() int
	ReadIntFrom64() int

	ReadFloat32() float32
	ReadFloat64() float64

	ReadString() string
	ReadBlock() Unpacker
}

type unpacker struct {
	reader    io.Reader
	byteOrder binary.ByteOrder
}

func (unpacker *unpacker) SetByteOrder(v binary.ByteOrder) {
	unpacker.byteOrder = v
}

func (unpacker *unpacker) ReadUInt8() uint8 {
	var ret uint8
	binary.Read(unpacker.reader, unpacker.byteOrder, &ret)
	return ret
}
func (unpacker *unpacker) ReadUInt16() uint16 {
	var ret uint16
	binary.Read(unpacker.reader, unpacker.byteOrder, &ret)
	return ret
}
func (unpacker *unpacker) ReadUInt32() uint32 {
	var ret uint32
	binary.Read(unpacker.reader, unpacker.byteOrder, &ret)
	return ret
}
func (unpacker *unpacker) ReadUInt64() uint64 {
	var ret uint64
	binary.Read(unpacker.reader, unpacker.byteOrder, &ret)
	return ret
}
func (unpacker *unpacker) ReadUIntFrom8() uint {
	var ret uint8
	binary.Read(unpacker.reader, unpacker.byteOrder, &ret)
	return uint(ret)
}
func (unpacker *unpacker) ReadUIntFrom16() uint {
	var ret uint16
	binary.Read(unpacker.reader, unpacker.byteOrder, &ret)
	return uint(ret)
}
func (unpacker *unpacker) ReadUIntFrom32() uint {
	var ret uint32
	binary.Read(unpacker.reader, unpacker.byteOrder, &ret)
	return uint(ret)
}
func (unpacker *unpacker) ReadUIntFrom64() uint {
	var ret uint64
	binary.Read(unpacker.reader, unpacker.byteOrder, &ret)
	return uint(ret)
}

func (unpacker *unpacker) ReadInt8() int8 {
	var ret int8
	binary.Read(unpacker.reader, unpacker.byteOrder, &ret)
	return ret
}
func (unpacker *unpacker) ReadInt16() int16 {
	var ret int16
	binary.Read(unpacker.reader, unpacker.byteOrder, &ret)
	return ret
}
func (unpacker *unpacker) ReadInt32() int32 {
	var ret int32
	binary.Read(unpacker.reader, unpacker.byteOrder, &ret)
	return ret
}
func (unpacker *unpacker) ReadInt64() int64 {
	var ret int64
	binary.Read(unpacker.reader, unpacker.byteOrder, &ret)
	return ret
}
func (unpacker *unpacker) ReadIntFrom8() int {
	var ret int8
	binary.Read(unpacker.reader, unpacker.byteOrder, &ret)
	return int(ret)
}
func (unpacker *unpacker) ReadIntFrom16() int {
	var ret int16
	binary.Read(unpacker.reader, unpacker.byteOrder, &ret)
	return int(ret)
}
func (unpacker *unpacker) ReadIntFrom32() int {
	var ret int32
	binary.Read(unpacker.reader, unpacker.byteOrder, &ret)
	return int(ret)
}
func (unpacker *unpacker) ReadIntFrom64() int {
	var ret int64
	binary.Read(unpacker.reader, unpacker.byteOrder, &ret)
	return int(ret)
}

func (unpacker *unpacker) ReadFloat32() float32 {
	var ret float32
	binary.Read(unpacker.reader, unpacker.byteOrder, &ret)
	return ret
}
func (unpacker *unpacker) ReadFloat64() float64 {
	var ret float64
	binary.Read(unpacker.reader, unpacker.byteOrder, &ret)
	return ret
}

func (unpacker *unpacker) ReadString() string {
	strLength := unpacker.ReadUInt32()
	buffer := make([]byte, strLength)
	unpacker.reader.Read(buffer)
	return string(buffer)
}

func (unpacker *unpacker) ReadBlock() Unpacker {
	bufferLength := unpacker.ReadUInt32()
	buffer := make([]byte, bufferLength)
	unpacker.reader.Read(buffer)
	ret, _ := NewUnpacker(bytes.NewBuffer(buffer))
	return ret
}

// NewUnpacker creates a new unpacker from a reader
func NewUnpacker(reader io.Reader) (Unpacker, error) {
	if reader == nil {
		return nil, errors.New("passed reader is nil")
	}

	return &unpacker{
		reader:    reader,
		byteOrder: binary.LittleEndian,
	}, nil
}
