package main

import (
	"bytes"
	"reflect"
	"testing"
)

type spySleeper struct {
	Calls int
}

func (s *spySleeper) Sleep() {
	s.Calls++
}

type CountdownOperationsSpy struct {
	Calls []string
}

func (c *CountdownOperationsSpy) Sleep() {
	c.Calls = append(c.Calls, sleepOperation)
}

func (c *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, writeOperation)
	return
}

const sleepOperation = "sleep"
const writeOperation = "write"

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

	t.Run("sleep is run before any write", func(t *testing.T) {
		countdownOperations := &CountdownOperationsSpy{}
		Countdown(countdownOperations, countdownOperations)
		want := []string{
			sleepOperation,
			writeOperation,
			sleepOperation,
			writeOperation,
			sleepOperation,
			writeOperation,
			sleepOperation,
			writeOperation,
		}
		got := countdownOperations.Calls

		if reflect.DeepEqual(want, got) == false {
			t.Errorf("got %v wanted %v", got, want)
		}
	})
}
