package evaluator

func Distance(str1, str2 string) int {
	// Copyright (c) 2015, Arbo von Monkiewitsch All rights reserved.
	// Use of this source code is governed by a BSD-style
	// license that can be found in the LICENSE file.
	var cost, lastdiag, olddiag int
	s1 := []rune(str1)
	s2 := []rune(str2)

	len_s1 := len(s1)
	len_s2 := len(s2)

	column := make([]int, len_s1+1)

	for y := 1; y <= len_s1; y++ {
		column[y] = y
	}

	for x := 1; x <= len_s2; x++ {
		column[0] = x
		lastdiag = x - 1
		for y := 1; y <= len_s1; y++ {
			olddiag = column[y]
			cost = 0
			if s1[y-1] != s2[x-1] {
				cost = 1
			}
			column[y] = min(
				column[y]+1,
				column[y-1]+1,
				lastdiag+cost)
			lastdiag = olddiag
		}
	}
	return column[len_s1]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}

const MISS_THRESHOLD = 10

func TypoSuggestions(keys []string, miss string) []string {
	matches := make(map[string]int)
	for _, k := range keys {
		matches[k] = Distance(k, miss)
	}

	display := []string{}
	for k := range matches {
		if matches[k] < MISS_THRESHOLD {
			display = append(display, k)
		}
	}
	return display
}
