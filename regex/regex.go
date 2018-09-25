package regex

import (
	"github.com/paulgriffiths/automata/dfa"
	"github.com/paulgriffiths/automata/nfa"
	"github.com/paulgriffiths/goeval/lar"
	"strings"
)

// Regex represents a compiled regular expression.
type Regex struct {
	d dfa.Dfa
}

// Match tests if the supplied string matches the regular expression.
func (r *Regex) Match(s string) bool {
	return r.d.Accepts(s)
}

// MatchPrefix tests if there is a prefix of the supplied string
// which matches the regular expression. If there is, it returns true
// and the length of the longest matching prefix. Otherwise, it returns
// false and zero.
func (r *Regex) MatchPrefix(s string) (bool, int) {
	return r.d.AcceptsPrefix(s)
}

// Compile compiles a regular expression provided in string form.
func Compile(r string) *Regex {
	lar, err := lar.NewLookaheadReader(strings.NewReader(r))
	if err != nil {
		return nil
	}

	expr := getExpr(&lar)
	if expr == nil || !lar.EndOfInput() {
		return nil
	}

	rx := Regex{(*expr).ToDfa()}
	return &rx
}

func getExpr(lar *lar.LookaheadReader) *nfa.Nfa {
	concat := getConcat(lar)
	if concat == nil {
		return nil
	}

	for lar.MatchOneOf('|') {
		next := getConcat(lar)
		if next == nil {
			return nil
		}
		temp := nfa.NewUnionNfa(*concat, *next)
		concat = &temp
	}

	return concat
}

func getConcat(lar *lar.LookaheadReader) *nfa.Nfa {
	closure := getClosure(lar)
	if closure == nil {
		return nil
	}
	for next := getClosure(lar); next != nil; next = getClosure(lar) {
		temp := nfa.NewConcatNfa(*closure, *next)
		closure = &temp
	}

	return closure
}

func getClosure(lar *lar.LookaheadReader) *nfa.Nfa {
	term := getTerm(lar)
	if term == nil {
		return nil
	}
	if lar.MatchOneOf('*') {
		closure := nfa.NewClosureNfa(*term)
		return &closure
	}
	return term
}

func getTerm(lar *lar.LookaheadReader) *nfa.Nfa {
	switch {
	case lar.MatchLetter():
		letter := nfa.NewRuneNfa(lar.Result.Value[0])
		return &letter
	case lar.MatchDigit():
		digit := nfa.NewRuneNfa(lar.Result.Value[0])
		return &digit
	case lar.MatchOneOf('('):
		expr := getExpr(lar)
		if expr == nil || !lar.MatchOneOf(')') {
			return nil
		}
		return expr
	}
	return nil
}
