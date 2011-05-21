package main

import (
	"fmt"
	"math"
)

func GenerateCompleteSequence(number uint) chan uint {
	channel := make(chan uint)
	go func(){
		for i := uint(2); i <= number; i++ {
			channel <- i
		}
		channel <- 0
	}()

	return channel
}

func FilterSequence(prime uint, channel chan uint) chan uint {
	out := make(chan uint)
	go func(){
		i := <-channel
		for i != 0{
			if i % prime != 0 {
				out <- i
			}
			i = <-channel
		}
		out <- 0
	}()
	return out
}

func GetPrimes(number float64) chan uint {
	var greater uint
	square := uint(math.Sqrt(number))
	value := uint(number)

	out := make(chan uint)
	go func(){
		for i := uint(2); i <= uint(number); i++ {
			if i > square {
				break
			}
			greater = i
		}

		channel := GenerateCompleteSequence(value)
		prime := <-channel
		for prime != 0 {
			out <- prime
			if prime <= square  {
				channel = FilterSequence(prime, channel)
			}
			prime = <-channel
		}
		out <- 0
	}()

	return out
}

func Factor(i float64) uint {
	if i < 4 {
		return uint(i)
	}

	var greater uint
	value := uint(i)
	primesChannel := GetPrimes(i)
	prime := <-primesChannel
	for prime != 0 {
		if (prime != 0 && value % prime == 0) {
			greater = prime
		}
		prime = <-primesChannel
	}
	return greater
}

func main() {
	fmt.Println(Factor(600851475143))
}
