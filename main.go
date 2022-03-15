package main

import (
	"myapp/metrics"
	"myapp/server"
)

func main() {
	//Initializing child go routines
	go metrics.Start()
	go server.Start()

	//Waiting for a bool val from channel, which never happens
	//Prevents main go routine from exiting before child routines are complete
	<-make(chan bool)

}


