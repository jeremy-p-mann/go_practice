package main

import (
	"log"
	"time"
)

func workerx(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		log.Println("worker ", id, "started job", j)
		time.Sleep(time.Second)
		log.Println("worker ", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	for w := 1; w <= 3; w++ {
		go workerx(w, jobs, results)
	}
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a < numJobs; a++ {
		log.Println(<-results)
	}
}
