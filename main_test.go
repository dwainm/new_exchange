package main

import "testing"

func TestPlaceOrderConfirmationIdGeneration(t *testing.T) {
	ob := &OrderBook{}
	o := newOrder("ZZZ", SELL, 5, 200)
	got := placeOrder(o, ob).confirmation
	want := "1-SELL-ZZZ"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestPlaceOrderInvalidType(t *testing.T) {
	ob := &OrderBook{}
	o := newOrder("ZZZ", SELL, 5, 200)
	got := placeOrder(o, ob).confirmation
	want := "1-SELL-ZZZ"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestNewOrder(t *testing.T) {
	got := newOrder("ZZZ", BUY, 5, 200)
	want := "ZZZ"
	if got.product != want {
		t.Errorf("got %q want %q", got.product, want)
	}
}
