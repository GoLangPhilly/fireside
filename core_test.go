package main

import (
	"testing"
)

func TestQuery(t *testing.T) {
	q := Query{Username: "djworth"}

	if q.Username != "djworth" {
		t.Errorf("expected q.Username to equal djworth, actual ", q.Username)
	}
}
