package main

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"
)

var NumCPU int
var CountTo float64

func main() {
	NumCPU = runtime.NumCPU()

	runtime.GOMAXPROCS(NumCPU)
	http.HandleFunc("/", billion)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func billion(res http.ResponseWriter, req *http.Request) {

	startingTime := time.Now().UTC()
	fmt.Println(startingTime)
	CountTo = 1000000000
	CountToPer := math.Floor(CountTo / float64(NumCPU))

	var wg sync.WaitGroup
	for i := 0; i < NumCPU-1; i++ {
		wg.Add(1)
		go func() {
			for i := float64(0); i < CountToPer; i++ {
			}
			wg.Done()
		}()
		fmt.Println("started 1")
	}

	wg.Wait()

	endingTime := time.Now().UTC()

	var duration time.Duration = endingTime.Sub(startingTime)

	fmt.Fprintln(res, fmt.Sprintf(`<html><head><title>Go counts to %.0f</title></head><body><h1>Go counts to 1 billion</h1><p>This is how long it took me to count to 1billion, with %v CPUs:</p><p>%vms</p></body></html>`, CountTo, NumCPU, duration.Nanoseconds()/1e6))

}
