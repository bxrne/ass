package main

import (
	"fmt"
	"github.com/bxrne/ass"
)

type Counter struct {
	Value int
}

func main() {
	inv := ass.New[Counter]("NonNegativeCounter").
		Check(func(c Counter) bool { return c.Value >= 0 }).
		Msg("Counter value cannot go below zero")

	c := Counter{Value: -1}

	suite := ass.InvSuite[Counter]{inv}

	for _, err := range suite.Check(c) {
		fmt.Println(err)
	}
}
