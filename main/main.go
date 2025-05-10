package main

import (
	// system packages
	"log"
	"net/http"
)

// bot configuration values read from file
var config Config

// initialize a session Discord servers
func init() {
  // get configuration value
  GetBotConfiguration(&config)
}

func main() {
  // connect to server & start listening

  // send HTTP request
  client := &http.Client{}

  req, err := http.NewRequest("GET", "https://kapi.kakao.com/v2/user/me", nil)
  if err !=  nil {
    log.Fatalf("wtf? [%v]", err)
  }
  req.Header.Add("Authorization", "rVTtgAFwAyZvF7CAHAhjn5Lm_2ZKUpJiAAAAAQo8IpsAAAGU7rnnqM6SpOckXrb0")
  resp, err:= client.Do(req)

  log.Printf("Printing results: [%v]", resp)

  // TODO: remove after listener is added
	//WaitOnInterrupt()

  // disconnect all services, then die
}
