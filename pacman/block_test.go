package pacman

import "testing"

func TestBlock(t *testing.T) {
	b, p := NewBlock()

	p.WriteUint32(0)
	if b.Len() != 4 {
		t.Error("unexpected block length")
	}
}
