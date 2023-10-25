package main

import "testing"

func TestPblocks(t *testing.T) {
	for i := 0; i < 256; i++ {
		if PblockR(Pblock(uint8(i))) != uint8(i) {
			t.Errorf("Error with data %d in Pblocks", i)
		}
	}
}

func TestSblocks(t *testing.T) {
	for i := 0; i < 256; i++ {
		if SblockR(Sblock(uint8(i))) != uint8(i) {
			t.Errorf("Error with data %d in Sblocks", i)
		}
	}
}
