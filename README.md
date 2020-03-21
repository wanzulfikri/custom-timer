# Work-rest timer with configurable time interval

A simple work-rest timer (ala Pomodoro) that allows for configurable time interval. Instead of sticking to a strict 25-5 - aka 25 work minutes and 5 rest minutes - interval, you can instead set it to 10-3, 20-1, or whatever you like.

Note: As of now, the intervals don't repeat themselves. If you set the interval to 10-3, the program will end after the second interval runs to completion. To have longer timer sessions, enter more inputs eg. 10-5-10-5.

## Prerequisites

[Go](https://golang.org/) - Go is needed to run the timer. 
[Beep](https://github.com/faiface/beep) - dependency for playing sounds

*Later, I'll create a release thus removing the need for the above prerequisites*

## Install

With Go installed, run the following command in your terminal:

`go get -u github.com/wanzulfikri/custom-timer`

## How to use

Let's say that you want an interval of 30 minutes and then 10 minutes.

To run, `cd` to the `custom-timer` directory, do either of the following:

i) Run immediately

`go run . 30 10`

ii) Build and run

`go build . && ./custom-timer 30 10`

iii) Install and run (make sure that [your path is configured correctly](https://golang.org/doc/install#install))

`go install . && custom-timer 30 10`

The intervals can be as long as you want and they must be integers. Examples:

**10 minutes**

`go run . 10`

**10 minutes -> 5 minutes**

`go run . 10 5`

**10 minutes -> 5 minutes -> 30 minutes -> 10 minutes**

`go run . 10 5 30 10`

**Error-inducing**

``` go run .
go run .

// or

go run . -10

// or

go run . "work yo!"
```

## Configuration

### Change sound 

The sound after each intervals are configurable.

There are three types of sounds:

 i)  work  - after every odd interval

ii)  rest  - after every even interval

iii) end   - at the end of program

First, put the sounds that you want to use into the same location as `main.go`.

Then, modify the following section of `main.go`:

```go
var sounds = map[string]string{
  // If you want to use "moo.wav" after work,
  // change "ohyeah.wav" to "moo.wav"
  "work": "moo.wav", 
	"rest": "ding.wav",
	"end":  "cheers.wav",
}
```



### Credits

A million thanks to the folks from [freesound.org](http://freesound.org) who shared the default sounds:
ohyeah.wav by [metrostock99](https://freesound.org/people/metrostock99/sounds/345086/)

ding.wav by [JohnsonBrandEditing](https://freesound.org/people/JohnsonBrandEditing/sounds/173932/)

cheers.wav by [FoolBoyMedia](https://freesound.org/people/FoolBoyMedia/sounds/397435/)