package fastrand

import (
	"testing"
)

func TestRNGUint32(t *testing.T) {
	var r RNG
	m := make(map[uint32]struct{})
	for i := 0; i < 1e6; i++ {
		n := r.Uint32()
		if _, ok := m[n]; ok {
			t.Fatalf("number %v already exists", n)
		}
		m[n] = struct{}{}
	}
}

func TestRNGSeed(t *testing.T) {
	var r RNG
	for _, seed := range []uint32{1, 2, 1234, 432432, 34324432} {
		r.Seed(seed)
		m := make(map[uint32]struct{})
		for i := 0; i < 1e6; i++ {
			n := r.Uint32()
			if _, ok := m[n]; ok {
				t.Fatalf("number %v already exists", n)
			}
			m[n] = struct{}{}
		}
	}
}
