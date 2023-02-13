package main

import (
	"strings"
	"fmt"
	"flag"
	"log"
	"runtime"
)

const (
	SymM      = 'M'
	SymI      = 'I'
	SymU      = 'U'
	StringIII = "III"
	StringUU  = "UU"
)

type StepDesc string

var (
	StepStart   StepDesc = "start"
	StepRuleI   StepDesc = "ruleI"
	StepRuleII  StepDesc = "ruleII"
	StepRuleIII StepDesc = "ruleIII"
	StepRuleIV  StepDesc = "ruleIV"
)

type StepString struct {
	Result string
	Steps  []StepDesc
}

var startString = flag.String("s", "MI", "The initial string") 
var endString = flag.String("e", "MU", "The target string")
var maxDepth = flag.Int("d", 1000, "Maximal number of steps to apply to a string")

func main() {
	flag.Parse()
	if len(*startString) < 2 {
		panic("initial string must have a length of at least two characters")
	}
	log.Printf("start: `%s', end: `%s', max: %d", *startString, *endString, *maxDepth)
	res := Solve(*startString, *endString, *maxDepth)
	fmt.Printf("TOOK %d STEPS\n", len(res.Steps))
	for _, s := range res.Steps {
		fmt.Printf("%s -> ", s)
	}
	fmt.Println("RESULT FOUND!")
}

func Solve(startString, endString string, maxDepth int) StepString {
	coll := []StepString{
		{startString, []StepDesc{StepStart}},
	}
	for {
		ncoll := []StepString{}
		for _, s := range coll {
			r1 := CreateDesc(RuleI(s.Result), s.Steps, StepRuleI)
			r2 := CreateDesc(RuleII(s.Result), s.Steps, StepRuleII)
			r3 := CreateDesc(RuleIII(s.Result), s.Steps, StepRuleIII)
			r4 := CreateDesc(RuleIV(s.Result), s.Steps, StepRuleIV)
			for _, r := range r1 {
				if len(r.Steps) < maxDepth {
					ncoll = append(ncoll, r)
				}
			}
			for _, r := range r2 {
				if len(r.Steps) < maxDepth {
					ncoll = append(ncoll, r)
				}
			}
			for _, r := range r3 {
				if len(r.Steps) < maxDepth {
					ncoll = append(ncoll, r)
				}
			}
			for _, r := range r4 {
				if len(r.Steps) < maxDepth {
					ncoll = append(ncoll, r)
				}
			}
			//ncoll = append(ncoll, r1...)
			//ncoll = append(ncoll, r2...)
			//ncoll = append(ncoll, r3...)
			//ncoll = append(ncoll, r4...)
		}
		for _, c := range ncoll {
			if c.Result == endString {
				return c
			}
		}
		if len(ncoll) == 0 {
			log.Printf("No solution")
			var nop StepString
			return nop
		}
		coll = ncoll
		log.Printf("Permutations: %d", len(coll))
		PrintMemUsage()
	}
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}

func CreateDesc(ss []string, prevSteps []StepDesc, currStep StepDesc) []StepString {
	ncoll := []StepString{}
	for _, s := range ss {
		numPrevSteps := len(prevSteps)
		nsteps := make([]StepDesc, numPrevSteps+1)
		for i, li := range prevSteps {
			nsteps[i] = li
		}
		nsteps[numPrevSteps] = currStep
		ncoll = append(ncoll, StepString{s, nsteps})
	}
	return ncoll
}

// RuleI if the string ends in an 'I', a 'U' can be appended.
func RuleI(s string) []string {
	last := s[len(s)-1]
	if last == SymI {
		return []string{s + string(SymU)}
	}
	return []string{}
}

// RuleII the string Mx can be turned into Mxx.
func RuleII(s string) []string {
	x := s[1:]
	mxx := string(s[0]) + x + x
	return []string{mxx}
}

// RuleIII one occurrence of III can be replaced by U.
func RuleIII(s string) []string {
	parts := strings.Split(s, StringIII)
	strs := []string{}
	start := parts[0]
	rest := parts[1:]
	for i := 0; i < len(rest); i++ {
		ns := start
		for j, p := range rest {
			if i == j {
				ns = ns + string(SymU) + p
			} else {
				ns = ns + StringIII +  p
			}
		}
		strs = append(strs, ns)
	}
	return strs
}

// RuleIV one occurrence of UU can be removed from the string.
func RuleIV(s string) []string {
	parts := strings.Split(s, StringUU)
	strs := []string{}
	start := parts[0]
	rest := parts[1:]
	for i := 0; i < len(rest); i++ {
		ns := start
		for j, p := range rest {
			if j == i {
				ns = ns + p
			} else {
				ns = ns + StringUU + p
			}
		}
		strs = append(strs, ns)
	}
	return strs
}
