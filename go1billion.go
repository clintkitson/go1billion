package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", billion)
	fmt.Println("listening...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func billion(res http.ResponseWriter, req *http.Request) {

	startingTime := time.Now().UTC()

	for i := 0; i < 1000000000; i++ {

	}

	endingTime := time.Now().UTC()

	var duration time.Duration = endingTime.Sub(startingTime)

	fmt.Fprintln(res, fmt.Sprintf("Native [%v]\nMilliseconds [%d]\nSeconds [%.3f]\n",
		duration,
		duration.Nanoseconds()/1e6,
		duration.Seconds()))

	//	fmt.Fprintln(res, fmt.Sprintf(`<!DOCTYPE html><html><head><title>GO counts to 1 billion</title><link rel="stylesheet" href="/stylesheets/style.css"></head><body><h1>Go counts to 1 billion</h1><p>This is how long it took me to count to 1billion:</p><p>xx</p></body></html>`))
}
