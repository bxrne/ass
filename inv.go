package ass

import "fmt"

// Inv represents a software invariant testable by a condition
type Inv[T any] struct {
	Name      string
	Condition func(T) bool
	ErrMsg    string
}

// NewInv creates a new invariant with a name
func New[T any](name string) *Inv[T] {
	return &Inv[T]{Name: name}
}

// Check sets the condition for the invariant (fluent style)
func (inv *Inv[T]) Check(f func(T) bool) *Inv[T] {
	inv.Condition = f
	return inv
}

// Msg sets an optional error message
func (inv *Inv[T]) Msg(msg string) *Inv[T] {
	inv.ErrMsg = msg
	return inv
}

// Validate checks the invariant against the state
func (inv *Inv[T]) Validate(state T) error {
	if inv.Condition == nil {
		return fmt.Errorf("Invariant %s has no condition defined", inv.Name)
	}
	if !inv.Condition(state) {
		if inv.ErrMsg != "" {
			return fmt.Errorf("Invariant %s violated: %s", inv.Name, inv.ErrMsg)
		}
		return fmt.Errorf("%s invariant violated", inv.Name)
	}
	return nil
}

// InvSuite represents many invariants
type InvSuite[T any] []*Inv[T]

// Check checks all invariants in the suite
func (suite InvSuite[T]) Check(state T) []error {
	var errs []error
	for _, inv := range suite {
		if err := inv.Validate(state); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}
