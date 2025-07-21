package data

import (
	"errors"
	"sync"
	"task_manager/models"
)

type TaskService struct {
	tasks  map[int]models.Task
	mu     sync.Mutex
	nextID int
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks:  make(map[int]models.Task),
		nextID: 1,
	}
}

func (s *TaskService) GetAllTasks() []models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	tasks := []models.Task{}
	for _, t := range s.tasks {
		tasks = append(tasks, t)
	}
	return tasks
}

func (s *TaskService) GetTask(id int) (models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	task, ok := s.tasks[id]
	if !ok {
		return models.Task{}, errors.New("task not found")
	}
	return task, nil
}

func (s *TaskService) CreateTask(task models.Task) models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	task.ID = s.nextID
	s.nextID++
	s.tasks[task.ID] = task
	return task
}

func (s *TaskService) UpdateTask(id int, updated models.Task) (models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.tasks[id]
	if !ok {
		return models.Task{}, errors.New("task not found")
	}
	updated.ID = id
	s.tasks[id] = updated
	return updated, nil
}

func (s *TaskService) DeleteTask(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.tasks[id]
	if !ok {
		return errors.New("task not found")
	}
	delete(s.tasks, id)
	return nil
}
