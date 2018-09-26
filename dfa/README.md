# dfa

This package simulates a deterministic finite automaton (DFA).

Formally, a DFA is a 5-tuple (𝑄, 𝛴, 𝛿, 𝑞𝟢, 𝐹) where:

* 𝑄 is a finite set called the states;
* 𝛴 is a finite set called the alphabet;
* 𝛿 : 𝑄 × 𝛴 ⟶  𝑄, is the transition function;
* 𝑞𝟢 ∈ 𝑄 is the start state; and
* 𝐹 ⊆ 𝑄 is the set of accept states.

and our Go implementation follows this definition closely:

```go
import "github.com/paulgriffiths/gods/sets"

type Dfa struct {
	Q  int            // Number of states
	S  sets.SetRune   // Alphabet
	D  []map[rune]int // Transition function
	Qs int            // Start state
	F  sets.SetInt    // Set of accepting states
}
```

The `Accepts` method then checks if a string is accepted by the DFA. The
`AcceptsPrefix` method checks if the DFA accepts any prefix of a string,
with the longest prefix being preferred.

### Example

The following DFA recognizes any string consisting solely of 'a's and 'b's
and which starts and ends with the same letter:

![dfa](https://user-images.githubusercontent.com/5059971/46049249-5af30380-c0fb-11e8-88a8-44b76edf8f4f.png)

and we can implement it as follows:

```go
automaton := dfa.Dfa{
    Q: 5,
    S: sets.NewSetRune('a', 'b'),
    D: []map[rune]int{
        {'a': 1, 'b': 3},
        {'a': 1, 'b': 2},
        {'a': 1, 'b': 2},
        {'a': 4, 'b': 3},
        {'a': 4, 'b': 3},
    },
    Qs: 0,
    F:  sets.NewSetInt(1, 3),
}

for _, s := range []string{"abbba", "abbb"} {
    if automaton.Accepts(s) {
        fmt.Printf("DFA accepts string %q.\n", s)
    } else {
        fmt.Printf("DFA does not accept string %q.\n", s)
    }
}

for _, s := range []string{"ababb", "abbbb"} {
    accepted, n := automaton.AcceptsPrefix(s)
    if accepted {
        fmt.Printf("DFA accepts prefix %q of string %q.\n", s[:n], s)
    } else {
        fmt.Printf("DFA does not accept any prefix of string %q.\n", s)
    }
}

// Output:
// DFA accepts string "abbba".
// DFA does not accept string "abbb".
// DFA accepts prefix "aba" of string "ababb".
// DFA accepts prefix "a" of string "abbbb".
```
