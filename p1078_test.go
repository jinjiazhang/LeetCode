package leetcode

import "testing"

func Test_1078(t *testing.T) {
	text := "alice is a good girl she is a good student"
	first := "a"
	second := "good"

	result := findOcurrences(text, first, second)
	t.Log(result)
}
