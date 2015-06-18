package main

import "fmt"

func main() {
	g1 := G{a: 10, b: 10}
	g2 := &G{a: 20, b: 20}
	f1 := F{a: 30, b: 30}
	f2 := &F{a: 40, b: 40}
	g1.calc()
	g2.calc()
	f1.calc()
	f2.calc()
	fmt.Println(g1.c)
	fmt.Println(g2.c)
	fmt.Println(f1.c)
	fmt.Println(f2.c)
}

type G struct {
	a int
	b int
	c int
}

func (g G) calc() {
	g.c = g.a + g.b
}

type F struct {
	a int
	b int
	c int
}

func (f *F) calc() {
	f.c = f.a + f.b
}
