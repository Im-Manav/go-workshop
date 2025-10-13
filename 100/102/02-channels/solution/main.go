package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	start := time.Now()

	ctx, cancel := context.WithCancel(context.Background())

	// Handle interrupts to allow graceful shutdown. If you press Ctrl+C,
	// the program will stop accepting new jobs and finish processing before
	// exiting.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigs
		cancel()
	}()

	inputs := make(chan int, 100)
	primes := make(chan int)

	// Create some workers.
	var wg sync.WaitGroup
	// If there's only one worker, on my machine...
	// Time to find first 1000,000 primes:
	//  1 Worker:       1000.331149704ms
	//  NumCPU workers:  993.285018ms
	for range runtime.NumCPU() {
		wg.Go(func() {
			for job := range inputs {
				if isPrime(job) {
					primes <- job
				}
			}
		})
	}

	// Push jobs into the input queue for the workers.
	go func() {
		// When this function exits, wait for all workers to finish and then close
		// the primes channel to signal that no more primes will be sent.
		defer func() {
			wg.Wait()
			close(primes)
		}()

		var n int
		for {
			select {
			case inputs <- n:
				n++
				if n%100000 == 0 {
					fmt.Printf("Jobs submitted: %d\n", n)
				}
			case <-ctx.Done():
				fmt.Println("Shutting down job submission")
				close(inputs)
				return
			}
		}
	}()

	// Read from the primes channel until it's closed.
	primeOrdinal := 1
	var timeToFirst1000KPrimes time.Duration
	for p := range primes {
		fmt.Printf("Prime %d found: %d\n", primeOrdinal, p)
		primeOrdinal++

		if primeOrdinal == 1000_000 {
			timeToFirst1000KPrimes = time.Since(start)
		}
	}

	if primeOrdinal <= 1000_000 {
		fmt.Printf("Total primes found: %d\n", primeOrdinal-1)
	} else {
		fmt.Printf("Time to find first 1000,000 primes: %v\n", timeToFirst1000KPrimes)
	}
}
