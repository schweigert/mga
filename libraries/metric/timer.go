package metric

import (
	"fmt"
	"time"
)

func Timer(key string, f func()) {
	t := time.Now()
	f()

	duration := time.Since(t)
	err := GRAPH_CLIENT.SimpleSend(fmt.Sprintf("rudyc.%s", key), fmt.Sprintf("%v", duration.Seconds()))
	if err != nil {
		panic(err)
	}
}
