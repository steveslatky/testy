package main

import (
	"fmt"
	"testing"
)

const Version = "0.0.1"

type test struct {
	name string
	vars []Fv
}

// Formatd var
type Fv struct {
	name string
	T    any
}

// AssertEqual checks if two values are equal
func AssEq(t *testing.T, expected, actual interface{}) {
	t.Helper()
	if expected != actual {
		t.Errorf("Expected result: %d, Actual Result %d", expected, actual)
	}
}

func main(){
    fmt.Println("Worked")
}
