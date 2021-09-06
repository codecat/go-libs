// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package pacman

import "encoding/binary"

func (packer *packer) WriteUint8(v interface{}) {
	switch v.(type) {
	case uint:
		binary.Write(packer.writer, packer.byteOrder, uint8(v.(uint)))
	case uint8:
		binary.Write(packer.writer, packer.byteOrder, uint8(v.(uint8)))
	case uint16:
		binary.Write(packer.writer, packer.byteOrder, uint8(v.(uint16)))
	case uint32:
		binary.Write(packer.writer, packer.byteOrder, uint8(v.(uint32)))
	case uint64:
		binary.Write(packer.writer, packer.byteOrder, uint8(v.(uint64)))

	case int:
		binary.Write(packer.writer, packer.byteOrder, uint8(v.(int)))
	case int8:
		binary.Write(packer.writer, packer.byteOrder, uint8(v.(int8)))
	case int16:
		binary.Write(packer.writer, packer.byteOrder, uint8(v.(int16)))
	case int32:
		binary.Write(packer.writer, packer.byteOrder, uint8(v.(int32)))
	case int64:
		binary.Write(packer.writer, packer.byteOrder, uint8(v.(int64)))

	case float32:
		binary.Write(packer.writer, packer.byteOrder, uint8(v.(float32)))
	case float64:
		binary.Write(packer.writer, packer.byteOrder, uint8(v.(float64)))

	default:
		panic("unhandled pack type")
	}
}

func (packer *packer) WriteUint16(v interface{}) {
	switch v.(type) {
	case uint:
		binary.Write(packer.writer, packer.byteOrder, uint16(v.(uint)))
	case uint8:
		binary.Write(packer.writer, packer.byteOrder, uint16(v.(uint8)))
	case uint16:
		binary.Write(packer.writer, packer.byteOrder, uint16(v.(uint16)))
	case uint32:
		binary.Write(packer.writer, packer.byteOrder, uint16(v.(uint32)))
	case uint64:
		binary.Write(packer.writer, packer.byteOrder, uint16(v.(uint64)))

	case int:
		binary.Write(packer.writer, packer.byteOrder, uint16(v.(int)))
	case int8:
		binary.Write(packer.writer, packer.byteOrder, uint16(v.(int8)))
	case int16:
		binary.Write(packer.writer, packer.byteOrder, uint16(v.(int16)))
	case int32:
		binary.Write(packer.writer, packer.byteOrder, uint16(v.(int32)))
	case int64:
		binary.Write(packer.writer, packer.byteOrder, uint16(v.(int64)))

	case float32:
		binary.Write(packer.writer, packer.byteOrder, uint16(v.(float32)))
	case float64:
		binary.Write(packer.writer, packer.byteOrder, uint16(v.(float64)))

	default:
		panic("unhandled pack type")
	}
}

func (packer *packer) WriteUint32(v interface{}) {
	switch v.(type) {
	case uint:
		binary.Write(packer.writer, packer.byteOrder, uint32(v.(uint)))
	case uint8:
		binary.Write(packer.writer, packer.byteOrder, uint32(v.(uint8)))
	case uint16:
		binary.Write(packer.writer, packer.byteOrder, uint32(v.(uint16)))
	case uint32:
		binary.Write(packer.writer, packer.byteOrder, uint32(v.(uint32)))
	case uint64:
		binary.Write(packer.writer, packer.byteOrder, uint32(v.(uint64)))

	case int:
		binary.Write(packer.writer, packer.byteOrder, uint32(v.(int)))
	case int8:
		binary.Write(packer.writer, packer.byteOrder, uint32(v.(int8)))
	case int16:
		binary.Write(packer.writer, packer.byteOrder, uint32(v.(int16)))
	case int32:
		binary.Write(packer.writer, packer.byteOrder, uint32(v.(int32)))
	case int64:
		binary.Write(packer.writer, packer.byteOrder, uint32(v.(int64)))

	case float32:
		binary.Write(packer.writer, packer.byteOrder, uint32(v.(float32)))
	case float64:
		binary.Write(packer.writer, packer.byteOrder, uint32(v.(float64)))

	default:
		panic("unhandled pack type")
	}
}

