package main

import (
	"fmt"
	"strconv"
)

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
	number_length := 0

	fmt.Print("Enter the starting value of the x axis: ")
	_, err := fmt.Scanln(&x_start)
	if err != nil {
		handle_error(err)
		return
	}
	fmt.Print("\nEnter the end value of the x axis: ")
	_, err = fmt.Scanln(&x_end)
	if err != nil {
		handle_error(err)
		return
	}
	
	fmt.Print("\nEnter the starting value of the y axis: ")
	_, err = fmt.Scanln(&y_start)
	if err != nil {
        handle_error(err)
        return
    }
	fmt.Print("\nEnter the end value of the y axis: ")
	_, err = fmt.Scanln(&y_end)
	if err != nil {
        handle_error(err)
        return
    }

	x_length := x_end - x_start + 1
	y_length := y_end - y_start + 1
	
	if x_length <= 0 || y_length <= 0 {
		fmt.Println("You entered the wrong range. Please try again.")
		return
	}

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

	x_negative, x_positive := count_positive_negative(x)
	y_negative, y_positive := count_positive_negative(y)

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
