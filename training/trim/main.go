package main

import "fmt"

type T string

func (t T) trim() {
	fmt.Println(t[:3])
}

func main() {
	var t T = "123456"
	t.trim()

}