func (packer *packer) WriteUint64(v interface{}) {
	switch v.(type) {
	case uint:
		binary.Write(packer.writer, packer.byteOrder, uint64(v.(uint)))
	case uint8:
		binary.Write(packer.writer, packer.byteOrder, uint64(v.(uint8)))
	case uint16:
		binary.Write(packer.writer, packer.byteOrder, uint64(v.(uint16)))
	case uint32:
		binary.Write(packer.writer, packer.byteOrder, uint64(v.(uint32)))
	case uint64:
		binary.Write(packer.writer, packer.byteOrder, uint64(v.(uint64)))

	case int:
		binary.Write(packer.writer, packer.byteOrder, uint64(v.(int)))
	case int8:
		binary.Write(packer.writer, packer.byteOrder, uint64(v.(int8)))
	case int16:
		binary.Write(packer.writer, packer.byteOrder, uint64(v.(int16)))
	case int32:
		binary.Write(packer.writer, packer.byteOrder, uint64(v.(int32)))
	case int64:
		binary.Write(packer.writer, packer.byteOrder, uint64(v.(int64)))

	case float32:
		binary.Write(packer.writer, packer.byteOrder, uint64(v.(float32)))
	case float64:
		binary.Write(packer.writer, packer.byteOrder, uint64(v.(float64)))

	default:
		panic("unhandled pack type")
	}
}

func (packer *packer) WriteInt8(v interface{}) {
	switch v.(type) {
	case uint:
		binary.Write(packer.writer, packer.byteOrder, int8(v.(uint)))
	case uint8:
		binary.Write(packer.writer, packer.byteOrder, int8(v.(uint8)))
	case uint16:
		binary.Write(packer.writer, packer.byteOrder, int8(v.(uint16)))
	case uint32:
		binary.Write(packer.writer, packer.byteOrder, int8(v.(uint32)))
	case uint64:
		binary.Write(packer.writer, packer.byteOrder, int8(v.(uint64)))

	case int:
		binary.Write(packer.writer, packer.byteOrder, int8(v.(int)))
	case int8:
		binary.Write(packer.writer, packer.byteOrder, int8(v.(int8)))
	case int16:
		binary.Write(packer.writer, packer.byteOrder, int8(v.(int16)))
	case int32:
		binary.Write(packer.writer, packer.byteOrder, int8(v.(int32)))
	case int64:
		binary.Write(packer.writer, packer.byteOrder, int8(v.(int64)))

	case float32:
		binary.Write(packer.writer, packer.byteOrder, int8(v.(float32)))
	case float64:
		binary.Write(packer.writer, packer.byteOrder, int8(v.(float64)))

	default:
		panic("unhandled pack type")
	}
}

func (packer *packer) WriteInt16(v interface{}) {
	switch v.(type) {
	case uint:
		binary.Write(packer.writer, packer.byteOrder, int16(v.(uint)))
	case uint8:
		binary.Write(packer.writer, packer.byteOrder, int16(v.(uint8)))
	case uint16:
		binary.Write(packer.writer, packer.byteOrder, int16(v.(uint16)))
	case uint32:
		binary.Write(packer.writer, packer.byteOrder, int16(v.(uint32)))
	case uint64:
		binary.Write(packer.writer, packer.byteOrder, int16(v.(uint64)))

	case int:
		binary.Write(packer.writer, packer.byteOrder, int16(v.(int)))
	case int8:
		binary.Write(packer.writer, packer.byteOrder, int16(v.(int8)))
	case int16:
		binary.Write(packer.writer, packer.byteOrder, int16(v.(int16)))
	case int32:
		binary.Write(packer.writer, packer.byteOrder, int16(v.(int32)))
	case int64:
		binary.Write(packer.writer, packer.byteOrder, int16(v.(int64)))

	case float32:
		binary.Write(packer.writer, packer.byteOrder, int16(v.(float32)))
	case float64:
		binary.Write(packer.writer, packer.byteOrder, int16(v.(float64)))

	default:
		panic("unhandled pack type")
	}
}

func (packer *packer) WriteInt32(v interface{}) {
	switch v.(type) {
	case uint:
		binary.Write(packer.writer, packer.byteOrder, int32(v.(uint)))
	case uint8:
		binary.Write(packer.writer, packer.byteOrder, int32(v.(uint8)))
	case uint16:
		binary.Write(packer.writer, packer.byteOrder, int32(v.(uint16)))
	case uint32:
		binary.Write(packer.writer, packer.byteOrder, int32(v.(uint32)))
	case uint64:
		binary.Write(packer.writer, packer.byteOrder, int32(v.(uint64)))

	case int:
		binary.Write(packer.writer, packer.byteOrder, int32(v.(int)))
	case int8:
		binary.Write(packer.writer, packer.byteOrder, int32(v.(int8)))
	case int16:
		binary.Write(packer.writer, packer.byteOrder, int32(v.(int16)))
	case int32:
		binary.Write(packer.writer, packer.byteOrder, int32(v.(int32)))
	case int64:
		binary.Write(packer.writer, packer.byteOrder, int32(v.(int64)))

	case float32:
		binary.Write(packer.writer, packer.byteOrder, int32(v.(float32)))
	case float64:
		binary.Write(packer.writer, packer.byteOrder, int32(v.(float64)))

	default:
		panic("unhandled pack type")
	}
}

