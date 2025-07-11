package test

import (
	"fmt"
	"testing"
)

type A struct {
	A int
}

type B struct {
	B int
}

type IObj interface {
	TTT()
}

func (a *A) AAA() bool {
	return true
}

func (a *A) TTT() {
}

func (b *B) BBB() bool {
	return true
}

func (b *B) TTT() {

}

func Test_S1034(t *testing.T) {
	// a := &A{A: 1}
	b := &B{B: 2}
	var c IObj
	c = b
	switch v := c.(type) {
	case *A:
		if v.AAA() {
			fmt.Println("A")
		}
	case *B:
		if v.BBB() {
			fmt.Println("B")
		}
	}
}
