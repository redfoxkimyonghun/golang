package main

import "fmt"

type IF interface {
	get() string
	set(string)
}

type Doc struct {
	text string
}

func (d Doc) get() string {
	return d.text
}

func (d Doc) set(s string) {
	d.text = s
}
