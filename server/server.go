package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"myapp/metrics"
	"time"
)

func Start() {
	router := gin.Default()

	//Defining endpoints
	router.GET("/first_service", firstService)
	router.GET("/second_service", secondService)

	log.Println("Starting HTTP server at port 8080...")

	//Listen and serve HTTP requests on port :8080 by default
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to make connection....")

	}

}
func firstService(c *gin.Context) {
	//increase counter
	log.Println("First Service called")
	metrics.FirstServiceCount.Inc()

	//mimicking task duration
	metrics.TasksQueued.Inc()
	time.Sleep(5 * time.Second)
	metrics.TasksQueued.Dec()
}

func secondService(c *gin.Context) {
	//increase counter
	log.Println("Second Service called")
	metrics.SecondServiceCount.Inc()

	//mimicking task duration
	metrics.TasksQueued.Inc()
	time.Sleep(3 * time.Second)
	metrics.TasksQueued.Dec()
}
