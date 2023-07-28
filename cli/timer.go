package cli

import (
	"fmt"
	"time"
)

type Stopwatch struct {
    startTime time.Time
    isRunning bool
    interval  time.Duration
}

func NewStopWatch(interval time.Duration) Stopwatch {
    return Stopwatch{
        startTime: time.Now(),
        isRunning: false,
        interval:  interval,
    }
}

func (t *Stopwatch) Start() {
    t.startTime = time.Now()
    t.isRunning = true
    fmt.Print("\0337") // save cursor position
    go t.loop()
}

func (t *Stopwatch) SetStartTimeAndStart(time time.Time) {
    t.startTime = time
    t.isRunning = true
    fmt.Print("\0337") // save cursor position
    go t.loop()
}

func (t *Stopwatch) Stop() {
    t.isRunning = false
    fmt.Printf("\n")
}

func (t *Stopwatch) Reset() {
    t.startTime = time.Now()
}

func (t *Stopwatch) loop() {
    for t.isRunning {
        t.Update()
        time.Sleep(t.interval)
   }
}

func (t *Stopwatch) Update() {
    elapsed := time.Since(t.startTime).Truncate(t.interval).String()

    fmt.Print("\0338") // restore cursor position
    fmt.Printf("\033[90m%s\033[0m", elapsed)
}
