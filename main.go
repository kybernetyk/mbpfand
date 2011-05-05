package main

import (
	"fmt"
	"time"
	"math"
)

var control_chan chan string = make(chan string)

//stupid
func calc_fan_speed(temp float64) float64 {
	rpm := 6200.0/90.0*temp
	return rpm
}

func calc_fan_speed_log(temp float64) float64 {
	x := math.Log10(temp/40.0)/math.Log10(2.0)*6200.00
	//clipping to min/max will be done by SetFanSpeed()
	return x
}

func DoWork() {
	f := GetAverageTemp()
	fmt.Println("Average temperature:", f)
	speed := GetFanSpeed()
	fmt.Println("Fan Speed:", speed)

	rpm := calc_fan_speed_log(f)
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
