// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package pacman

import "encoding/binary"

func (unpacker *unpacker) ReadUint8() uint8 {
	var ret uint8
	binary.Read(unpacker.reader, unpacker.byteOrder, &ret)
	return ret
}

func (unpacker *unpacker) ReadUint16() uint16 {
	var ret uint16
	binary.Read(unpacker.reader, unpacker.byteOrder, &ret)
	return ret
}

func (unpacker *unpacker) ReadUint32() uint32 {
	var ret uint32
	binary.Read(unpacker.reader, unpacker.byteOrder, &ret)
	return ret
}

func (unpacker *unpacker) ReadUint64() uint64 {
	var ret uint64
	binary.Read(unpacker.reader, unpacker.byteOrder, &ret)
	return ret
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
