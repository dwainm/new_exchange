package main

import "fmt"

type Order struct {
	product   string
	orderType string
	qty       int
	amount    int
}

func newOrder(p string, ot string, q int, a int) *Order {
	o := Order{product: p, orderType: ot, qty: q, amount: a}
	return &o
}

func main() {
	fmt.Println("hello exchange")
}
