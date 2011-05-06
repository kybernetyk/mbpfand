package main

import (
	"fmt"
	"time"
	"os"
	"flag"
)

var control_chan = make(chan string)

var (
	mode         = flag.String("m", "default", "The mode to use. default or aggressive")
	show_version = flag.Bool("V", false, "Print mbpfand's version and exit.")
	be_verbose   = flag.Bool("v", false, "Verbose output.")
	update_rate  = flag.Int64("u", 10, "Update rate in seconds.")
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

	if *show_version {
		fmt.Println("mbpfand", g_mbpfand_version)
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
	verbOutp("Using Mode:", modes[g_opt_mode])

	g_max_fan_speed = readSensor(g_fan_max)
	verbOutp("Max Fan Speed for this system:", g_max_fan_speed)

	verbOutp("Update Rate is:", *update_rate)
	ticker := time.NewTicker(seconds(*update_rate))
	verbOutp("Scheduled ...\n")
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
