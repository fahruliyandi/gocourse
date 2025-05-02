package advance

import (
	"fmt"
	"sync"
	"time"
)

// ========example using struct============
type Worker struct {
	ID   int
	task string
}

func (w Worker) working(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d is working on task %s\n", w.ID, w.task)
	time.Sleep(2 * time.Second)
	fmt.Printf("Worker %d finished on task %s\n", w.ID, w.task)
}

func main() {
	var wg sync.WaitGroup
	tasks := []string{"digging", "laying brick", "painting"}
	wg.Add(len(tasks))

	for i, task := range tasks {
		worker := Worker{
			ID:   i + 1,
			task: task,
		}
		go worker.working(&wg)
	}

	wg.Wait()

	fmt.Println("End of program")
}

// ==== example using channel=========
// func worker(id int, jobs <-chan int, result chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Workker %d starting..\n", id)
// 	for job := range jobs {
// 		time.Sleep(3 * time.Second) // simulate some work
// 		result <- job
// 	}
// 	fmt.Printf("Worker %d finished..\n", id)

// }

// func main() {
// 	numWorkers := 3
// 	numJobs := 5
// 	result := make(chan int, numJobs)
// 	jobs := make(chan int, numJobs)
// 	var wg sync.WaitGroup

// 	wg.Add(numWorkers)

// 	for i := range numWorkers {
// 		go worker(i+1, jobs, result, &wg)
// 	}

// 	for i := range numJobs {
// 		jobs <- i + 1
// 	}

// 	close(jobs)

// 	go func() {
// 		wg.Wait()
// 		close(result)
// 	}()

// 	// wg.Wait()
// 	// close(result)

// 	for res := range result {
// 		fmt.Println("result", res)
// 	}

// 	fmt.Println("End of program")

// }

// ==========basic example without using channel===============
// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	// wg.Add(1) wrong practice
// 	fmt.Printf("Worker %d starting\n", id)
// 	time.Sleep(time.Second) //simulate prosesing task
// 	fmt.Printf("Worker %d finished\n", id)
// }

// func main() {

// 	var wg sync.WaitGroup
// 	numWorkers := 3

// 	wg.Add(numWorkers) // this is the correc way to add wait group

// 	for i := range numWorkers {
// 		go worker(i, &wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("All workers finished")

// }
