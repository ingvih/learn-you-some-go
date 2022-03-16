package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
    buffer := &bytes.Buffer{}
    observableSleeper := &ObservableSleeper{}

    Countdown(buffer, observableSleeper)

    got := buffer.String()
    want := `3
2
1
Go!`

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }

    if observableSleeper.Calls != 4 {
        t.Errorf("not enough calls to sleeper, want 4 got %d", observableSleeper.Calls)
    }
}

type ObservableSleeper struct {
    Calls int
}

func (s *ObservableSleeper) Sleep() {
    s.Calls++
}
