package main

import (
	"fmt"
	"bufio"
	"os"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	//Write a program that accepts a user inputted string and prints its reverse
	//get an input from the user
	//try and reverse it some how

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	word1 := text
	reversed1 := Reverse(word1)
	fmt.Println(reversed1)


}