package test

import (
	"fmt"
	"testing"
)

func TestIneffAssign(t *testing.T) {
	m := make(map[string]*p)
	m["a"] = &p{x: 100}
	f1(m)
	fmt.Println(m["a"].x)
}

type p struct {
	x int
}

func f1(m map[string]*p) {
	newM := make(map[string]*p)
	newM["a"] = &p{x: 1}
	newM["b"] = &p{x: 2}
	m = newM
}
