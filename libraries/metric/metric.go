package metric

import (
	"fmt"
	"strconv"
)

// Int send a function into a key
func Int(key string, value int) {
	err := GraphClient.SimpleSend(fmt.Sprintf("go.%s", key), strconv.Itoa(value))
	if err != nil {
		panic(err)
	}
}
