package main

import (
	"fmt"
	"math"
)

func Generate(number uint) chan uint {
	channel := make(chan uint)
	go func(){
		for i := uint(2); i <= number; i++ {
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
	square := uint(math.Sqrt(number))

	channel := make(chan uint)
	go func() {
		ch := Generate(uint(number))
		for {
			prime := <-ch
			channel <- prime
			if  prime > square {
				break
			}
			ch = FilterPrimesOnly(ch, prime)
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
