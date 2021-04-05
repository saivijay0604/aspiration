/*
	1 . Write a function with this signature:

	func CapitalizeEveryThirdAlphanumericChar(s string) string {
		// your code goes here
	}

	that capitalizes *only* every third alphanumeric character of a string, so that if I hand you

	Aspiration.com

	You hand me back

	asPirAtiOn.cOm

	Please note:

	- Characters other than each third should be downcased
	- For the purposes of counting and capitalizing every three characters, consider only alphanumeric
	  characters, ie [a-z][A-Z][0-9].

	2. Do the same problem, but this time create a "mapper" package that looks like this:
*/

package mapper

import (
	"strings"
)

type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

type Caps struct {
	I int
	S string
	R strings.Builder
}

var ct = 0

func MapString(i Interface) {
	for pos, _ := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}
}

func (c *Caps) GetValueAsRuneSlice() []rune {
	return []rune(c.S)
}

func (c *Caps) TransformRune(p int) {

	if (c.S[p] >= 'a' && c.S[p] <= 'z') || (c.S[p] >= 'A' && c.S[p] <= 'Z') || (c.S[p] >= '0' && c.S[p] <= '9') {
		if (p+1-ct)%3 == 0 {
			c.R.WriteString(strings.ToUpper(string(c.S[p])))
			return
		}
	} else {
		ct = ct + 1
	}
	c.R.WriteString(string(c.S[p]))
	return
}

// And change your code to look like this:

/*
  Make sure your fmt.Println(s) output looks nice and clean, ie implement the Stringer interface.
  Zip up your file(s) and email them to the recruiter
*/
