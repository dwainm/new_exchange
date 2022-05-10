package main

import (
	"fmt"
	"strconv"
	"time"
)

var orderBookIDIncrementor = 0

type Order struct {
	product   string
	orderType string
	qty       int
	amount    int
	timestamp int64
}

type OrderBookEntry struct {
	o            *Order
	confirmation string
}

// @todo possible issues with pass by value
func placeOrder(o *Order, ob []OrderBookEntry) OrderBookEntry {
	orderBookIDIncrementor++
	confirmation := strconv.Itoa(orderBookIDIncrementor) + "-" + o.orderType + "-" + o.product
	newEntry := OrderBookEntry{o, confirmation}
	ob = append(ob, newEntry)
	return newEntry
}

func newOrder(p string, ot string, q int, a int) *Order {
	o := Order{product: p, orderType: ot, qty: q, amount: a, timestamp: time.Now().Unix()}
	return &o
}

func main() {
	fmt.Println("hello exchange")
}
