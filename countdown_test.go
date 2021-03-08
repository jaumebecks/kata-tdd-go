package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	Countdown(buffer)

	want := "3\n2\n1\nGo!"
	got := buffer.String()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
