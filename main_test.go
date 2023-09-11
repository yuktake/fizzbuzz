package main

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		{name: "", input: 1, want: "1"},
		{name: "", input: 2, want: "2"},
		{name: "3で割り切れる", input: 3, want: "fizz"},
		{name: "5で割り切れる", input: 5, want: "buzz"},
		{name: "15で割り切れる", input: 15, want: "fizzbuzz"},
		{name: "15で割り切れる", input: 30, want: "fizzbuzz"},
	}

	for _, tc := range tests {
		got := fizzbuzz(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("%s: expected: %s, got %s:", tc.name, tc.want, got)
		}
	}
}
