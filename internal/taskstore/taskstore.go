// package to provide a simple in-memory data storage for our tasks API
package taskstore

import (
	"sync"
	"time"
)

type Task struct {
	Id   int       `json:"id"`
	Text string    `json:"text"`
	Tags []string  `json:"tags"`
	Due  time.Time `json:"due"`
}

// in-memory of tasks; thread-safe
type TaskStore struct {
	// give all methods of the Mutex type to this type
	sync.Mutex

	// hashmap of id:task
	tasks  map[int]Task
	nextId int
}
