package main

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"runtime"
	"strconv"
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

	requestNumCPU, _ := strconv.Atoi(req.URL.Query().Get("numcpu"))

	if requestNumCPU > NumCPU || requestNumCPU < 1 {
		requestNumCPU = 1
	}

	CountTo = 1000000000
	CountToPer := int(math.Floor(CountTo / float64(requestNumCPU)))
	fmt.Println(CountToPer)
	var wg sync.WaitGroup
	startingTime := time.Now().UTC()
	for i := 0; i < requestNumCPU; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(fmt.Sprintf("started %v", CountToPer))
			for i := 0; i < CountToPer; i++ {
			}
			wg.Done()
		}()

	}

	wg.Wait()
	endingTime := time.Now().UTC()

	var duration time.Duration = endingTime.Sub(startingTime)

	fmt.Fprintln(res, fmt.Sprintf(`<html><head><title>Go counts to %.0f</title></head><body><h1>Go counts to 1 billion</h1><p>This is how long it took me to count to 1billion, with %v CPUs:</p><p>%vms</p></body></html>`, CountTo, requestNumCPU, duration.Nanoseconds()/1e6))

}
