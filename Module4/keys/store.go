package keys

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// section: store
type Store struct {
	id     int
	values map[string]interface{}
	mu     sync.RWMutex
}

func NewStore() *Store {
	s := Store{}
	s.values = make(map[string]interface{})
	return &s
}

// section: store

// section: methods
func (vs *Store) Set(key string, value interface{}) {
	vs.id++
	// Make this an asynchronous call
	go func() {
		vs.mu.Lock()
		defer vs.mu.Unlock()

		// take a random amout of time to return to simulate the real world
		t := time.Duration(rand.Intn(3)) * time.Second
		time.Sleep(t)

		vs.values[key] = value
		log.Println("inserted: ", key, " with value of ", value)
	}()
}

func (vs *Store) Get(key string) (interface{}, error) {
	if v, ok := vs.values[key]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("%s not found", key)
}

func (vs *Store) Count() int {
	return len(vs.values)
}
