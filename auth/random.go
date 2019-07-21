package auth

import (
	"math/rand"
	"sync"
	"time"
)

// Adapted from this code: https://nishanths.svbtle.com/do-not-seed-the-global-random
type Erandom struct {
	Random *rand.Rand
}

// Create a new reference to the random number generator
func NewRandom() (*Erandom, error) {

	var ra = rand.New(
		&lockedRandSource{
			src: rand.NewSource(time.Now().UnixNano()),
		},
	)

	r := &Erandom{
		Random: ra,
	}

	return r, nil
}

// locked to prevent concurrent use of the underlying source
type lockedRandSource struct {
	lock sync.Mutex // protects src
	src  rand.Source
}

// to satisfy rand.Source interface
func (r *lockedRandSource) Int63() int64 {
	r.lock.Lock()
	ret := r.src.Int63()
	r.lock.Unlock()
	return ret
}

// to satisfy rand.Source interface
func (r *lockedRandSource) Seed(seed int64) {
	r.lock.Lock()
	r.src.Seed(seed)
	r.lock.Unlock()
}
