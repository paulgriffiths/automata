package nfa

import "github.com/paulgriffiths/gods/sets"

type dstate struct {
	nfaState sets.SetInt
	trans    map[rune]int
}

func newDstate(s sets.SetInt) dstate {
	return dstate{s, make(map[rune]int)}
}

type dtran []dstate

func newDtran(s sets.SetInt) dtran {
	return dtran{newDstate(s)}
}

func (d dtran) length() int {
	return len(d)
}

func (d *dtran) appendState(s sets.SetInt) {
	*d = append(*d, newDstate(s))
}

func (d dtran) addTrans(from, to int, a rune) {
	d[from].trans[a] = to
}

func (d dtran) stateExists(s sets.SetInt) (int, bool) {
	for i, state := range d {
		if state.nfaState.Equals(s) {
			return i, true
		}
	}
	return 0, false
}
