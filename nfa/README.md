# nfa

This package simulates a nondeterministic finite automaton (NFA).

Formally, an NFA is a 5-tuple (𝑄, 𝛴, 𝛿, 𝑞𝟢, 𝐹) where:

* 𝑄 is a finite set called the states;
* 𝛴 is a finite set called the alphabet;
* 𝛿 : 𝑄 × 𝛴𝜀 ⟶  𝒫(𝑄), is the transition function;
* 𝑞𝟢 ∈ 𝑄 is the start state; and
* 𝐹 ⊆ 𝑄 is the set of accept states.

and our Go implementation follows this definition closely:

```go
import "github.com/paulgriffiths/gods/sets"

type Nfa struct {
	Q  int                    // Number of states
	S  sets.SetRune           // Alphabet
	D  []map[rune]sets.SetInt // Transition function
	Qs int                    // Start state
	F  sets.SetInt            // Set of accepting states
}
```

The `Accepts` method then checks if a string is accepted by the NFA. The
`ToDfa` method converts the NFA to an equivalent DFA. Other methods allow
the construction of the union, concatenation, and closure of multiple NFAs,
enabling the construction of NFAs which match arbitrary regular expressions.

### Example

The following NFA recognizes strings matching aa*|bb*:

![nfa](https://user-images.githubusercontent.com/5059971/46050076-00a87180-c100-11e8-9c9b-4bcd63335306.png)

and we can implement it as follows:

```go
automaton := nfa.Nfa{
    Q: 5,
    S: sets.NewSetRune('a', 'b'),
    D: []map[rune]sets.SetInt{
        {0: sets.NewSetInt(1, 3)},
        {'a': sets.NewSetInt(2)},
        {'a': sets.NewSetInt(2)},
        {'b': sets.NewSetInt(4)},
        {'b': sets.NewSetInt(4)},
    },
    Qs: 0,
    F: sets.NewSetInt(2, 4),
}

for _, s := range []string{"aaaa", "bbb", "aaabba"} {
    if automaton.Accepts(s) {
        fmt.Printf("NFA accepts string %q.\n", s)
    } else {
        fmt.Printf("NFA does not accept string %q.\n", s)
    }
}

dfa := automaton.ToDfa()

for _, s := range []string{"aaabba", "baaab"} {
    accepted, n := dfa.AcceptsPrefix(s)
    if accepted {
        fmt.Printf("DFA accepts prefix %q of string %q.\n", s[:n], s)
    } else {
        fmt.Printf("DFA does not accept any prefix of string %q.\n", s)
    }
}

// Output:
// NFA accepts string "aaaa".
// NFA accepts string "bbb".
// NFA does not accept string "aaabba".
// DFA accepts prefix "aaa" of string "aaabba".
// DFA accepts prefix "b" of string "baaab".
```

Note the use of `0` in the transition function to denote 𝜀-transitions.
