package main

import (
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

func SumPrimes(number float64) uint64 {
	primesChannel := GetPrimes(number)
	prime := <-primesChannel
	sum := uint64(0)
	for {
		if (prime == 0) {
			break
		}
		sum += uint64(prime)
		prime = <-primesChannel
	}

	return sum
}
