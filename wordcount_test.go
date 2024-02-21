package main

import (
	"reflect"
	"testing"
	"fmt"
)

func TestWordCount(t *testing.T) {
	tests := []struct {
		input string        // input is the string to be tested
		want  map[string]int // want is the expected output map
	}{
		{
			input: "hello world hello",
			want:  map[string]int{"hello": 2, "world": 1},
		},
		{
			input: "apple banana apple banana",
			want:  map[string]int{"apple": 2, "banana": 2},
		},
		{
			input: "one two three",
			want:  map[string]int{"one": 1, "two": 1, "three": 1},
		},
	}

	for i, test := range tests {
		got := WordCount(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Test case %d failed. Input: %q, Got: %v, Want: %v", i+1, test.input, got, test.want)
		} else {
			fmt.Printf("Test case %d passed. Input: %q, Got: %v, Want: %v\n", i+1, test.input, got, test.want)
		}
	}
}
