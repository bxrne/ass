package ass_test

import (
	"fmt"
	"testing"

	"github.com/bxrne/ass"
)

func TestAutoInv_Get_Set(t *testing.T) {
	suite := ass.InvSuite[int]{ass.New[int]("positive").Check(func(v int) bool { return v > 0 })}
	a := ass.NewAuto(10, suite)

	if a.Get() != 10 {
		t.Fatalf("expected 10, got %v", a.Get())
	}

	a.Set(20)
	if a.Get() != 20 {
		t.Fatalf("expected 20, got %v", a.Get())
	}
}

func TestAutoInv_PanicOnViolation(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic on invariant violation, got none")
		} else {
			msg := fmt.Sprintf("%v", r)
			expected := "negative invariant violated"
			if msg != expected {
				t.Fatalf("expected panic message %q, got %q", expected, msg)
			}
		}
	}()

	suite := ass.InvSuite[int]{ass.New[int]("negative").Check(func(v int) bool { return v >= 0 })}
	a := ass.NewAuto(5, suite) // ok
	a.Set(-1)                  // triggers panic
}

func TestAutoInv_PanicOnNewViolation(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic on NewAuto with violating state, got none")
		} else {
			msg := fmt.Sprintf("%v", r)
			expected := "negative invariant violated"
			if msg != expected {
				t.Fatalf("expected panic message %q, got %q", expected, msg)
			}
		}
	}()

	suite := ass.InvSuite[int]{ass.New[int]("negative").Check(func(v int) bool { return v >= 0 })}
	_ = ass.NewAuto(-1, suite) // triggers panic on creation
}
