package main

import (
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
//
//  Example: if program input is 10, 5, 10, then the sequence of sound would be
//	work -> rest -> work -> end.
var sounds = map[string]string{
	// modify according to your file name
	"work": "ohyeah.wav",
	"rest": "ding.wav",
	"end":  "cheers.wav",
}

// timeModifier will be multiplied by program input to give actual timer duration.
// In testing, modify this to time.Millisecond or time.Nanosecond to shorten
// test duration.
var timeModifier = time.Minute

func main() {
	intervals := os.Args[1:]
	if len(intervals) == 0 {
		log.Fatal("No timer interval added. Please rerun the timer with at least one interval")
	}

	err := checkValidity(intervals)
	if err != nil {
		log.Fatal(err)
	}

	workBuffer, format := getBuffer(sounds["work"])
	restBuffer, _ := getBuffer(sounds["rest"])
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
			panic(err)
		}
		timer := time.NewTimer(time.Duration(minute) * timeModifier)
		sound := buffer.Streamer(0, buffer.Len())
		fmt.Printf("%v %v: %v -> %v", i, workType[isWorking], time.Now().Format(time.Kitchen), time.Now().Add(timeModifier*time.Duration(minute)).Format(time.Kitchen))
		<-timer.C
		speaker.Play(sound)
		isWorking = !isWorking
	}
	time.Sleep(time.Second * 2)
	playOnce(sounds["end"])
}

func checkValidity(intervals []string) error {
	for _, interval := range intervals {
		minute, err := strconv.Atoi(interval)
		if err != nil {
			return fmt.Errorf("You entered %q. The interval must be an integer", interval)
		} else if minute <= 0 {
			return fmt.Errorf("You entered %q. The interval must be more than 0", interval)
		}
	}
	return nil
}

func playOnce(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	streamer, _, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
	time.Sleep(time.Second)

	defer streamer.Close()
}

func getBuffer(filename string) (*(beep.Buffer), beep.Format) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	streamer.Close()
	return buffer, format
}
