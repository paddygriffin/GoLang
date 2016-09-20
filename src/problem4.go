package main

import (
	//"bufio"
	//"os"
	"fmt"
)
/*
 Write a program that takes as input a positive integer and applies the following operations until the sequence begins to repeat:
  if the number is even, divide it by 2, but if the nubmer is odd, multiply it by 3 and add 1.
  The program should print the generated sequence to the screen.
 You might want to consider whether the program always terminates, and what will happen should the program encounter a 0.
*/

//takes a positive integer input
// if even, divide by 2
//if odd multiply by 3 and plus 1


func main() {
	//var evenNum int
	var num int
	var newnum int

	fmt.Print("Enter num: ")
	fmt.Scanln(&num)

	//finds if num is even or odd and does the math
	if (num  % 2)==0{
		newnum = num / 2
		fmt.Print(newnum)
	}else{
		newnum = (num * 3) + 1
		fmt.Print(newnum)
	}



}
