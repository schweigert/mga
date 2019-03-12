package metric

import (
	"log"
	"os"
	"strconv"

	graphite "github.com/marpaia/graphite-golang"
)

// GraphClient store a sample graphite cliente connection
var GraphClient *graphite.Graphite

func init() {
	port, _ := strconv.Atoi(os.Getenv("GRAPHITE_PORT"))
	graph, err := graphite.NewGraphite(os.Getenv("GRAPHITE_ADDR"), port)
	if err != nil {
		log.Println("Error on connect to GRAPHITE:", err)
		graph = graphite.NewGraphiteNop(os.Getenv("GRAPHITE_ADDR"), port)
	}

	GraphClient = graph
}
