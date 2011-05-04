package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"bytes"
)

func GetAverageTemp() float64 {
	temp1 := readSensor(g_cpu_die_sensor)
	avg := temp1 / 1000.0
	return avg
}

func GetFanSpeed() float64 {
	speed := readSensor(g_fan_sensor)
	return speed
}

func readSensor(sensorname string) float64 {
	s, err := ioutil.ReadFile(g_sensors_base_dir + sensorname)
	if err != nil {
		fmt.Println("Couldn't open sensor: ", g_sensors_base_dir+sensorname)
		return 0.0
	}
	s = bytes.Trim(s, "\r\n")

	f, err := strconv.Atof64(string(s))
	if err != nil {
		fmt.Println("Couldn't convert value to float64: ", string(s), s)
		return 0.0
	}
	return f
}
