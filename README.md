# go-prometheus
A simple go application that makes use of the prometheus go client library. Two metrics are implemented in this app: a counter vector and a gauge. 

## Code structure
- `main.go` - App's entry point -- starts up metrics and HTTP server as go routines.
- `/Metrics`- Contains implementation of counter and gauge metrics, in addition to HTTP server to expose metrics
- `/Server` - Contains app's HTTP server and endpoints


## Starting the app
Run the following in project root directory to compile and run the main package 
- `go run main.go` <br />

This starts up app's HTTP server, making it ready to listen to client requests at http://localhost:8080. Next, HTTP server in metrics package initializes and starts exposing metrics at http://localhost:6060/metrics. 

## Starting prometheus
Run the following in project root directory to run prometheus in a docker container
- `docker-compose up -d`

You should see prometheus UI and scraped metrics at http://localhost:9090