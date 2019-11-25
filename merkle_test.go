package merkle

import (
	"testing"

	"github.com/umran/cowrie/crypto"
)

type A struct {
}

type B struct {
}

func (a *A) Hash() crypto.Hash {
	return crypto.GenerateHash(make([]byte, 256))
}

func (b *B) Hash() crypto.Hash {
	return crypto.GenerateHash(make([]byte, 256))
}

func NewA() *A {
	return &A{}
}

func NewB() *B {
	return &B{}
}

func BenchmarkNew(b *testing.B) {
	// prepare input
	A1 := NewA()
	A2 := NewA()
	B1 := NewB()
	B2 := NewB()

	hashables := []Hashable{
		A1,
		A2,
		B1,
		B2,
	}

	for i := 0; i < b.N; i++ {
		New(hashables)
	}
}
