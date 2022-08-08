package work_pools_test

import (
	"fmt"
	"testing"
	"time"
)

func TestWorkPool(t *testing.T) {
	jobs := make(chan int, 100)
	receivers := make(chan int, 100)
	// 先创建池
	for i := 0; i < 3; i++ {
		go worker(i, jobs, receivers)
	}
	// 写入
	for i := 0; i < 9; i++ {
		jobs <- i
	}
	close(jobs)
	for i := 0; i < 9; i++ {
		fmt.Println("result", <-receivers)
	}
	close(receivers)
}

func worker(id int, jobs chan int, receivers chan int) {
	for job := range jobs {
		fmt.Println("job id", id, "value", job)
		time.Sleep(time.Second)
		receivers <- job
	}
}
