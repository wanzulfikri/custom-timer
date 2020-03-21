package main

import (
	"os"
	"testing"
	"time"
)

func TestRunTimerWithValidIntegers(t *testing.T) {
	timeModifier = time.Nanosecond
	sleepModifier = time.Nanosecond
	playSound = false
	logIntervals = false

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

	t.Run("Non-Integers", func(t *testing.T) {
		os.Args = []string{"", "0.2", "string"}
		err := runTimer()
		if err == nil {
			t.Error("Non-integers should return an error.")
		}
	})

}
