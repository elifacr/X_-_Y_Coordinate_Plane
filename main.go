package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x_start, x_end int
	var y_start, y_end int

	var x_negative, x_positive int
	var y_negative, y_positive int

	var x_negative_str string
	number_length := 0

	fmt.Print("Enter the starting value of the x axis: ")
	fmt.Scanln(&x_start)
	fmt.Print("\nEnter the end value of the x axis: ")
	fmt.Scanln(&x_end)
	
	fmt.Print("\nEnter the starting value of the y axis: ")
	fmt.Scanln(&y_start)
	fmt.Print("\nEnter the end value of the y axis: ")
	fmt.Scanln(&y_end)

	x_length := x_end - x_start + 1
	y_length := y_end - y_start + 1

	//Points to display on the x axis.
	x := make([]int, x_length)
	for i := 0; i < x_length; i++ {
		x[i] = x_start + i
	}

	//Points to display on the y axis.
	y := make([]int, y_length)
	for i := 0; i < y_length; i++ {
		y[i] = y_start + i
	}

	fmt.Println("\nx axis:", x)
	fmt.Println("y axis:", y)

	for i := 0; i < x_length; i++ {
		if x[i] < 0 {
			x_negative++
		}
	}
	x_positive = x_length - x_negative - 1

	for i := 0; i < y_length; i++ {
		if y[i] < 0 {
			y_negative++
		}
	}
	y_positive = y_length - y_negative - 1

	fmt.Println("\nNumber of positive x values:", x_positive)
	fmt.Println("Number of negative x values:", x_negative)
	fmt.Println("Number of positive y values:", y_positive)
	fmt.Println("Number of negative y values:", y_negative)
	fmt.Print("\n")

	//It calculates the number of spaces that must be present before writing
	//the value of each point to be shown on the y axis.
	for i := 0; i < x_negative; i++ {
		x_negative_str = strconv.Itoa(x[i])
		number_length += len(x_negative_str)
	}
	number_of_spaces := number_length + x_negative*4
	//It forms the positive part of the y axis.
	for i := 0; i < y_positive; i++ {
		for j := 0; j < number_of_spaces; j++ {
			fmt.Print(" ")
		}
		fmt.Print("+")
		fmt.Println(y_length - y_negative - i)
		for j := 0; j < number_of_spaces; j++ {
			fmt.Print(" ")
		}
		fmt.Println("|")
	}

	//It forms the negative part of the x axis.
	for i := 0; i < x_negative; i++ {
		fmt.Print(x[i])
		fmt.Print(" -- ")
	}
	fmt.Print("O")
	//It forms the positive part of the x axis.
	for i := 0; i < x_positive; i++ {
		fmt.Print(" -- ")
		fmt.Print("+")
		fmt.Print(x[x_negative + 1 + i])
	}
	fmt.Print("\n")

	//It forms the negative part of the y axis.
	for i := 0; i < y_negative; i++ {
		for j := 0; j < number_of_spaces; j++ {
			fmt.Print(" ")
		}
		fmt.Println("|")
		for j := 0; j < number_of_spaces; j++ {
			fmt.Print(" ")
		}
		fmt.Println(y[y_negative - i - 1])
	}
}
