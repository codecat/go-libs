//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "VPackerType=uint8,uint16,uint32,uint64,int8,int16,int32,int64,float32,float64"

package pacman

import (
	"encoding/binary"

	"github.com/cheekybits/genny/generic"
)

// VPackerType is the integer type
type VPackerType generic.Type

func (packer *packer) WriteVPackerType(v interface{}) {
	switch v.(type) {
	case uint:
		binary.Write(packer.writer, packer.byteOrder, VPackerType(v.(uint)))
	case uint8:
		binary.Write(packer.writer, packer.byteOrder, VPackerType(v.(uint8)))
	case uint16:
		binary.Write(packer.writer, packer.byteOrder, VPackerType(v.(uint16)))
	case uint32:
		binary.Write(packer.writer, packer.byteOrder, VPackerType(v.(uint32)))
	case uint64:
		binary.Write(packer.writer, packer.byteOrder, VPackerType(v.(uint64)))

	case int:
		binary.Write(packer.writer, packer.byteOrder, VPackerType(v.(int)))
	case int8:
		binary.Write(packer.writer, packer.byteOrder, VPackerType(v.(int8)))
	case int16:
		binary.Write(packer.writer, packer.byteOrder, VPackerType(v.(int16)))
	case int32:
		binary.Write(packer.writer, packer.byteOrder, VPackerType(v.(int32)))
	case int64:
		binary.Write(packer.writer, packer.byteOrder, VPackerType(v.(int64)))

	case float32:
		binary.Write(packer.writer, packer.byteOrder, VPackerType(v.(float32)))
	case float64:
		binary.Write(packer.writer, packer.byteOrder, VPackerType(v.(float64)))
	}
}
