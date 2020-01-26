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

func main() {
	intervals := os.Args[1:]
	len := len(intervals)
	workBuffer, format := getBuffer("ohyeah.wav")
	restBuffer, _ := getBuffer("ding.wav")
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	isWorking := true
	buffer := workBuffer
	for i := 0; i < len; i++ {
		if isWorking {
			buffer = workBuffer
		} else {
			buffer = restBuffer
		}
		minute, err := strconv.Atoi(intervals[i])
		if err != nil {
			panic(err)
		}
		timer := time.NewTimer(time.Duration(minute) * time.Minute)
		sound := buffer.Streamer(0, buffer.Len())
		fmt.Printf("%v %v %v ", i, workType[isWorking], time.Now().Format(time.Kitchen))
		<-timer.C
		speaker.Play(sound)
		isWorking = !isWorking
	}
	time.Sleep(time.Second * 2)
	playOnce("cheers.wav")
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
