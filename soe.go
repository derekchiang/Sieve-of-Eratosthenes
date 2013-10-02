package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const STOP_SIGNAL = -1

var (
	numPrimes = 0
	max       = flag.Int("max", 1000, "The maximum number to be tested.")
)

func Consumer(inputChan chan int, output io.Writer, stopChan chan bool) {
	for {
		n := <-inputChan
		if n == STOP_SIGNAL {
			break
		}
		output.Write([]byte(fmt.Sprintf("%d  ", n)))
		numPrimes++

		// Create a new filter
		outputChan := make(chan int)
		go Filter(n, inputChan, outputChan)
		inputChan = outputChan
	}
	close(stopChan)
}

func Producer(outputChan chan int, max int) {
	for i := 2; i <= max; i++ {
		outputChan <- i
	}
	outputChan <- STOP_SIGNAL
}

func Filter(prime int, inputChan, outputChan chan int) {
	for {
		n := <-inputChan
		if n == STOP_SIGNAL {
			outputChan <- n
			return
		}
		if (n % prime) != 0 {
			outputChan <- n
		}
	}
}

func main() {
	flag.Parse()

	output, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := output.Close(); err != nil {
			panic(err)
		}
	}()

	c := make(chan int)
	stop := make(chan bool)
	go Producer(c, *max)
	go Consumer(c, output, stop)
	<-stop
	fmt.Printf("Number of primes found: %d\n", numPrimes)
	fmt.Printf("Number of goroutines generated: %d\n", numPrimes+2)
}
