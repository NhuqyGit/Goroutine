package main

import (
	"fmt"
	"sync"
	"time"
)

func printNum(wg *sync.WaitGroup, chanPrintNum chan int){
	defer wg.Done()
	sum := 0
	for i := 1; i <= 5; i++{
		sum += i
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	chanPrintNum <- sum
}

func printProcessedNum(wg *sync.WaitGroup, chanPrintNum chan int, chanProcessedNum chan int) {
	defer wg.Done()
	sum := <-chanPrintNum
	fmt.Println("bbbb: ", sum)
	for i := 1; i <= 5; i++{
		sum += i
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	chanProcessedNum <- sum
}

func printNum2(wg *sync.WaitGroup, chanPrintNum chan int){
	defer wg.Done()
	for i := 1; i <= 5; i++{
		times2 := i * 2
		fmt.Printf("print(%d): %d\n", i, times2)
		chanPrintNum <- times2
		// time.Sleep(1 * time.Second)
	}

}

func timeNum(wg *sync.WaitGroup, chanPrintNum chan int, chanProcessedNum chan int) {
	defer wg.Done()
	sum := 0
	for i := 1; i <= 5; i++{
		sum += <-chanPrintNum
		fmt.Printf("time(%d): %d\n", i, sum)
		// time.Sleep(1 * time.Second)
	}

	chanProcessedNum <- sum
}

func main(){
	chanPrintNum := make(chan int)
	chanProcessedNum := make(chan int)

	var wg sync.WaitGroup

	// Add 1 goroutine to wait
	wg.Add(2)

	go printNum2(&wg, chanPrintNum)
	go timeNum(&wg, chanPrintNum, chanProcessedNum)
	
	fmt.Println("processedTimeNum: ", <-chanProcessedNum)
	
	//Waiting for the sendNotification goroutine done and return
	wg.Wait()
	fmt.Println("DONEEEEEE")
}