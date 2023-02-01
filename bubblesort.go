package main

import (
	"fmt"
	"strconv"
)

func main() {
	TakeInput()
}

// BubbleSort function takes a slice of integers as input and sorts it
func BubbleSort(s []int) []int {

	//compare each element to each other
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-1; j++ {
			if s[j] > s[j+1] {
				Swap(s, j)
			}
		}
	}
	fmt.Println("Your sorted number is: ", s)
	return nil
}

// Swap function swaps the elements in the slice
func Swap(s []int, j int) {
	s[j], s[j+1] = s[j+1], s[j]
}

func TakeInput() []int {
	var input string

	//declare an empty slice of integer up to 10
	s := make([]int, 0, 10)

	fmt.Println(" Enter the string of integers you want to sort, up to 10")

	//loop through the slice and take input from the user
	for i := 0; i < 10; i++ {
		fmt.Println("Enter the integer , enter X to stop")
		fmt.Scan(&input)
		num, err := strconv.Atoi(input)

		//if the user enters X, break the loop and call the bubblesort function
		if input == "X" {
			break
		}

		if err != nil {
			fmt.Println("Invalid input , please enter a integer")
		} else {
			//append the integer to the slice
			s = append(s, num)
		}
		BubbleSort(s)
	}
	return nil
}
