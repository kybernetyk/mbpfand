package main

import (
	"fmt"
	"time"
	"os"
	"os/signal"
	"syscall"
	"flag"
)

var control_chan = make(chan string)

var (
	mode         = flag.String("m", "default", "The mode to use. default or aggressive")
	show_version = flag.Bool("V", false, "Print mbpfand's version and exit.")
	be_verbose   = flag.Bool("v", false, "Verbose output.")
	update_rate  = flag.Int64("u", 10, "Update rate in seconds.")
	set_speed    = flag.Bool("S", false, "Set fan speed and exit. Speed specified by -s.")
	print_speed  = flag.Bool("Pf", false, "Print current fan speed and exit.")
	print_temp   = flag.Bool("Pt", false, "Print current temp and exit.")
	force_speed  = flag.Int64("s", 2000, "Fan speed for -S option.")
	show_usage   = flag.Bool("h", false, "Show usage options and exit.")
)


//turns seconds into nanoseconds ... for all the folks who hate zeros
func seconds(n int64) int64 {
	return 1000000000 * n
}

func verbOutp(v ...interface{}) {
	if *be_verbose {
		fmt.Fprintf(os.Stdout, fmt.Sprintln(v...))
	}
}

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\nOptions:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = Usage
	flag.Parse()

	if *show_usage {
		flag.Usage()
		return
	}

	if *show_version {
		fmt.Println("mbpfand", g_mbpfand_version)
		fmt.Println("\t(c) Leon Szpilewski, 2011")
		fmt.Println("\tLicensed under GPL v3")
		return
	}

	switch *mode {
	case "aggressive":
		g_opt_mode = mode_Aggressive
	default:
		g_opt_mode = mode_Default
	}
	modes := map[ModeType]string{
		mode_Aggressive: "[aggressive]",
		mode_Default:    "[default]",
	}

	g_max_fan_speed = readSensor(g_fan_max)

	if *set_speed {
		SetFanSpeed(float64(*force_speed))
		return
	}

	if *print_speed {
		speed := GetFanSpeed()
		fmt.Println("Current fan speed:", speed)
		if !*print_temp {
			return
		}
	}

	if *print_temp {
		temp := GetAverageTemp()
		fmt.Println("Current Avg Temp:", temp)
		return
	}

	defer SetFanSpeed(g_min_fan_speed)

	verbOutp("Using Mode:", modes[g_opt_mode])
	verbOutp("Max Fan Speed for this system:", g_max_fan_speed)
	verbOutp("Update Rate is:", *update_rate)
	ticker := time.NewTicker(seconds(*update_rate))
L:
	for {
		select {
		case msg := <-control_chan:
			if msg == "quit" {
				break L
			}
		case <-ticker.C:
			DoWork()

		case sig := <-signal.Incoming:
			fmt.Println("Got signal: " + sig.String())
			switch sig.(signal.UnixSignal) {
			case syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGHUP:
				return
			}
		}
	}
}
