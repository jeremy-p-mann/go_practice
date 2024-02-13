package main

import (
	"fmt"
	"time"
)

func main() {
	duration := 100 * time.Millisecond
	ticker := time.NewTicker(duration)
	done := make(chan bool)
	ticks := []time.Time{}
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				ticks = append(ticks, t)
			}
		}
	}()
	time.Sleep(5100 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker Stopped")
	deltas := make([]int64, len(ticks)-1)
	duration_micro := duration.Microseconds()
	for index, tick := range ticks[1:] {
		diff := tick.Sub(ticks[index]).Microseconds()
		delta := (duration_micro - diff)
		deltas[index] = delta
	}
	fmt.Println(deltas)
}
