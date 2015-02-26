package main

import (
	"encoding/json"
	"fmt"
	"log"
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

type testInfo struct {
	InstanceNumCPU int
	RequestNumCPU  int
	CountTo        float64
	CountToPer     int
	StartingTime   time.Time
	EndingTime     time.Time
	Duration       time.Duration
	Durationms     int64
}

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
	var testInfo testInfo
	testInfo.InstanceNumCPU = NumCPU
	testInfo.RequestNumCPU, _ = strconv.Atoi(req.URL.Query().Get("numcpu"))

	if testInfo.RequestNumCPU > testInfo.InstanceNumCPU || testInfo.RequestNumCPU < 1 {
		testInfo.RequestNumCPU = 1
	}

	testInfo.CountTo = 1000000000
	testInfo.CountToPer = int(math.Floor(testInfo.CountTo / float64(testInfo.RequestNumCPU)))
	var wg sync.WaitGroup
	testInfo.StartingTime = time.Now().UTC()
	for i := 0; i < testInfo.RequestNumCPU; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(fmt.Sprintf("started %v", testInfo.CountToPer))
			for i := 0; i < testInfo.CountToPer; i++ {
			}
			wg.Done()
		}()

	}

	wg.Wait()
	EndingTime := time.Now().UTC()

	testInfo.Duration = EndingTime.Sub(testInfo.StartingTime)
	testInfo.Durationms = testInfo.Duration.Nanoseconds() / 1e6

	//fmt.Fprintln(res, fmt.Sprintf(`<html><head><title>Go counts to %.0f</title></head><body><pre>`, testInfo.CountTo))
	jsonOutput, err := json.Marshal(&testInfo)
	if err != nil {
		log.Fatalf("error marshaling: %s", err)
	}

	fmt.Fprintln(res, string(jsonOutput))
	//fmt.Fprintln(res, fmt.Sprintf(`</pre></body></html>`))

}
