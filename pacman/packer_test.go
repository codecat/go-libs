package pacman

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func assertPanic(t *testing.T, name string, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("%s code did not panic", name)
		}
	}()
	f()
}

func TestNewPacker(t *testing.T) {
	packer, err := NewPacker(new(bytes.Buffer))
	if err != nil {
		t.Errorf("creating packer got unexpected error: %s", err.Error())
	}

	if packer == nil {
		t.Error("packer must not be nil")
	}
}

func TestNewPackerNil(t *testing.T) {
	packer, err := NewPacker(nil)
	if err == nil {
		t.Error("was expecting an error, got nil")
	}

	if packer != nil {
		t.Error("packer must be nil")
	}
}

func TestPacker(t *testing.T) {
	buffer := new(bytes.Buffer)
	packer, _ := NewPacker(buffer)

	packer.Write((uint32)(0))
	if buffer.Len() != 4 {
		t.Error("unexpected buffer length")
	}

	packer.WriteUint8(0)
	if buffer.Len() != 5 {
		t.Error("unexpected buffer length")
	}

	packer.WriteUint16(0)
	if buffer.Len() != 7 {
		t.Error("unexpected buffer length")
	}

	packer.WriteUint32(0)
	if buffer.Len() != 11 {
		t.Error("unexpected buffer length")
	}

	packer.WriteUint64(0)
	if buffer.Len() != 19 {
		t.Error("unexpected buffer length")
	}

	packer.WriteInt8(0)
	if buffer.Len() != 20 {
		t.Error("unexpected buffer length")
	}

	packer.WriteInt16(0)
	if buffer.Len() != 22 {
		t.Error("unexpected buffer length")
	}

	packer.WriteInt32(0)
	if buffer.Len() != 26 {
		t.Error("unexpected buffer length")
	}

	packer.WriteInt64(0)
	if buffer.Len() != 34 {
		t.Error("unexpected buffer length")
	}

	packer.WriteFloat32(1.234)
	if buffer.Len() != 38 {
		t.Error("unexpected buffer length")
	}

	packer.WriteFloat64(1.234)
	if buffer.Len() != 46 {
		t.Error("unexpected buffer length")
	}

	packer.WriteString("hi")
	if buffer.Len() != 52 {
		t.Error("unexpected buffer length")
	}

	b, p := NewBlock()
	p.WriteInt32(0)
	packer.WriteBlock(b)
	if buffer.Len() != 60 {
		t.Error("unexpected buffer length")
	}

	pp := packer.BeginBlock()
	pp.WriteInt32(0)
	packer.EndBlock()

	if buffer.Len() != 68 {
		t.Error("unexpected buffer length")
	}

	assertPanic(t, "EndBlock", packer.EndBlock)
}

func TestPackerByteOrder(t *testing.T) {
	buffer := new(bytes.Buffer)
	packer, _ := NewPacker(buffer)
	packer.WriteUint32(0x12345678)
	packer.SetByteOrder(binary.BigEndian)
	packer.WriteUint32(0x12345678)

	res := bytes.Compare(
		buffer.Bytes(),
		[]byte{
			0x78, 0x56, 0x34, 0x12, // Little Endian
			0x12, 0x34, 0x56, 0x78, // Big Endian
		},
	)
	if res != 0 {
		t.Error("byte order doesn't work")
	}
}
