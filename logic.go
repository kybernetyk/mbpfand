package main

import (
	"math"
)

func calc_fan_speed(temp float64) float64 {
	switch g_opt_mode {
	case mode_Default:
		return math.Log10(temp/40.0) / 0.3 * g_max_fan_speed //quiet but not so cool
	case mode_Aggressive:
		return math.Log10(temp/30.0) / 0.35 * g_max_fan_speed //cooler but also louder
	}
	return g_min_fan_speed
}

//check temp, set speed
func DoWork() {
	f := GetAverageTemp()
	verbOutp("Average temperature:", f)
	speed := GetFanSpeed()
	verbOutp("Fan Speed:", speed)

	rpm := calc_fan_speed(f)
	SetFanSpeed(rpm)
}
