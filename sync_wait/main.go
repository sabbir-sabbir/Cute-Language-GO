package main

import (
	"fmt"
	"sync"
)


func worker(i int, wg *sync.WaitGroup){
	defer wg.Done() //decrement the routine after completion
	fmt.Println("worker work started", i)
	// some task are happening
	fmt.Println("worker work end", i)
	

}

func main() {
	fmt.Println("worker task main function")
     var wg sync.WaitGroup
	// three workers
	for i:=1; i <= 3; i++ {
		wg.Add(1)  //increment the ruitine
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("worker task completed")
}