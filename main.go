package main

import (
	"aspiration/mapper"
	"fmt"
	"strings"
)

func main() {
	s := NewSkipString(3, "Aspiration.com")
	mapper.MapString(&s)
	fmt.Println("Transformed using mapper package : ", s.R.String())
	fmt.Println("Transformed using CapitalizeEveryThirdAlphanumericChar func:", CapitalizeEveryThirdAlphanumericChar("Aspiration.com"))
}

func NewSkipString(i int, s string) mapper.Caps {
	var r strings.Builder
	return mapper.Caps{i, strings.ToLower(s), r}
}

func CapitalizeEveryThirdAlphanumericChar(s string) string {
	// your code goes here
	s = strings.ToLower(s)
	r := ""
	c := 0
	for k, v := range s {

		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || (v >= '0' && v <= '9') {
			if (k+1-c)%3 == 0 {
				r = r + strings.ToUpper(string(v))
				continue
			}
		} else {
			c = c + 1
		}
		r = r + string(v)
	}

	return r
}
