package main

import (
	"fmt"
	"strconv"
)

type (
	Box__int struct {
		v int
	}
	Box__string struct {
		v string
	}
)

type A__int []int

type B__string__int map[string]int
type C__int__string func(int) string

type D__int struct {
	a int
	b A__int
}

func main() {

	box1 := Box__int{v: 123}
	fmt.Println(box1)

	box2 := Box__string{v: "gopher"}
	fmt.Println(box2)

	var a A__int

	a = append(a, 100)
	a = append(a, 200)
	fmt.Println(a)

	b := make(B__string__int)
	b["foo"] = 1
	b["bar"] = 2
	fmt.Println(b)

	cFunc := func(i int) string { return strconv.Itoa(i) }
	_ = C__int__string(cFunc)

	d := D__int{a: 300, b: a}
	fmt.Println(d)

}
