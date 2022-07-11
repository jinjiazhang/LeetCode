package leetcode

import "strings"

func findOcurrences(text string, first string, second string) []string {
	results := make([]string, 0)
	words := strings.Split(text, " ")
	for i := 0; i < len(words)-2; i++ {
		if words[i] == first && words[i+1] == second {
			results = append(results, words[i+2])
		}
	}

	return results
}
