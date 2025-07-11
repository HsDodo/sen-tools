package test

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) Study() string {
	return fmt.Sprintf("%s正在学习", p.Name)
}

type Data struct {
	val int
}

func (d Data) AddOneByValue() {
	// 打印d的地址
	fmt.Printf("d的指针地址: %p\n", &d)
}
func (d *Data) AddOneByPointer() {
	// 打印d的地址
	fmt.Printf("d的指针地址: %p\n", d)
	d.val++
}

func TestF1(t *testing.T) {

	var d Data
	fmt.Printf("d的指针地址: %p\n", &d)
	d.AddOneByValue()
	d.AddOneByPointer()
	fmt.Println(d.val)
}

func TTT(student IStudent) {
	fmt.Println(student.Study())
}

func TTT3(student IStudent) {
	fmt.Println(student.Study())
}

func F1(p ...*Person) {
	// 打印p的实际地址
	fmt.Printf("F1中 p的地址: %p\n", &p)
}

type IStudent interface {
	Study() string
}
