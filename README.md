# Sieve of Eratosthenes

## Why?

I wrote this program mainly to test the limit of the number of goroutines that can be running concurrently.

My implementation is such that one goroutine is spawned for one prime, so the number of primes found is roughly equal to the number of goroutines that were run concurrently.

## Usage

`clone` or `go get` this repo, then run

	go run soe.go -max=100000

The flag `max` specifies the largest number to be tested.

A file named `output.txt` will be generated.  It will contain all primes found.  The number of goroutines spawned in total will be printed to the console.