package main

import (
	"fmt"
	"strconv"
	"time"
)

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
	id           int
	o            *Order
	confirmation string
}

type OrderBook struct {
	idIncrementor int
	entries       []OrderBookEntry
}

// @todo possible issues with pass by value
func placeOrder(o *Order, ob *OrderBook) OrderBookEntry {
	ob.idIncrementor++
	confirmation := strconv.Itoa(ob.idIncrementor) + "-" + o.orderType.String() + "-" + o.product
	newEntry := OrderBookEntry{ob.idIncrementor, o, confirmation}
	ob.entries = append(ob.entries, newEntry)
	return newEntry
}

func newOrder(p string, ot OrderType, q int, a int) *Order {
	o := Order{product: p, orderType: ot, qty: q, amount: a, timestamp: time.Now().Unix()}
	return &o
}

func main() {
	fmt.Println("hello exchange")
}
