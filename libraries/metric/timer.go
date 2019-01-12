package metric

import (
	"fmt"
	"log"
	"time"
)

// Timer a function storing into the key
func Timer(key string, f func()) {
	t := time.Now()
	log.Println("Timer:", key)
	f()

	duration := time.Since(t)
	err := GRAPH_CLIENT.SimpleSend(fmt.Sprintf("go.%s", key), fmt.Sprintf("%v", duration.Seconds()))
	if err != nil {
		panic(err)
	}
}
