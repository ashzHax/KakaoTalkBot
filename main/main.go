package main

import (
  // system packages
  "log"
)

// bot configuration values read from file
var config Config

// initialize a session Discord servers
func init() {
	var e error

  // get configuration value
  GetBotConfiguration(&config)

	// if error is found
	if e != nil {
		log.Fatalf("invalid bot parameters [%v]", e)
	}
}

func main() {
  // connect to server & start listening
	WaitOnInterrupt()
  // disconnect all services, then die
}
