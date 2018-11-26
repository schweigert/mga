package metric

import (
	"fmt"
	"time"
)

func Timer(key string, f func()) {
	t := time.Now()
	f()

	duration := time.Since(t)
	GRAPH_CLIENT.SimpleSend(fmt.Sprintf("rudyc.%s", key), fmt.Sprintf("%v", duration.Seconds()))
}
