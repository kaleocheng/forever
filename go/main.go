package main

import (
	"os"
	"os/signal"
	"time"
)

func main() {
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt)
		<-c
		os.Exit(0)
	}()
	for {
		time.Sleep(time.Hour)
	}
}
