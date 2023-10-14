package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // 模拟工作
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2
	}
}

func main() {
	numWorkers := 2
	numJobs := 9

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 启动多个协程来执行工作
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// 向通道发送工作
	for j := 1; j <= numJobs; j++ {
		time.Sleep(1 * time.Second)
		jobs <- j
	}

	close(jobs) // 关闭工作通道，告知协程不再有工作

	// 收集结果
	for r := 1; r <= numJobs; r++ {
		result := <-results
		fmt.Printf("Received result: %d\n", result)
	}
}
