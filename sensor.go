package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

//todo: read moar sensors (preferably 2 left / 2 right)
func GetAverageTemp() float64 {
	temp1 := readSensor(g_cpu_die_sensor)
	avg := temp1 / 1000.0
	return avg
}

func GetFanSpeed() float64 {
	speed := readSensor(g_fan_sensor)
	return speed
}

func SetFanSpeed(new_speed float64) {
	if new_speed > g_max_fan_speed {
		new_speed = g_max_fan_speed
	}
	if new_speed < g_min_fan_speed {
		new_speed = g_min_fan_speed
	}
	s := fmt.Sprintf("%d", int32(new_speed))
	b := []byte(s)

	err := ioutil.WriteFile(g_sensors_base_dir+g_fan_out, b, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't set fan speed: %s\n", err)
		return
	}
	verbOutp("Setting Fan Speed to:", int32(new_speed))
	verbOutp("")
}

func readSensor(sensorname string) float64 {
	s, err := ioutil.ReadFile(g_sensors_base_dir + sensorname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't open sensor: %s\n", g_sensors_base_dir+sensorname)
		return 0.0
	}
	s = bytes.Trim(s, "\r\n")

	f, err := strconv.ParseFloat(string(s), 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't convert value to float64: ", string(s), s)
		return 0.0
	}
	return f
}
