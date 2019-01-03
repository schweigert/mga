package metric

import (
	"log"
	"os"
	"strconv"

	graphite "github.com/marpaia/graphite-golang"
)

var GRAPH_CLIENT *graphite.Graphite

func init() {
	port, _ := strconv.Atoi(os.Getenv("GRAPHITE_PORT"))
	graph, err := graphite.NewGraphite(os.Getenv("GRAPHITE_ADDR"), port)
	if err != nil {
		log.Println("Error on connect to GRAPHITE:", err)
		graph = graphite.NewGraphiteNop(os.Getenv("GRAPHITE_ADDR"), port)
	}

	GRAPH_CLIENT = graph
}
