package main

// Generate sieve
func Generate(sieve chan<- int) {
	for prime := 2; prime < 100; prime++ {
		sieve <- prime // send numbers to sieve
	}
	close(sieve)
}

// Filter - remove items divisible by prime from sieve
func Filter(in <-chan int, out chan<- int, prime int) {
	for number := range in {
		// if input not divisible by prime, send to output
		if number%prime != 0 {
			out <- number
		}
	}
}

func main() {
	sieve := make(chan int)
	go Generate(sieve)
	for prime := range sieve {
		println(prime)
		ch := make(chan int)
		go Filter(sieve, ch, prime)
		sieve = ch
	}
}
