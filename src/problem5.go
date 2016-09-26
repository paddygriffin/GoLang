package main

import (
	"fmt"
	"bufio"
	"os"

)

/*
Write a program that accepts four characters as input,
and outputs all permutations of those four characters.
What should your program do if two or more of the characters are the same?

*/

//accepts 4 cahracters only
//


func main() {
	//var word4 String;
	var word4 string
	var fourWord bool = true

	for fourWord{

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter four letter word: ")
	word4, _ = reader.ReadString('\n')
	fmt.Println(word4)
		if(len(word4)==4){
			fourWord = false
		}

	}




}

