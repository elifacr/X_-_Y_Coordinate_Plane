package main

import (
	"fmt"
	"strconv"
)

//Handles the error.
func handle_error(err error) {
	if err != nil {
		switch err.Error() {
		case "strconv.ParseInt: error parsing, invalid syntax":
			fmt.Println("Invalid input. Please enter a valid integer value.")
		case "strconv.ParseInt: value out of range":
			fmt.Println("Input value is out of range. Please enter a value within the specified axis range.")
		default:
			fmt.Println("An error occurred:", err)
		}
	}
}

//Calculates the number of points with positive and negative
//values to display on the x and y axis.
func count_positive_negative(arr []int) (int, int) {
	positive_count := 0
	negative_count := 0

	for _, value := range arr {
		if value > 0 {
			positive_count++
		} else if value < 0 {
			negative_count++
		}
	}

	return negative_count, positive_count
}

func main() {
	var x_start, x_end int
	var y_start, y_end int

	var x_negative_str string
	n_number_length := 0

	//Draws the x axis.
	fmt.Print("Enter the starting value of the x axis (-): ")
	_, err := fmt.Scanln(&x_start)
	if err != nil {
		handle_error(err)
		return
	}
	fmt.Print("\nEnter the end value of the x axis (+): ")
	_, err = fmt.Scanln(&x_end)
	if err != nil {
		handle_error(err)
		return
	}
	
	//Draws the y axis.
	fmt.Print("\nEnter the starting value of the y axis (-): ")
	_, err = fmt.Scanln(&y_start)
	if err != nil {
		handle_error(err)
		return
	}
	fmt.Print("\nEnter the end value of the y axis (+): ")
	_, err = fmt.Scanln(&y_end)
	if err != nil {
		handle_error(err)
		return
	}

	//Calculates the number of points on the x and y axis.
	x_length := x_end - x_start + 1
	y_length := y_end - y_start + 1
	
	//Error checking.
	//An axis length cannot be zero.
	if x_length <= 0 || y_length <= 0 {
		fmt.Println("You entered the wrong range. Please try again.")
		return
	}

	//It keeps the points to be displayed on the x axis in an array.
	x := make([]int, x_length)
	for i := 0; i < x_length; i++ {
		x[i] = x_start + i
	}

	//It keeps the points to be displayed on the y axis in an array.
	y := make([]int, y_length)
	for i := 0; i < y_length; i++ {
		y[i] = y_start + i
	}

	//Prints the data of the points on the x and y axis to the screen.
	fmt.Println("\nx axis:", x)
	fmt.Println("y axis:", y)

	//Calculates the number of points with positive and negative values
	//to display on the x and y axis.
	x_negative, x_positive := count_positive_negative(x)
	y_negative, y_positive := count_positive_negative(y)

	//Prints the number of points with positive and negative values
	//to be displayed on the x and y axis.
	fmt.Println("\nNumber of positive x values:", x_positive)
	fmt.Println("Number of negative x values:", x_negative)
	fmt.Println("Number of positive y values:", y_positive)
	fmt.Println("Number of negative y values:", y_negative)
	fmt.Print("\n")

	//It determines the x and y values of the point P(x, y)
	//with data input from the user.
	x_value, y_value := 0, 0
	fmt.Print("Enter the x value of the point: ")
	_, err = fmt.Scanln(&x_value)
	if err != nil {
		handle_error(err)
		return
	}
	fmt.Print("Enter the y value of the point: ")
	_, err = fmt.Scanln(&y_value)
	if err != nil {
		handle_error(err)
		return
	}
	fmt.Print("\n")

	//It calculates the number of spaces that must be present before writing
	//the value of each point to be shown on the y axis.
	for i := 0; i < x_negative; i++ {
		x_negative_str = strconv.Itoa(x[i])
		n_number_length += len(x_negative_str)
	}
	number_of_spaces := n_number_length + x_negative*4

	//It forms the positive part of the y axis.
	for i := 0; i < y_positive; i++ {
		for j := 0; j < number_of_spaces; j++ {
			fmt.Print(" ")
		}
		fmt.Print("+")
		fmt.Println(y_length - y_negative - 1 - i)
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

	//It forms the origin point.
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
