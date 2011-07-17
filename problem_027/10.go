package main

import (
	"math"
)

func GenerateCompleteSequence(number int) chan int {
	channel := make(chan int)
	go func(){
		for i := int(2); i <= number; i++ {
			channel <- i
		}
		channel <- 0
	}()

	return channel
}

func FilterSequence(prime int, channel chan int) chan int {
	out := make(chan int)
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

func GetPrimes(number float64) chan int {
	var greater int
	square := int(math.Sqrt(number))
	value := int(number)

	out := make(chan int)
	go func(){
		for i := int(2); i <= int(number); i++ {
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

func SumPrimes(number float64) int64 {
	primesChannel := GetPrimes(number)
	prime := <-primesChannel
	sum := int64(0)
	for {
		if (prime == 0) {
			break
		}
		sum += int64(prime)
		prime = <-primesChannel
	}

	return sum
}
