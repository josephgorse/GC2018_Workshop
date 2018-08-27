package src_test

import (
	"strings"
	"testing"
)

// Write a test for strings.HasPrefix
// https://golang.org/pkg/strings/#HasPrefix
// Given the value "main.go", test that it has the prefix "main"
// Remember that your test has to start with the name `Test` and be in an `_test.go` file
func Test_HasPrefix(t *testing.T) {
	value := "main.go"
	if !strings.HasPrefix(value, "main") {
		t.Fatalf("expected %s to have suffix %s", value, "main")
	}
}

// Write a table drive test for strings.Index
// https://golang.org/pkg/strings/#Index
// Use the following test conditions
// |------------------------------------------------|
// | Value                     | Substring | Answer |
// |===========================|===========|========|
// | "Gophers are amazing!"    | "are"     | 8      |
// | "Testing in Go is fun."   | "fun"     | 17     |
// | "The answer is 42."       | "is"      | 11     |
// |------------------------------------------------|
func Test_Index(t *testing.T) {
	tests := []struct {
		value  string
		substr string
		exp    int
	}{
		{"Gophers are amazing", "are", 8},
		{"Testing in Go is fun.", "fun", 17},
		{"The answer is 42.", "is", 11},
	}
	for _, test := range tests {
		got := strings.Index(test.value, test.substr)
		if got != test.exp {
			t.Errorf("unexpected value.  got %d, exp %d", got, test.exp)
		}
	}
}

// Rewrite the above test for strings.Index using subtests
func Test_Index_Subtest(t *testing.T) {
	tests := []struct {
		value  string
		substr string
		exp    int
	}{
		{"Gophers are amazing", "are", 8},
		{"Testing in Go is fun.", "fun", 17},
		{"The answer is 42.", "is", 11},
	}
	for _, test := range tests {
		t.Run(test.value, func(st *testing.T) {
			got := strings.Index(test.value, test.substr)
			if got != test.exp {
				st.Errorf("unexpected value.  got %d, exp %d", got, test.exp)
			}
		})
	}
}

// Here is a simple test that tests `strings.HasSuffix`.i
// https://golang.org/pkg/strings/#HasSuffix
func Test_HasSuffix(t *testing.T) {
	value := "main.go"
	if !strings.HasSuffix(value, ".go") {
		t.Fatalf("expected %s to have suffix %s", value, ".go")
	}
}
