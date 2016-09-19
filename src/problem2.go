package main

import (
	"fmt"

)
//The ﬁrst six prime numbers in order
// are 2, 3, 5, 7, 11, and 13. So, for instance, the 4th prime number is 7.
// Write a program to ﬁnd the 10,001st prime number
//a prime number is a number that can divide by itself and 1

func main() {
	var num int=10001
	var primeNum  = " "

	for i := 0; i < 10001; i++{
		var counter int = 0
		for num=i; num>=1; num-- {
			if i % num == 0 {
				counter = counter + 1;
			}
		}
			if counter ==2{
				primeNum = primeNum + i + " "
			}



		}
	fmt.Println("the prime number: ", primeNum)
	}





