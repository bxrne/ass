package ass

// AutoInv wraps a state and auto-checks invariants
type AutoInv[T any] struct {
	state T
	suite InvSuite[T]
}

// NewAutoInv creates a new wrapper with invariants
func NewAuto[T any](state T, suite InvSuite[T]) *AutoInv[T] {
	a := &AutoInv[T]{state: state, suite: suite}
	a.check() // initial check
	return a
}

// Set replaces the state and auto-checks invariants
func (a *AutoInv[T]) Set(newState T) {
	a.state = newState
	a.check()
}

// Get returns the current state
func (a *AutoInv[T]) Get() T {
	return a.state
}

// Internal check function
func (a *AutoInv[T]) check() {
	for _, inv := range a.suite {
		if err := inv.Validate(a.state); err != nil {
			panic(err) // or log/fatal, depending on your policy
		}
	}
}
