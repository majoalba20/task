package queue

import (
	"go-repaso/internal/domain"
	"go-repaso/internal/repository"
	"log"
	"time"
)

type TaskWorker struct {
	repo  repository.TaskRepository
	queue TaskQueue
}

func NewTaskWorker(repo repository.TaskRepository, queue TaskQueue) *TaskWorker {
	return &TaskWorker{
		repo:  repo,
		queue: queue,
	}
}

func (w *TaskWorker) Start() {
	go func() {
		for taskID := range w.queue.Dequeue() {
			log.Printf("processing task %d\n", taskID)
			task, err := w.repo.FindByID(taskID)
			if err != nil {
				log.Printf("task %d not found: %v\n", taskID, err)
				continue
			}
			task.Status = domain.StatusProcessing
			if err := w.repo.Update(task); err != nil {
				log.Printf("failed to set processing for task %d: %v\n", taskID, err)
				continue
			}
			time.Sleep(5 * time.Second)
			// Simulación
			if task.Title == "fail" {
				msg := "simulated processing failure"
				task.Status = domain.StatusFailed
				task.ErrorMessage = &msg
				task.Result = nil
			} else {
				result := "task processed successfully"
				task.Status = domain.StatusDone
				task.Result = &result
				task.ErrorMessage = nil
			}
			if err := w.repo.Update(task); err != nil {
				log.Printf("failed to finalize task %d: %v\n", taskID, err)
				continue
			}
			log.Printf("task %d finished with status %s\n", taskID, task.Status)
		}
	}()
}
