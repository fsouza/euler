package main

import (
	"fmt"
	"math"
)

func Generate(number float64) chan uint {
	square := uint(math.Sqrt(number))

	channel := make(chan uint)
	go func(){
		for i := uint(2); i <= uint(number); i++ {
			if i <= square {
				channel<-i
			}
		}
		channel<-4
	}()

	return channel
}

func Factor(i float64) uint {
	if i < 4 {
		return uint(i)
	}

	var greater	uint
	primesChannel := Generate(i)
	prime := <-primesChannel
	for j := 0; prime != 4; j++ {
		if prime != 4 {
			greater = prime
		}
		prime = <-primesChannel
	}
	return greater
}

func main() {
	fmt.Println(Factor(600851475143))
}
