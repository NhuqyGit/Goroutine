package main

import (
	"fmt"
	"sync"
	"time"
)

func sendNotification(wg *sync.WaitGroup, sendAt time.Time){
	duration := time.Until(sendAt)
	
	fmt.Println("WAITING........")
	time.Sleep(duration)
	fmt.Println("HOPE YOU HAVE A NICE DAY!!!")
	wg.Done()
}

func printHello(){
	for {
		fmt.Println("HELLO")
		time.Sleep(1*time.Second)
	}
}

func main(){
	var wg sync.WaitGroup

	// Add 1 goroutine to wait
	wg.Add(1)

	sendAt := time.Now().Add(5 * time.Second)
	go sendNotification(&wg, sendAt)
	
	go printHello()

	// Main is doing another job
	for i:=1; i<5; i++{
		fmt.Println(i)
	}
	
	//Waiting for the sendNotification goroutine done and return
	wg.Wait()
	fmt.Println("DONEEEEEE")
}