package main

import (
	"context"
	"fmt"
	"sync"
)

type Job struct {
	ID   int
	Name string
}

type Result struct {
	ID   int
	Name string
}

type WorkerPool struct {
	jobs     chan Job
	results  chan Result
	ctx      context.Context
	cancel   context.CancelFunc
	wg       sync.WaitGroup
}

func NewWorkerPool() *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	return &WorkerPool{
		jobs:     make(chan Job),         // Unbuffered
		results:  make(chan Result),      // Unbuffered
		ctx:      ctx,
		cancel:   cancel,
	}
}

func (wp *WorkerPool) StartWorkers(n int) {
	wp.wg.Add(n)
	for i := 0; i < n; i++ {
		go func(workerID int) {
			defer wp.wg.Done()
			wp.worker(workerID)
		}(i)
	}

	// Monitor workers and close results when done
	go func() {
		wp.wg.Wait()
		close(wp.results)
	}()
}

func (wp *WorkerPool) worker(id int) {
	for {
		select {
		case <-wp.ctx.Done():
			fmt.Printf("Worker %d shutting down\n", id)
			return
		// Worker only receives job when producer is ready to send
		case job, ok := <-wp.jobs:
			if !ok {
				fmt.Printf("Worker %d: jobs channel closed\n", id)
				return
			}
			// Process job
			fmt.Printf("Worker %d processing job %d\n", id, job.ID)
			
			select {
			// if the job is cancelled, the worker will stop processing
			case <-wp.ctx.Done():
				return
			case wp.results <- Result{ID: job.ID, Name: job.Name}:
				fmt.Printf("Worker %d completed job %d\n", id, job.ID)
			}
		}
	}
}

// Producer interface
func (wp *WorkerPool) SubmitJob(job Job) error {
	select {
	case <-wp.ctx.Done():
		return fmt.Errorf("worker pool is shutting down")
	// Producer automatically slows down if workers are busy
	case wp.jobs <- job: // Will block if no worker is ready
		return nil
	}
}

// Consumer interface
func (wp *WorkerPool) Results() <-chan Result {
	return wp.results
}

func (wp *WorkerPool) Shutdown() {
	wp.cancel()  // Signal shutdown
	close(wp.jobs)
	wp.wg.Wait()
}

func main() {
	pool := NewWorkerPool()
	
	// Start workers
	numWorkers := 3
	pool.StartWorkers(numWorkers)

	// Multiple producers
	var producerWg sync.WaitGroup
	numProducers := 2
	producerWg.Add(numProducers)

	for i := 0; i < numProducers; i++ {
		go func(producerID int) {
			defer producerWg.Done()
			
			// Producer will naturally slow down if workers are busy
			for j := 0; j < 5; j++ {
				job := Job{
					ID:   producerID*100 + j,
					Name: fmt.Sprintf("Job-%d-%d", producerID, j),
				}
				
				select {
				case <-pool.ctx.Done():
					return
				default:
					if err := pool.SubmitJob(job); err != nil {
						fmt.Printf("Failed to submit job: %v\n", err)
						return
					}
				}
			}
		}(i)
	}

	// Multiple consumers
	var consumerWg sync.WaitGroup
	numConsumers := 2
	consumerWg.Add(numConsumers)

	for i := 0; i < numConsumers; i++ {
		go func(consumerID int) {
			defer consumerWg.Done()
			
			for result := range pool.Results() {
				fmt.Printf("Consumer %d received result for job %d\n", 
					consumerID, result.ID)
			}
		}(i)
	}

	// Wait for producers to finish
	producerWg.Wait()
	
	// Shutdown pool
	pool.Shutdown()

	// Wait for consumers to finish
	consumerWg.Wait()
}
