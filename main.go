package main

import (
	"fmt"
	"time"
	"math"
)

var control_chan chan string = make(chan string)

func calc_fan_speed(temp float64) float64 {
	switch (g_opt_mode) {
	case mode_Default:
		return math.Log10(temp/40.0) / 0.3 * g_max_fan_speed  //quiet but not so cool
	case mode_Aggressive:
		return math.Log10(temp/30.0) / 0.35 * g_max_fan_speed //cooler but also louder
	}
	return g_min_fan_speed
}

//check temp, set speed
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

//turns seconds into nanoseconds ... for all the folks who hate zeros
func seconds(n int64) int64 {
	return 1000000000 * n
}

func main() {
	g_max_fan_speed = readSensor(g_fan_max)
	fmt.Println("Max Fan Speed for this system:", g_max_fan_speed)

	ticker := time.NewTicker(seconds(g_job_fire_time))
L:
	for {
		select {
		case msg := <-control_chan:
			if msg == "quit" {
				break L
			}
		case <-ticker.C:
			DoWork()
		}
	}
}
