/*
function random(a,b): returns an integer between "a" and "b",with each such integer being equally likely
and implement random(a,b) that only makes calls to rand.Intn(0,1).
T(n)=theta[lg(a-b)]

function unbiasedRandom: uses biasedRandom as a subroutine to return an unbiased answer,
returning 0 with probability 1/2 and 1 with probability 1/2.
T(n)=theta{ 1 / [2p(1-p)] }
function biasedRandom: outputs 1 with some probability p and 0 with probability 1−p (0<p<1)
*/
package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	for i := 1; i < 6; i++ {
		fmt.Println(random(3,7))
		fmt.Println(unbiasedRandom())
	}
}

func random(a,b int) int {
	if a == b {
		return a
	}
	if rand.Intn(2) == 0 {		//The random function generates integers in the interval [0, n)
		return random(a,(a+b)/2)
	} else {
		return random((a+b)/2+1,b)
	}
}

func unbiasedRandom() int {
	for true {
		x := rand.Intn(2)		//Assuming that it is the biasedRandom
		y := rand.Intn(2)
		if y != x {
			return x
		}
	}
	return 0
}