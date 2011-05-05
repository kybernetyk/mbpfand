package main
/*
 Note: I've got a 13" MBP 5,5 - so these settings are pretty
 specific to this model. It has only one Fan - whereas the
 bigger models have more fans (I heard). So if you've got
 a 15"/17" then you'll have to submit a patch to make the
 demon use your 2nd fan.
*/
const (
	g_sensors_base_dir = "/sys/devices/platform/applesmc.768/"

	//temperature sensors
	g_cpu_die_sensor = "temp5_input"

	//fan control
	g_fan_sensor     = "fan1_input"
	g_fan_max        = "fan1_max"
	g_fan_out        = "fan1_min" //write to ths to set the fan RPM
)

const (
	g_min_fan_speed = 2000.0
)
var g_max_fan_speed float64 = 6200.0

const (
	g_job_fire_time = 10.0	//how often the DoWork() function shall be called. time in seconds
)

type ModeType int32

const (
	mode_Default = iota
	mode_Aggressive
)

var g_opt_mode ModeType = mode_Default
