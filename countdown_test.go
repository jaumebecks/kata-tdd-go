package main

import (
	"bytes"
	"reflect"
	"testing"
)

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
	t.Run("3 2 1 Go! sequence is printed", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &CountdownOperationsSpy{})

		want := "3\n2\n1\nGo!"
		got := buffer.String()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

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
