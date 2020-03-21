package main

import (
	"os"
	"testing"
	"time"
)

func TestRunTimerWithValidIntegers(t *testing.T) {
	timeModifier = time.Nanosecond
	playSound = false

	t.Run("Valid Integers", func(t *testing.T) {
		os.Args = []string{"", "1", "1"}
		err := runTimer()
		if err != nil {
			t.Errorf("Positive integers should be valid. Error: %v", err)
		}
	})

	t.Run("Negative integers", func(t *testing.T) {

	})

}
