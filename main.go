package main

import (
	"fmt"
	"time"
)

var control_chan chan string = make(chan string)

func DoWork() {
	f := GetAverageTemp()
	fmt.Println("Average temperature:", f)
	speed := GetFanSpeed()
	fmt.Println("Fan Speed:", speed)

	rpm := 2000.0/40.0*f
	if rpm < 2000.0 {
		rpm = 2000.0
	}
	fmt.Println("Setting Fan Speed to:", int32(rpm))
	fmt.Println("")
	SetFanSpeed(rpm)
}

func seconds(n int64) int64 {
	return 1000000000 * n
}

func main() {
	ticker := time.NewTicker(seconds(10))
L:
	for {
		select {
		case msg := <-control_chan:
			if msg == "quit" {
				break L
			}
		case <-ticker.C:
			go DoWork()
		}
	}
}
