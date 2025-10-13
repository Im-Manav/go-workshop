package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
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

	primeOrdinal := 1
	var timeToFirst1000KPrimes time.Duration

	var i int
	for {
		if isPrime(i) {
			fmt.Printf("Prime %d found: %d\n", primeOrdinal, i)
			primeOrdinal++
		}
		if primeOrdinal == 1000_000 {
			timeToFirst1000KPrimes = time.Since(start)
		}
		if ctx.Err() != nil {
			break
		}
		i++
	}

	if primeOrdinal <= 1000_000 {
		fmt.Printf("Total primes found: %d\n", primeOrdinal-1)
	} else {
		fmt.Printf("Time to find first 1000,000 primes: %v\n", timeToFirst1000KPrimes)
	}
}
