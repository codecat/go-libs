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

	ReadUint8() uint8
	ReadUint16() uint16
	ReadUint32() uint32
	ReadUint64() uint64

	ReadInt8() int8
	ReadInt16() int16
	ReadInt32() int32
	ReadInt64() int64

	ReadUintFrom8() uint
	ReadUintFrom16() uint
	ReadUintFrom32() uint
	ReadUintFrom64() uint

	ReadIntFrom8() int
	ReadIntFrom16() int
	ReadIntFrom32() int
	ReadIntFrom64() int

	ReadFloat32() float32
	ReadFloat64() float64

	ReadString() string
	ReadStringSize(size int) string

	ReadBlock() Unpacker
	ReadBlockSize(size int) Unpacker

	ReadBytes(size int) []byte
}

type unpacker struct {
	reader    io.Reader
	byteOrder binary.ByteOrder
}

func (unpacker *unpacker) SetByteOrder(v binary.ByteOrder) {
	unpacker.byteOrder = v
}

func (unpacker *unpacker) ReadUintFrom8() uint {
	var ret uint8
	binary.Read(unpacker.reader, unpacker.byteOrder, &ret)
	return uint(ret)
}
func (unpacker *unpacker) ReadUintFrom16() uint {
	var ret uint16
	binary.Read(unpacker.reader, unpacker.byteOrder, &ret)
	return uint(ret)
}
func (unpacker *unpacker) ReadUintFrom32() uint {
	var ret uint32
	binary.Read(unpacker.reader, unpacker.byteOrder, &ret)
	return uint(ret)
}
func (unpacker *unpacker) ReadUintFrom64() uint {
	var ret uint64
	binary.Read(unpacker.reader, unpacker.byteOrder, &ret)
	return uint(ret)
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

func (unpacker *unpacker) ReadString() string {
	strLength := unpacker.ReadIntFrom32()
	return unpacker.ReadStringSize(strLength)
}

func (unpacker *unpacker) ReadStringSize(size int) string {
	buffer := make([]byte, size)
	unpacker.reader.Read(buffer)
	return string(buffer)
}

func (unpacker *unpacker) ReadBlock() Unpacker {
	bufferLength := unpacker.ReadIntFrom32()
	return unpacker.ReadBlockSize(bufferLength)
}

func (unpacker *unpacker) ReadBlockSize(size int) Unpacker {
	buffer := make([]byte, size)
	unpacker.reader.Read(buffer)
	ret, _ := NewUnpacker(bytes.NewBuffer(buffer))
	return ret
}

func (unpacker *unpacker) ReadBytes(size int) []byte {
	buffer := make([]byte, size)
	unpacker.reader.Read(buffer)
	return buffer
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
