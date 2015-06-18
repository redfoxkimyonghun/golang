package main

import "fmt"

func main() {
	g1 := G{a: 10, b: 10}
	g2 := &G{a: 20, b: 20}
	f1 := F{a: 30, b: 30}
	f2 := &F{a: 40, b: 40}
	fmt.Println(g1.sum())
	fmt.Println(g2.sum())
	fmt.Println(f1.sum())
	fmt.Println(f2.sum())
}

type G struct {
	a int
	b int
}

func (g G) sum() int {
	return g.a + g.b
}

type F struct {
	a int
	b int
}

func (f *F) sum() int {
	return f.a + f.b
}
