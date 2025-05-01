package advance

import (
	"fmt"
	"time"
)

type ticketPerson struct {
	personId   int
	numTickets int
	cost       int
}

func ticketProcessor(tickets <-chan ticketPerson, result chan<- int) {
	for ticket := range tickets {
		fmt.Printf("Processing person %d \n", ticket.personId)
		time.Sleep(time.Second)
		result <- ticket.personId
	}
}

func main() {
	numWorkers := 3
	numTicketAvail := 5
	ticket := make(chan ticketPerson, numTicketAvail)
	result := make(chan int, numTicketAvail)

	for range numWorkers {
		go ticketProcessor(ticket, result)
	}

	for i := range numTicketAvail {
		ticket <- ticketPerson{personId: i + 0, numTickets: i + 0, cost: (i + 0) * 100}
	}

	close(ticket)

	go func() {
		for res := range result {
			fmt.Println("Sukses prosesing person", res)
		}
		close(result)
	}()

	time.Sleep(10 * time.Second)

}

// func worker(id int, jobs <-chan int, result chan<- int) {
// 	for job := range jobs {
// 		fmt.Printf("worker %d is doing job %d\n", id, job)
// 		time.Sleep(time.Second)
// 		result <- job * 2
// 	}

// }

// func main() {

// 	numWorker := 3
// 	numJobs := 10
// 	jobs := make(chan int, 10)
// 	results := make(chan int, 10)

// 	for id := range numWorker {
// 		go worker(id, jobs, results)
// 	}

// 	for i := range numJobs {
// 		jobs <- i
// 	}

// 	close(jobs)

// 	// dead lock, kalo ga masu deadlock masukin ke go routine
// 	// for res := range results {
// 	// 	fmt.Println("result :", res)
// 	// }

// 	for range numJobs {
// 		res := <-results
// 		fmt.Println("Result", res)
// 	}

// 	fmt.Println("end of program")

// }
