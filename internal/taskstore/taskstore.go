// package to provide a simple in-memory data storage for our tasks API
package taskstore

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	Id   int       `json:"id"`
	Text string    `json:"text"`
	Tags []string  `json:"tags"`
	Due  time.Time `json:"due"`
}

// in-memory storage of tasks; thread-safe
type TaskStore struct {
	// give all methods of the Mutex type to this type
	sync.Mutex

	// hashmap id:task
	tasks map[int]Task

	// ID to use for next task inserted
	nextId int
}

func NewTaskStore() *TaskStore {
	ts := &TaskStore{}
	ts.tasks = make(map[int]Task)
	ts.nextId = 0
	return ts
}

// get tasks in a TaskStore by id
func (ts *TaskStore) GetTask(id int) (Task, error) {
	// lock ts and release it at the end of func
	ts.Lock()
	defer ts.Unlock()

	t, exists := ts.tasks[id]
	if exists {
		return t, nil
	} else {
		return Task{}, fmt.Errorf("task with id: %d not found!", id)
	}
}

// attempt to delete a task, err if doesn't exist
func (ts *TaskStore) DeleteTask(id int) error {
	ts.Lock()
	defer ts.Unlock()

	// err if task doesn't exist
	if _, ok := ts.tasks[id]; !ok {
		return fmt.Errorf("no task of ID: %d exists, cannot be deleted", id)
	}

	// delete task otherwise
	delete(ts.tasks, id)
	return nil
}
