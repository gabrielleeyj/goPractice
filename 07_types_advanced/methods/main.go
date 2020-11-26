package main

import "fmt"

type Printer interface {
	Print()
}

var (
	ErrNotConnected Error = "not-connected"
)

type Error string

var (
	StatusOnline  Status = "online"
	StatusOffline Status = "offline"
)

type Status string

func (s Status) Equal(status Status) bool {
	return s == status
}

func (s Status) Print() {
	fmt.Println("Status:", s)
}

func main() {
	fmt.Println(StatusOnline.Equal(StatusOffline))
	PrintAll(ErrNotConnected, StatusOffline)
}

func PrintAll(a, b Printer) {
	a.Print()
	b.Print()
}
