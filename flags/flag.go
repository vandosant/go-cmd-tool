package main

import (
"flag"
"fmt"
"os"
)

func main() {
	var (
		cmd  string = "website"
		port int    = 8000
		log int = 1
	)

	fs := flag.NewFlagSet("default", flag.ExitOnError)

	fs.StringVar(&cmd, "cmd", cmd, "the command to run")
	fs.IntVar(&port, "p", port, "the port to run on")
	fs.IntVar(&log, "l", log, "the log level")

	fs.Parse(os.Args[1:])

	fmt.Printf("Running %q on port %d with log level %d", cmd, port, log)
}
