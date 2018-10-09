package main

import "fmt"

type Link struct {
	iData int
	dData float64
	next  *Link
}

func NewLink(i int, d float64) *Link {
	return &Link{
		iData: i,
		dData: d,
		next:  nil,
	}
}

func (l *Link) display() {
	fmt.Printf("{%d,%f}", l.iData, l.dData)
}

type DLink struct {
	data int
	next *DLink
}

func NewDLink(i int) *DLink {
	return &DLink{
		data: i,
		next: nil,
	}
}

func (l *DLink) display() {
	fmt.Printf("{%d}", l.data)
}
