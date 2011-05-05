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

	g_cpu_die_sensor = "temp5_input"
	g_fan_sensor     = "fan1_input"
	g_fan_max        = "fan1_max"
	g_fan_out        = "fan1_min" //write to ths to set the fan RPM
)

const (
	g_min_temp      = 40.0 //what we wish our temp would always be in a world full of rainbow shitting unicorns
	g_min_fan_speed = 2000.0
)

var g_max_fan_speed float64 = 6200.0
