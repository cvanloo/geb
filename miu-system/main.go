package main

import (
	"strings"
)

const (
	SymM      = 'M'
	SymI      = 'I'
	SymU      = 'U'
	StringIII = "III"
	StringUU  = "UU"
)

func main() {

}

// RuleI if the string ends in an 'I', a 'U' can be appended.
func RuleI(s string) string {
	last := s[len(s)-1]
	if last == SymI {
		return s + string(SymU)
	}
	return ""
}

// RuleII the string Mx can be turned into Mxx.
func RuleII(s string) string {
	x := s[1:]
	xx := x + x
	mxx := string(s[0]) + xx
	return mxx
}

// RuleIII one occurrence of III can be replaced by U.
func RuleIII(s string) []string {
	parts := strings.Split(s, StringIII)
	strs := []string{}
	for i := 0; i < len(parts); i++ {
		ns := parts[0]
		other := parts[1:]

		if i < len(other) {
			for j, p := range other {
				if i == j {
					ns = ns + string(SymU) + p
				} else {
					ns = ns + p
				}
			}
			strs = append(strs, ns)
		}

	}
	return strs
}

// RuleIV one occurrence of UU can be removed from the string.
func RuleIV(s string) []string {
	parts := strings.Split(s, StringUU)
	strs := []string{}
	for i := 0; i < len(parts); i++ {
		ns := parts[0]
		for j, p := range parts[1:] {
			if j == i {
				ns = ns + p
			} else {
				ns = ns + StringUU + p
			}
		}
		if ns != s {
			strs = append(strs, ns)
		}
	}
	return strs
}
