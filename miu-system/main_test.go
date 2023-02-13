package main

import "testing"

func TestRuleI(t *testing.T) {
	tests := []struct {
		test    string
		results map[string]bool
	}{
		{test: "MI", results: map[string]bool{"MIU": false}},
		{test: "MIII", results: map[string]bool{"MIIIU": false}},
		{test: "MIUUUIIU", results: map[string]bool{}},
	}

	for caseNum, testCase := range tests {
		res := RuleI(testCase.test)
		for _, r := range res {
			if _, ok := testCase.results[r]; !ok {
				t.Errorf("%d: invalid result: `%s'", caseNum, r)
			} else {
				testCase.results[r] = true
			}
		}
		for k, v := range testCase.results {
			if !v {
				t.Errorf("%d: missing result: `%s'", caseNum, k)
			}
		}
	}
}

func TestRuleII(t *testing.T) {
	tests := []struct {
		test    string
		results map[string]bool
	}{
		{test: "MI", results: map[string]bool{"MII": false}},
		{test: "MIII", results: map[string]bool{"MIIIIII": false}},
        {test: "MIUUUIIU", results: map[string]bool{"MIUUUIIUIUUUIIU": false}},
	}

	for caseNum, testCase := range tests {
		res := RuleII(testCase.test)
		for _, r := range res {
			if _, ok := testCase.results[r]; !ok {
				t.Errorf("%d: invalid result: `%s'", caseNum, r)
			} else {
				testCase.results[r] = true
			}
		}
		for k, v := range testCase.results {
			if !v {
				t.Errorf("%d: missing result: `%s'", caseNum, k)
			}
		}
	}
}

func TestRuleIII(t *testing.T) {
	tests := []struct {
		test    string
		results map[string]bool
	}{
		{test: "MIIIUU", results: map[string]bool{"MUUU": false}},
		{test: "MIII", results: map[string]bool{"MU": false}},
		{test: "MIIIUIII", results: map[string]bool{"MUUIII": false, "MIIIUU": false}},
		{test: "MIIIIII", results: map[string]bool{"MUIII": false, "MIIIU": false}},
		{test: "MIUUUIIU", results: map[string]bool{}},
	}

	for caseNum, testCase := range tests {
		res := RuleIII(testCase.test)
		for _, r := range res {
			if _, ok := testCase.results[r]; !ok {
				t.Errorf("%d: invalid result: `%s'", caseNum, r)
			} else {
				testCase.results[r] = true
			}
		}
		for k, v := range testCase.results {
			if !v {
				t.Errorf("%d: missing result: `%s'", caseNum, k)
			}
		}
	}
}

func TestRuleIV(t *testing.T) {
	tests := []struct {
		test    string
		results map[string]bool
	}{
		{test: "MIUU", results: map[string]bool{"MI": false}},
		{test: "MIUUIII", results: map[string]bool{"MIIII": false}},
		{test: "MIUUIUUI", results: map[string]bool{"MIIUUI": false, "MIUUII": false}},
		{test: "MIUUUIIU", results: map[string]bool{"MIUIIU": false}},
		{test: "MIUUUUIIU", results: map[string]bool{"MIUUIIU": false}},
		{test: "MIUIIU", results: map[string]bool{}},
	}

	for caseNum, testCase := range tests {
		res := RuleIV(testCase.test)
		for _, r := range res {
			if _, ok := testCase.results[r]; !ok {
				t.Errorf("%d: invalid result: `%s'", caseNum, r)
			} else {
				testCase.results[r] = true
			}
		}
		for k, v := range testCase.results {
			if !v {
				t.Errorf("%d: missing result: `%s'", caseNum, k)
			}
		}
	}
}

func TestSolve1(t *testing.T) {
	start := "MI"
	end := "MIU"
	expect := StepString{
		"MIU",
		[]StepDesc{StepStart, StepRuleI},
	}
	res := Solve(start, end, 1000)

	if res.Result != expect.Result {
		t.Errorf("want: %s, got invalid result: `%s'", expect.Result, res.Result)
	}
	for i, step := range expect.Steps {
		actualStep := res.Steps[i]
		if actualStep != step {
			t.Errorf("want: %s, got invalid step: `%s'", step, actualStep)
		}
	}
}

// Already takes much too long / too much memory
//func TestSolve2(t *testing.T) {
//	start := "MI"
//	end := "MUII"
//	expect := StepString{
//		"MUII",
//		[]StepDesc{StepStart, StepRuleI},
//	}
//	res := Solve(start, end, 1000)
//
//	if res.Result != expect.Result {
//		t.Errorf("want: %s, got invalid result: `%s'", expect.Result, res.Result)
//	}
//	for i, step := range expect.Steps {
//		actualStep := res.Steps[i]
//		if actualStep != step {
//			t.Errorf("want: %s, got invalid step: `%s'", step, actualStep)
//		}
//	}
//}
