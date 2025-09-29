package ass_test

import (
	"testing"

	"github.com/bxrne/ass"
)

func TestInv_Validate_Success(t *testing.T) {
	inv := ass.New[int]("positive").Check(func(v int) bool { return v > 0 }).Msg("must be positive")
	if err := inv.Validate(5); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestInv_Validate_Failure(t *testing.T) {
	inv := ass.New[int]("positive").Check(func(v int) bool { return v > 0 }).Msg("must be positive")
	err := inv.Validate(-1)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	expected := "Invariant positive violated: must be positive"
	if err.Error() != expected {
		t.Fatalf("expected %q, got %q", expected, err.Error())
	}
}

func TestInv_Validate_NoCondition(t *testing.T) {
	inv := ass.New[int]("noCond")
	err := inv.Validate(0)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	expected := "Invariant noCond has no condition defined"
	if err.Error() != expected {
		t.Fatalf("expected %q, got %q", expected, err.Error())
	}
}

func TestInvSuite_Check(t *testing.T) {
	suite := ass.InvSuite[int]{
		ass.New[int]("positive").Check(func(v int) bool { return v > 0 }),
		ass.New[int]("even").Check(func(v int) bool { return v%2 == 0 }),
	}
	errs := suite.Check(4)
	if len(errs) != 0 {
		t.Fatalf("expected no errors, got %v", errs)
	}

	errs = suite.Check(3)
	if len(errs) != 1 || errs[0].Error() != "even invariant violated" {
		t.Fatalf("expected one error 'even invariant violated', got %v", errs)
	}
}
