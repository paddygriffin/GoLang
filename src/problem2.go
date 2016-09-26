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

	var primeNumber int // Store the 10000th Prime Number
	var count int // Counter
	var i int =1
	var z int
	var check bool = true

	for check{ // this is like a while loop using a bool so the check = false will exit the loop when called upon
		i++
		primeNumber=i

		if (i % i == 0) && (i % 1 == 0){ // check for Prime Number divide by itself and 1
			if i == 2 { // 2 is a prime number so add it
				count++
			} else { // Continue
				for z = 2; z < i; z++ {
					if i % z == 0{// If it is divisible then its not a Prime Number
						break // Try the next number
					} else if z == i-1  { //If its not divisible then its a Prime Number
						count++ //Add 1 to the count
						if count == 10001 {
							check = false//Exit the for loop as count is = to 10001
						}
					} else {
						continue
					}
				}
			}
		}
	}
	fmt.Print("The 10001st prime number is: ", primeNumber)// Print the 1000th Prime Number


}





