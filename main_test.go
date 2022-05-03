package main

import "testing"

func TestNewOrder(t *testing.T) {
	got := newOrder("ZZZ", "bid", 5, 200)
	want := "ZZZ"
	if got.product != want {
		t.Errorf("got %q want %q", got.product, want)
	}
}
