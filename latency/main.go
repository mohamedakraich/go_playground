package main

import (
	"fmt"
	"time"
)

func measureTime(funcName string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("Time taken by %s function is %v \n", funcName, time.Since(start))
	}
}

func expensivePrint() {
	defer measureTime("expensivePrint")()
	for i := 1; i <= 5; i++ {
		fmt.Printf("Current number is %d \n", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	expensivePrint()
	fmt.Println("Finished executing main")
}

/*
func expensivePrint() {
	start := time.Now()
	for i := 1; i <= 5; i++ {
		fmt.Printf("Current number is %d \n", i)
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Printf("Time taken by expensivePrint function is %v \n", time.Since(start))
}
*/

/*
func printTimeMeasurement(funcName string, start time.Time) func() {
	return func() {
		fmt.Printf("Time taken by %s function is %v \n", funcName, time.Since(start))
	}
}
*/
