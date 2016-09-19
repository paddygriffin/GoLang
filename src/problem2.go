package main

import (
	"fmt"

)
//The ﬁrst six prime numbers in order
// are 2, 3, 5, 7, 11, and 13. So, for instance, the 4th prime number is 7.
// Write a program to ﬁnd the 10,001st prime number
//a prime number is a number that can divide by itself and 1

func main() {
	/*var num int
	var primeNum int
	var counter int = 0 //1

	for i := 0; i < 10001; i++{
		i++

		if (i % i == 0) && (i % 1 == 0){
			counter = counter + 1;
		}

		if i ==2{
			counter++
		}



		}
	fmt.Println("the prime number: ", primeNum)

	//variable

	//continuos loop until 10001
	//inside loop for checking

	//check if num is prime

	//print out the amount of prime numbers*/

//	var i int //loop 1
	var j int // loop 2
	var primeNumCount int //counting the prime

	for i := 1; i < 10001; i++{
		i++ //increment
		//checking
		if(i%1==0)&&(i%i==0){
			if i==2{
				primeNumCount++
			}else{
				for j =2;j<i;j++{
					if i % j ==0{
						break
					}else if j == i-1 {
						primeNumCount++
						if primeNumCount == 1001{
							primeNumCount = i
						}
					}else {
						continue
					}
				}
			}


		}//if

	}//for
	fmt.Println("the number:", primeNumCount)
}





