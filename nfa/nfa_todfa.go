package nfa

import (
	"github.com/paulgriffiths/automata/dfa"
	"github.com/paulgriffiths/gods/sets"
)

// makeDtran builds a transition table that can be used to
// build a deterministic finite automaton.
func (n Nfa) makeDtran() dtran {
	ds := newDtran(n.EclosureS(n.Start))

	i := 0
	for i < ds.length() {
		for _, letter := range n.S.Elements() {
			nextState := n.EclosureT(n.Move(ds[i].nfaState, letter))
			if j, yes := ds.stateExists(nextState); yes {
				ds.addTrans(i, j, letter)
			} else {
				ds.appendState(nextState)
				ds.addTrans(i, ds.length()-1, letter)
			}
		}
		i++
	}

	return ds
}

// ToDfa converts a nondeterministic finite automaton to a
// deterministic finite automaton.
func (n Nfa) ToDfa() dfa.Dfa {
	ds := n.makeDtran()

	accepts := sets.NewSetInt()
	tfunc := []map[rune]int{}

	for i := 0; i < ds.length(); i++ {
		if !n.Accept.Intersection(ds[i].nfaState).IsEmpty() {
			accepts.Insert(i)
		}
		tfunc = append(tfunc, ds[i].trans)
	}

	return dfa.Dfa{ds.length(), n.S, tfunc, 0, accepts}
}
