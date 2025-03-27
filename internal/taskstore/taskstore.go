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
		return Task{}, fmt.Errorf("task with id: %d not found", id)
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

func (ts *TaskStore) DeleteAllTasks() error {
	ts.Lock()
	defer ts.Unlock()

	clear(ts.tasks)

	return nil
}

func (ts *TaskStore) GetAllTasks() []Task {
	ts.Lock()
	defer ts.Unlock()

	var tasks = make([]Task, 0, len(ts.tasks))
	for _, task := range ts.tasks {
		tasks = append(tasks, task)
	}

	return tasks
}

func (ts *TaskStore) GetTasksByTag(tag string) []Task {
	ts.Lock()
	defer ts.Unlock()

	var tasks []Task

taskloop:
	for _, task := range ts.tasks {
		for _, taskTag := range task.Tags {
			if taskTag == tag {
				tasks = append(tasks, task)
				continue taskloop // early iteration for efficiency
			}
		}
	}

	return tasks
}

func (ts *TaskStore) GetTasksByDueDate(year int, month time.Month, day int) []Task {
	ts.Lock()
	defer ts.Unlock()

	var tasks []Task

	for _, task := range ts.tasks {
		y, m, d := task.Due.Date()
		if y == year && m == month && d == day {
			tasks = append(tasks, task)
		}
	}

	return tasks
}
