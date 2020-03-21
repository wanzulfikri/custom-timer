package main

import (
	"os"
	"testing"
	"time"
)

func TestRunTimerWithValidIntegers(t *testing.T) {
	timeModifier = time.Nanosecond
	playSound = false

	t.Run("ValidIntegers", func(t *testing.T) {
		os.Args = []string{"", "1", "1"}
		err := runTimer()
		if err != nil {
			t.Errorf("Positive integers should be valid. Error: %v", err)
		}
	})

	t.Run("NegativeIntegers", func(t *testing.T) {
		os.Args = []string{"", "-1", "-2"}
		err := runTimer()
		if err == nil {
			t.Error("Negative integers should return an error.")
		}
	})
}
