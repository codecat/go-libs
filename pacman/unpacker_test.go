package pacman

import (
	"bytes"
	"encoding/binary"
	"io"
	"testing"
)

func TestNewUnpacker(t *testing.T) {
	unpacker, err := NewUnpacker(new(bytes.Buffer))
	if err != nil {
		t.Errorf("creating unpacker got unexpected error: %s", err.Error())
	}

	if unpacker == nil {
		t.Error("unpacker must not be nil")
	}
}

func TestNewUnpackerNil(t *testing.T) {
	unpacker, err := NewUnpacker(nil)
	if err == nil {
		t.Error("was expecting an error, got nil")
	}

	if unpacker != nil {
		t.Error("unpacker must be nil")
	}
}

func TestUnpacker(t *testing.T) {
	buffer := new(bytes.Buffer)
	packer, _ := NewPacker(buffer)
	packer.WriteUint8(0x12)
	packer.WriteUint16(0x1234)
	packer.WriteUint32(0x12345678)
	packer.WriteUint64(0x1234567890ABCDEF)
	packer.WriteInt8(0x12)
	packer.WriteInt16(0x1234)
	packer.WriteInt32(0x12345678)
	packer.WriteInt64(0x1234567890ABCDEF)
	packer.WriteFloat32(2.5)
	packer.WriteFloat64(3.14159)
	packer.WriteString("Hello")
	p := packer.BeginBlock()
	p.WriteInt32(1)
	p.WriteInt32(2)
	p.WriteInt32(3)
	packer.EndBlock()
	packer.WriteString("World")

	unpacker, err := NewUnpacker(buffer)
	if err != nil {
		t.Errorf("unable to create unpacker: %s", err.Error())
	}

	if unpacker.ReadUint8() != 0x12 {
		t.Error("unexpected read value")
	}

	if unpacker.ReadUint16() != 0x1234 {
		t.Error("unexpected read value")
	}

	if unpacker.ReadUint32() != 0x12345678 {
		t.Error("unexpected read value")
	}

	if unpacker.ReadUint64() != 0x1234567890ABCDEF {
		t.Error("unexpected read value")
	}

	if unpacker.ReadInt8() != 0x12 {
		t.Error("unexpected read value")
	}

	if unpacker.ReadInt16() != 0x1234 {
		t.Error("unexpected read value")
	}

	if unpacker.ReadInt32() != 0x12345678 {
		t.Error("unexpected read value")
	}

	if unpacker.ReadInt64() != 0x1234567890ABCDEF {
		t.Error("unexpected read value")
	}

	if unpacker.ReadFloat32() != 2.5 {
		t.Error("unexpected read value")
	}

	if unpacker.ReadFloat64() != 3.14159 {
		t.Error("unexpected read value")
	}

	if unpacker.ReadString() != "Hello" {
		t.Error("unexpected read value")
	}

	b := unpacker.ReadBlock()

	if unpacker.ReadString() != "World" {
		t.Error("unexpected read value")
	}

	if b.ReadInt32() != 1 || b.ReadInt32() != 2 || b.ReadInt32() != 3 {
		t.Error("unexpected read value")
	}

	_, eof := buffer.ReadByte()
	if eof != io.EOF {
		t.Error("unexpected non-EOF found")
	}
}

func TestUnpackerByteOrder(t *testing.T) {
	buffer := new(bytes.Buffer)

	packer, _ := NewPacker(buffer)
	packer.WriteUint32(0x12345678)
	packer.SetByteOrder(binary.BigEndian)
	packer.WriteUint32(0x12345678)

	unpacker, _ := NewUnpacker(buffer)

	if unpacker.ReadUint32() != 0x12345678 {
		t.Error("unexpected read value")
	}

	unpacker.SetByteOrder(binary.BigEndian)

	if unpacker.ReadUint32() != 0x12345678 {
		t.Error("unexpected read value")
	}
}