func (packer *packer) WriteInt64(v interface{}) {
	switch v.(type) {
	case uint:
		binary.Write(packer.writer, packer.byteOrder, int64(v.(uint)))
	case uint8:
		binary.Write(packer.writer, packer.byteOrder, int64(v.(uint8)))
	case uint16:
		binary.Write(packer.writer, packer.byteOrder, int64(v.(uint16)))
	case uint32:
		binary.Write(packer.writer, packer.byteOrder, int64(v.(uint32)))
	case uint64:
		binary.Write(packer.writer, packer.byteOrder, int64(v.(uint64)))

	case int:
		binary.Write(packer.writer, packer.byteOrder, int64(v.(int)))
	case int8:
		binary.Write(packer.writer, packer.byteOrder, int64(v.(int8)))
	case int16:
		binary.Write(packer.writer, packer.byteOrder, int64(v.(int16)))
	case int32:
		binary.Write(packer.writer, packer.byteOrder, int64(v.(int32)))
	case int64:
		binary.Write(packer.writer, packer.byteOrder, int64(v.(int64)))

	case float32:
		binary.Write(packer.writer, packer.byteOrder, int64(v.(float32)))
	case float64:
		binary.Write(packer.writer, packer.byteOrder, int64(v.(float64)))

	default:
		panic("unhandled pack type")
	}
}

func (packer *packer) WriteFloat32(v interface{}) {
	switch v.(type) {
	case uint:
		binary.Write(packer.writer, packer.byteOrder, float32(v.(uint)))
	case uint8:
		binary.Write(packer.writer, packer.byteOrder, float32(v.(uint8)))
	case uint16:
		binary.Write(packer.writer, packer.byteOrder, float32(v.(uint16)))
	case uint32:
		binary.Write(packer.writer, packer.byteOrder, float32(v.(uint32)))
	case uint64:
		binary.Write(packer.writer, packer.byteOrder, float32(v.(uint64)))

	case int:
		binary.Write(packer.writer, packer.byteOrder, float32(v.(int)))
	case int8:
		binary.Write(packer.writer, packer.byteOrder, float32(v.(int8)))
	case int16:
		binary.Write(packer.writer, packer.byteOrder, float32(v.(int16)))
	case int32:
		binary.Write(packer.writer, packer.byteOrder, float32(v.(int32)))
	case int64:
		binary.Write(packer.writer, packer.byteOrder, float32(v.(int64)))

	case float32:
		binary.Write(packer.writer, packer.byteOrder, float32(v.(float32)))
	case float64:
		binary.Write(packer.writer, packer.byteOrder, float32(v.(float64)))

	default:
		panic("unhandled pack type")
	}
}

func (packer *packer) WriteFloat64(v interface{}) {
	switch v.(type) {
	case uint:
		binary.Write(packer.writer, packer.byteOrder, float64(v.(uint)))
	case uint8:
		binary.Write(packer.writer, packer.byteOrder, float64(v.(uint8)))
	case uint16:
		binary.Write(packer.writer, packer.byteOrder, float64(v.(uint16)))
	case uint32:
		binary.Write(packer.writer, packer.byteOrder, float64(v.(uint32)))
	case uint64:
		binary.Write(packer.writer, packer.byteOrder, float64(v.(uint64)))

	case int:
		binary.Write(packer.writer, packer.byteOrder, float64(v.(int)))
	case int8:
		binary.Write(packer.writer, packer.byteOrder, float64(v.(int8)))
	case int16:
		binary.Write(packer.writer, packer.byteOrder, float64(v.(int16)))
	case int32:
		binary.Write(packer.writer, packer.byteOrder, float64(v.(int32)))
	case int64:
		binary.Write(packer.writer, packer.byteOrder, float64(v.(int64)))

	case float32:
		binary.Write(packer.writer, packer.byteOrder, float64(v.(float32)))
	case float64:
		binary.Write(packer.writer, packer.byteOrder, float64(v.(float64)))

	default:
		panic("unhandled pack type")
	}
}