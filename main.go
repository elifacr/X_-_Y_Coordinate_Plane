package main

import (
	"fmt"
	"strconv"
)

//Handles the error.
func HandleError(Err error) {
	if Err != nil {
		switch Err.Error() {
		case "strconv.ParseInt: error parsing, invalid syntax":
			fmt.Println("Invalid input. Please enter a valid integer value.")
		case "strconv.ParseInt: value out of range":
			fmt.Println("Input value is out of range. Please enter a value within the specified axis range.")
		default:
			fmt.Println("An error occurred:", Err)
		}
	}
}

//Calculates the number of points with positive and negative
//values to display on the x and y axis.
func CountPositiveNegative(Arr []int) (int, int) {
	PositiveCount := 0
	NegativeCount := 0

	for _, Value := range Arr {
		if Value > 0 {
			PositiveCount++
		} else if Value < 0 {
			NegativeCount++
		}
	}

	return NegativeCount, PositiveCount
}

func main() {
	var XStart, XEnd int
	var YStart, YEnd int

	var XNegativeStr string
	NNumberLength := 0

	//Draws the x axis.
	fmt.Print("Enter the starting value of the x axis (-): ")
	_, Err := fmt.Scanln(&XStart)
	if Err != nil {
		HandleError(Err)
		return
	}
	fmt.Print("\nEnter the end value of the x axis (+): ")
	_, Err = fmt.Scanln(&XEnd)
	if Err != nil {
		HandleError(Err)
		return
	}
	
	//Draws the y axis.
	fmt.Print("\nEnter the starting value of the y axis (-): ")
	_, Err = fmt.Scanln(&YStart)
	if Err != nil {
		HandleError(Err)
		return
	}
	fmt.Print("\nEnter the end value of the y axis (+): ")
	_, Err = fmt.Scanln(&YEnd)
	if Err != nil {
		HandleError(Err)
		return
	}

	//Calculates the number of points on the x and y axis.
	XLength := XEnd - XStart + 1
	YLength := YEnd - YStart + 1
	
	//Error checking.
	//An axis length cannot be zero.
	if XLength <= 0 || YLength <= 0 {
		fmt.Println("You entered the wrong range. Please try again.")
		return
	}

	//It keeps the points to be displayed on the x axis in an array.
	X := make([]int, XLength)
	for i := 0; i < XLength; i++ {
		X[i] = XStart + i
	}

	//It keeps the points to be displayed on the y axis in an array.
	Y := make([]int, YLength)
	for i := 0; i < YLength; i++ {
		Y[i] = YStart + i
	}

	//Prints the data of the points on the x and y axis to the screen.
	fmt.Println("\nx axis:", X)
	fmt.Println("y axis:", Y)

	//Calculates the number of points with positive and negative values
	//to display on the x and y axis.
	XNegative, XPositive := CountPositiveNegative(X)
	YNegative, YPositive := CountPositiveNegative(Y)

	//Prints the number of points with positive and negative values
	//to be displayed on the x and y axis.
	fmt.Println("\nNumber of positive x values:", XPositive)
	fmt.Println("Number of negative x values:", XNegative)
	fmt.Println("Number of positive y values:", YPositive)
	fmt.Println("Number of negative y values:", YNegative)
	fmt.Print("\n")

	//It determines the x and y values of the point P(x, y)
	//with data input from the user.
	XValue, YValue := 0, 0
	fmt.Print("Enter the x value of the point: ")
	_, Err = fmt.Scanln(&XValue)
	if Err != nil {
		HandleError(Err)
		return
	}
	fmt.Print("Enter the y value of the point: ")
	_, Err = fmt.Scanln(&YValue)
	if Err != nil {
		HandleError(Err)
		return
	}
	fmt.Print("\n")

	//It calculates the number of spaces that must be present before writing
	//the value of each point to be shown on the y axis.
	for i := 0; i < XNegative; i++ {
		XNegativeStr = strconv.Itoa(X[i])
		NNumberLength += len(XNegativeStr)
	}
	NumberOfSpaces := NNumberLength + XNegative*4

	//It forms the positive part of the y axis.
	for i := 0; i < YPositive; i++ {
		for j := 0; j < NumberOfSpaces; j++ {
			fmt.Print(" ")
		}
		fmt.Print("+")
		fmt.Println(YLength - YNegative - 1 - i)
		for j := 0; j < NumberOfSpaces; j++ {
			fmt.Print(" ")
		}
		fmt.Println("|")
	}

	//It forms the negative part of the x axis.
	for i := 0; i < XNegative; i++ {
		fmt.Print(X[i])
		fmt.Print(" -- ")
	}

	//It forms the origin point.
	fmt.Print("O")

	//It forms the positive part of the x axis.
	for i := 0; i < XPositive; i++ {
		fmt.Print(" -- ")
		fmt.Print("+")
		fmt.Print(X[XNegative + 1 + i])
	}
	fmt.Print("\n")

	//It forms the negative part of the y axis.
	for i := 0; i < YNegative; i++ {
		for j := 0; j < NumberOfSpaces; j++ {
			fmt.Print(" ")
		}
		fmt.Println("|")
		for j := 0; j < NumberOfSpaces; j++ {
			fmt.Print(" ")
		}
		fmt.Println(Y[YNegative - 1 - i])
	}
}
