# match

**match** demonstrates the simple regular expression package.

## Notes

The simple regular expression package is an exercise in generating finite
automata from regular expressions in string format, and the functionality
of the usable regular expressions is therefore relatively limited.

* Any letter or number may be used as a matching character

* There are no wildcard characters or character classes

* A match is returned only if the entire string matches, rather than any
substring

* Aside from concatenation which requires no special characters, the
Kleene star or closure (*) and union (|) operators are available

* The precedences of operations, from highest to lowest, is closure,
concatenation, then union

* Parentheses may be used and nested to any depth for grouping or for
overriding default operation predecence

## Usage examples

	paul@horus:match$ ./match '0*|1*' numbers.txt
	0
	1
	00
	11
	000
	111
	0000
	1111
	00000
	11111
	000000
	111111
	paul@horus:match$ ./match '1*|b*' numbers.txt letters.txt
	1
	11
	111
	1111
	11111
	111111
	b
	bb
	bbb
	bbbb
	bbbbb
	bbbbbb
	paul@horus:match$ cat numbers.txt | ./match '1101(0|1)*'
	1101
	11010
	11011
	110100
	110101
	110110
	110111
	paul@horus:match$ ./match '(aa|bb)*' letters.txt
	aa
	bb
	aaaa
	aabb
	bbaa
	bbbb
	aaaaaa
	aaaabb
	aabbaa
	aabbbb
	bbaaaa
	bbaabb
	bbbbaa
	bbbbbb
	paul@horus:match$ 
