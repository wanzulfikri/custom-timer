package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

var workType = map[bool]string{
	true:  "Work",
	false: "Rest",
}

// Set the sounds for the timer here.
//
// Three types of sounds:
//	i)  work  - after every odd interval
//  ii) rest  - after every even interval
//  ii) end   - at the end of program
var sounds = map[string]string{
	// modify according to your file name
	"work": "ohyeah.wav",
	"rest": "ding.wav",
	"end":  "cheers.wav",
}

// timeModifier will be multiplied by program input to give actual timer duration.
// In testing, modify this to time.Millisecond or time.Nanosecond to shorten
// test duration.
// Do the same for sleepModifier. Change it to time.Millisecond or time.Nanosecond
var timeModifier = time.Minute
var sleepModifier = time.Second
var playSound = true
var logIntervals = true

func main() {
	err := runTimer()
	if err != nil {
		log.Fatal(err)
	}
}

func runTimer() error {
	intervals := os.Args[1:]
	if len(intervals) == 0 {
		return errors.New("No timer interval added. Please rerun the timer with at least one interval")
	}

	err := checkValidity(intervals)
	if err != nil {
		return err
	}

	workBuffer, format, err := getBuffer(sounds["work"])
	restBuffer, _, err := getBuffer(sounds["rest"])
	if err != nil {
		return err
	}
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	isWorking := true
	buffer := workBuffer
	for i := 0; i < len(intervals); i++ {
		if isWorking {
			buffer = workBuffer
		} else {
			buffer = restBuffer
		}
		minute, err := strconv.Atoi(intervals[i])
		if err != nil {
			return err
		}
		timer := time.NewTimer(time.Duration(minute) * timeModifier)
		sound := buffer.Streamer(0, buffer.Len())
		if logIntervals {
			fmt.Printf("%v %v: %v -> %v\n", i, workType[isWorking], time.Now().Format(time.Kitchen), time.Now().Add(timeModifier*time.Duration(minute)).Format(time.Kitchen))
		}
		<-timer.C
		if playSound {
			speaker.Play(sound)
		}
		isWorking = !isWorking
	}
	time.Sleep(sleepModifier * 2)
	if playSound {
		playOnce(sounds["end"])
	}
	return nil
}

func checkValidity(intervals []string) error {
	for _, interval := range intervals {
		minute, err := strconv.Atoi(interval)
		if err != nil {
			return err
		} else if minute <= 0 {
			return fmt.Errorf("You entered %q. The interval must be more than 0", interval)
		}
	}
	return nil
}

func playOnce(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	streamer, _, err := wav.Decode(f)
	if err != nil {
		return err
	}

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
	time.Sleep(time.Second)

	defer streamer.Close()
	return nil
}

func getBuffer(filename string) (*(beep.Buffer), beep.Format, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, beep.Format{}, err
	}

	streamer, format, err := wav.Decode(f)
	if err != nil {
		return nil, beep.Format{}, err
	}
	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	streamer.Close()
	return buffer, format, nil
}
