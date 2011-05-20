package main

import (
	"fmt"
	"math"
)

func Generate() chan uint {
	channel := make(chan uint)
	go func(){
		for i := uint(2); ; i++ {
			channel <- i
		}
	}()
	return channel
}

func FilterPrimesOnly(inChannel chan uint, prime uint) chan uint {
	outChannel := make(chan uint)
	go func(){
		for {
			i := <-inChannel
			if i % prime != 0 {
				outChannel <- i
			}
		}
	}()
	return outChannel
}

func Sieve(number float64) chan uint {
	channel := make(chan uint)

	go func() {
		ch := Generate()
		for {
			prime := <-ch
			channel <- prime
			ch = FilterPrimesOnly(ch, prime)
		}
	}()

	square := uint(math.Sqrt(number))
	primesToCheck := make(chan uint)
	go func(){
		for i := uint(1); i <= square; i++ {
			prime := <-channel
			if prime > square {
				break
			}
			primesToCheck <- prime
		}
	}()

	return channel
}

func Factor(i float64) uint {
	Sieve(i)
	return uint(i)
}

func main() {
	fmt.Println(Factor(600851475143))
}
