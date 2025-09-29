
# ass

Runtime assertions for your Go programs.

Go provides `assert` only for testing, but it is often valuable to **test invariants** at runtime, especially in **simulation, stateful systems, or safety-critical code**.

`ass` is largely inspired by **Tiger Style** from TigerBeetle's dev team.  
[Read more about Tiger Style](https://github.com/tigerbeetle/tigerbeetle/blob/main/docs/TIGER_STYLE.md#safety)

---

## Installation

```bash
go get github.com/bxrne/ass
```

---

## Features

- Define **invariants** as reusable, named predicates over your state.  
- Group invariants into **suites** for convenient batch checking.  
- Fluent API to define invariants: `NewInv().Check().Msg()`.  
- **Automatic invariant checking** on state updates via `AutoInv`.  
- Optional error messages for clearer diagnostics.

---

## Usage

### Manual invariant checking

```go
package main

import (
    "fmt"

    "github.com/bxrne/ass"
)

type Counter struct {
    Value int
}

func main() {
    // Define an invariant
    inv := ass.NewInv[Counter]("NonNegativeCounter").
        Check(func(c Counter) bool { return c.Value >= 0 }).
        Msg("Counter value cannot go below zero")

    // Example state
    c := Counter{Value: -1}

    // Create a suite and check all invariants
    suite := ass.InvSuite[Counter]{inv}
    errs := suite.Check(c)
    for _, err := range errs {
        fmt.Println(err)
    }
}
```

**Output:**
```
Invariant NonNegativeCounter violated: Counter value cannot go below zero
```

---

### Automatic invariant checking

```go
package main

import (
    "github.com/bxrne/ass"
)

type Counter struct {
    Value int
}

func main() {
    inv := ass.NewInv[Counter]("NonNegative").
        Check(func(c Counter) bool { return c.Value >= 0 }).
        Msg("Counter cannot be negative")

    counter := Counter{Value: 0}

    // Wrap state for automatic invariant checking
    wrapped := ass.NewAutoInv(counter, ass.InvSuite[Counter]{inv})

    wrapped.Set(Counter{Value: 5})  // ✅ passes

    wrapped.Set(Counter{Value: -1}) // ⚠ panics (in automatic mode)
}
```

> Note: `AutoInv` automatically checks invariants whenever state is updated. By default, violations **panic**, but this behavior can be customized by providing an `OnViolate` callback.

---

## License

MIT License

