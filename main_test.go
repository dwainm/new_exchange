package main

import "testing"

func TestPlaceOrderConfirmationIdGeneration(t *testing.T) {
	ob := make([]OrderBookEntry, 100, 100)
	o := newOrder("ZZZ", "sell", 5, 200)
	got := placeOrder(o, ob).confirmation
	want := "1-sell-ZZZ"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestNewOrder(t *testing.T) {
	got := newOrder("ZZZ", "bid", 5, 200)
	want := "ZZZ"
	if got.product != want {
		t.Errorf("got %q want %q", got.product, want)
	}
}
