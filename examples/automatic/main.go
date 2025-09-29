package main

import "github.com/bxrne/ass"

type Counter struct {
	Value int
}

func main() {
	inv := ass.New[Counter]("NonNegative").
		Check(func(c Counter) bool { return c.Value >= 0 }).
		Msg("Counter cannot be negative")

	counter := Counter{Value: 0}
	wrapped := ass.NewAuto(counter, ass.InvSuite[Counter]{inv})

	wrapped.Set(Counter{Value: 5})  // ✅ passes
	wrapped.Set(Counter{Value: -1}) // ⚠ panics: Invariant NonNegative violated: Counter cannot be negative
}
