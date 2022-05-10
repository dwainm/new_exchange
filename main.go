package main

import (
	"fmt"
	"strconv"
	"time"
)

var orderBookIDIncrementor = 0

type OrderType string

const (
	BUY  OrderType = "BUY"
	SELL           = "SELL"
)

func (t OrderType) String() string {
	if t == BUY {
		return "BUY"
	} else if t == SELL {
		return "SELL"
	} else {
		return "unknown"
	}
}

type Order struct {
	product   string
	orderType OrderType
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
	confirmation := strconv.Itoa(orderBookIDIncrementor) + "-" + o.orderType.String() + "-" + o.product
	newEntry := OrderBookEntry{o, confirmation}
	ob = append(ob, newEntry)
	return newEntry
}

func newOrder(p string, ot OrderType, q int, a int) *Order {
	o := Order{product: p, orderType: ot, qty: q, amount: a, timestamp: time.Now().Unix()}
	return &o
}

func main() {
	fmt.Println("hello exchange")
}
