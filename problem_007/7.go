package main

import (
	"fmt"
)

func GenerateNumbersChannel() chan int {
	channel := make(chan int)

	go func(){
		for i := 2; ; i++ {
			channel <- i
		}
	}()

	return channel
}

func FilterPrimes(prime int, numbersChannel chan int) chan int {
	channel := make(chan int)

	go func(){
		for {
			i := <- numbersChannel
			if i % prime != 0 {
				channel <- i
			}
		}
	}()

	return channel
}

func PrimeAt(position int) int {
	primesChannel := make(chan int)

	go func(){
		channel := GenerateNumbersChannel()
		for {
			primeNumber := <- channel
			primesChannel <- primeNumber
			channel = FilterPrimes(primeNumber, channel)
		}
	}()

	var prime int
	for i := 0; i < position; i++ {
		prime = <- primesChannel
	}

	return prime
}

func main() {
	fmt.Println(PrimeAt(10001))
}
