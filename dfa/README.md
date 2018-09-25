# dfa

This package simulates a deterministic finite automaton (DFA).

A DFA is represented as its formal 5-tuple definition:

```go
import "github.com/paulgriffiths/gods/sets"

type Dfa struct {
	Q      int            // Number of states
	S      sets.SetRune   // Alphabet
	D      []map[rune]int // Transition function
	Start  int            // Start state
	Accept sets.SetInt    // Set of accepting states
}
```

The `Accepts` method then checks if a string is accepted by the DFA. The
`AcceptsPrefix` method checks if the DFA accepts any prefix of a string,
with the longest prefix being preferred.
