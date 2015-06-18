package main

import "fmt"

func main() {

	var a string = "abcde"
	var b string = a
	a = "ggg"
	fmt.Println(a)
	fmt.Println(b)

	var c string = "abcde"
	var d *string = &c
	//d := &c
	c = "ggg"
	fmt.Println(c)
	fmt.Println(*d)

}
