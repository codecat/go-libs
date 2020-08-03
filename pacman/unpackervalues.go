//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "VUnpackerType=uint8,uint16,uint32,uint64,int8,int16,int32,int64,float32,float64"

package pacman

import (
	"encoding/binary"

	"github.com/cheekybits/genny/generic"
)

// VUnpackerType is the integer type
type VUnpackerType generic.Type

func (unpacker *unpacker) ReadVUnpackerType() VUnpackerType {
	var ret VUnpackerType
	binary.Read(unpacker.reader, unpacker.byteOrder, &ret)
	return ret
}
