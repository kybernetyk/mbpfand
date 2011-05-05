package main

import (
	"fmt"
	"time"
	"math"
)

var control_chan chan string = make(chan string)

func calc_fan_speed(temp float64) float64 {
	x := math.Log10(temp/g_min_temp)/math.Log10(2.0)*g_max_fan_speed
	//clipping to min/max will be done by SetFanSpeed()
	return x
}

func DoWork() {
	f := GetAverageTemp()
	fmt.Println("Average temperature:", f)
	speed := GetFanSpeed()
	fmt.Println("Fan Speed:", speed)

	rpm := calc_fan_speed(f)
	fmt.Println("Setting Fan Speed to:", int32(rpm))
	fmt.Println("")
	SetFanSpeed(rpm)
}

func seconds(n int64) int64 {
	return 1000000000 * n
}

func main() {
	g_max_fan_speed = readSensor(g_fan_max)
	fmt.Println("Max Fan Speed for this system:", g_max_fan_speed)

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
