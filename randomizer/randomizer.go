package randomizer

import (
	"math/rand"
	"sync"
)

var mutex *sync.Mutex

func init() {
	mutex = &sync.Mutex{}
}

func Int(max int) int {
	mutex.Lock()
	defer mutex.Unlock()

	if max <= 0 {
		return 0
	}

	return rand.Int() % max
}
