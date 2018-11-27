package randomizer

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var mutex *sync.Mutex

func init() {
	log.Println("Starting randomizer...")

	// Init mutex
	mutex = &sync.Mutex{}
}

func Int(max int) int {
	// Init Seed
	seed := time.Now().UnixNano()
	rand.Seed(seed)

	mutex.Lock()
	defer mutex.Unlock()

	if max <= 0 {
		return 0
	}

	return rand.Int() % max
}
