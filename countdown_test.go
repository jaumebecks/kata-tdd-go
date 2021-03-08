package main

import (
	"bytes"
	"testing"
)

type spySleeper struct {
	Calls int
}

func (s *spySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeper := &spySleeper{}
	Countdown(buffer, sleeper)

	want := "3\n2\n1\nGo!"
	got := buffer.String()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

	if sleeper.Calls != 4 {
		t.Errorf("got %d want 4 sleeper calls", sleeper.Calls)
	}
}
