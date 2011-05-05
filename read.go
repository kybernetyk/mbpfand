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

func SetFanSpeed(new_speed float64) {
	s := fmt.Sprintf("%d", int64(new_speed))	
	b := []byte(s)

	err := ioutil.WriteFile(g_sensors_base_dir + g_fan_out, b, 0644) 
	if err != nil {
		fmt.Println("Couldn't set fan speed:", err.String())
	}
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
