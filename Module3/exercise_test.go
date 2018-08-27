package src_test

import (
	"fmt"
	"strings"
	"testing"
)

// Write a test for strings.HasPrefix
// https://golang.org/pkg/strings/#HasPrefix
// Given the value "main.go", test that it has the prefix "main"
// Remember that your test has to start with the name `Test` and be in an `_test.go` file
func Test_HasPrefix(t *testing.T) {
	value := "main.go"
	expected := "main"
	if !strings.HasPrefix(value, expected) {
		t.Fatalf("expected %s to have suffix %s", value, expected)
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

func Test_IndexTest(t *testing.T) {
	tt := []struct {
		Value     string
		Substring string
		Expected  int
	}{
		{Value: "Gophers are amazing!", Substring: "are", Expected: 8},
		{Value: "Testing in Go is fun.", Substring: "fun", Expected: 17},
		{Value: "The answer is 42.", Substring: "is", Expected: 11},
	}

	for _, x := range tt {

		result := strings.Index(x.Value, x.Substring)
		if result != x.Expected {
			t.Errorf("expected %d, got %d", x.Expected, result)
		}
	}
}

// Rewrite the above test for strings.Index using subtests
func Test_IndexTest_WithSub(t *testing.T) {
	tt := []struct {
		Value     string
		Substring string
		Expected  int
	}{
		{Value: "Gophers are amazing!", Substring: "are", Expected: 8},
		{Value: "Testing in Go is fun.", Substring: "fun", Expected: 17},
		{Value: "The answer is 42.", Substring: "is", Expected: 11},
	}

	for i, x := range tt {
		t.Run(fmt.Sprintf("sub test (%d)", i), func(st *testing.T) {
			result := strings.Index(x.Value, x.Substring)
			if result != x.Expected {
				t.Errorf("expected %d, got %d", x.Expected, result)
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
