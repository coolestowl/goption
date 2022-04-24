package goption_test

import (
	"testing"

	"github.com/coolestowl/goption"
)

type Base struct {
	valA int
	valB int
}

func NewBase(opts ...goption.Option[Base]) *Base {
	b := new(Base)
	for _, opt := range opts {
		opt.Apply(b)
	}
	return b
}

func WithValA(val int) goption.Option[Base] {
	return goption.FuncOption(func(b *Base) {
		b.valA = val
	})
}

func WithValB(val int) goption.Option[Base] {
	return goption.FuncOption(func(b *Base) {
		b.valB = val
	})
}

func TestOption(t *testing.T) {
	b := NewBase(WithValA(1), WithValB(2))

	if b.valA != 1 {
		t.Errorf("Expected valA to be 1, got %d", b.valA)
	}

	if b.valB != 2 {
		t.Errorf("Expected valB to be 2, got %d", b.valB)
	}
}
