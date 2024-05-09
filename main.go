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

	return PositiveCount, NegativeCount
}

func main() {
	//Variables that hold the start and end values for the x and y axes.
	var XStart, XEnd int
	var YStart, YEnd int

	//Variables that hold the x and y values of the point P(x, y)
	//determined by data input from the user.
	var XValue, YValue int

	//Variable to keep negative values on the x-axis as string.
	var XNegativeStr string
	//Total length of negative numbers on the x axis.
	NNumberLengthX := 0

	//The variable YValueStr is created to calculate the character
	//length of the y value of point P.
	YValueStr := strconv.Itoa(YValue)
	//The variable NumberOfSpacesForPoint holds the number of spaces required
	//to the right of the y axis to mark point P in line with the x value.
	NumberOfSpacesForPoint := 0

	//If the P point is in the 1st or 4th region of the coordinate plane,
	//it is created to convert the positive numbers up to the x value
	//of the P point into a string.
	var X14Str string
	//If the P point is in the 1st or 4th region of the coordinate plane,
	//it is created to calculate the total character length of the positive
	//numbers up to the x value of the P point.
	var X14Length int

	//It gets the start and end value of the x axis
	//with data input from the user.
	fmt.Print("Enter the starting value of the x axis (-): ")
	_, Err := fmt.Scanln(&XStart)
	//Error checking.
	if Err != nil {
		HandleError(Err)
		return
	}
	//Error checking.
	if XStart >= 0 {
		fmt.Println("You entered the wrong starting value for the x axis. Please enter a negative number.")
		return
	}
	fmt.Print("\nEnter the end value of the x axis (+): ")
	_, Err = fmt.Scanln(&XEnd)
	//Error checking.
	if Err != nil {
		HandleError(Err)
		return
	}
	//Error checking.
	if XEnd <= 0 {
		fmt.Println("You entered the wrong end value for the x axis. Please enter a positive number.")
		return
	}
	
	//It gets the start and end value of the y axis
	//with data input from the user.
	fmt.Print("\nEnter the starting value of the y axis (-): ")
	_, Err = fmt.Scanln(&YStart)
	//Error checking.
	if Err != nil {
		HandleError(Err)
		return
	}
	//Error checking.
	if YStart >= 0 {
		fmt.Println("You entered the wrong starting value for the y axis. Please enter a negative number.")
		return
	}
	fmt.Print("\nEnter the end value of the y axis (+): ")
	_, Err = fmt.Scanln(&YEnd)
	//Error checking.
	if Err != nil {
		HandleError(Err)
		return
	}
	//Error checking.
	if YEnd <= 0 {
		fmt.Println("You entered the wrong end value for the y axis. Please enter a positive number.")
		return
	}

	//Calculates the number of points on the x and y axis.
	XLength := XEnd - XStart + 1
	YLength := YEnd - YStart + 1

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
	XPositive, XNegative := CountPositiveNegative(X)
	YPositive, YNegative := CountPositiveNegative(Y)

	//Prints the number of points with positive and negative values
	//to be displayed on the x and y axis.
	fmt.Println("\nNumber of positive x values:", XPositive)
	fmt.Println("Number of negative x values:", XNegative)
	fmt.Println("Number of positive y values:", YPositive)
	fmt.Println("Number of negative y values:", YNegative)
	fmt.Print("\n")

	//It determines the x and y values of the point P(x, y)
	//with data input from the user.
	fmt.Print("Enter the x value of the point: ")
	_, Err = fmt.Scanln(&XValue)
	//Error checking.
	if Err != nil {
		HandleError(Err)
		return
	}
	//Error checking.
	if XValue < XStart || XValue > XEnd {
		fmt.Println("You entered the x value of point P out of range. Please try again.")
	}
	fmt.Print("Enter the y value of the point: ")
	_, Err = fmt.Scanln(&YValue)
	//Error checking.
	if Err != nil {
		HandleError(Err)
		return
	}
	//Error checking.
	if YValue < YStart ||YValue > YEnd {
		fmt.Println("You entered the y value of point P out of range. Please try again.")
	}
	fmt.Print("\n")

	//Finds the location of point P(x, y).
	//
	//Point P(x, y) can be found in 7 different regions.
	//In the 1st region, it is in the form of P(+, +).
	//In the 2nd region, it is in the form of P(-, +).
	//In the 3rd region, it is in the form of P(-, -).
	//In the 4th region, it is in the form of P(+, -).
	//On the y axis, it is in the form of P(0, y).
	//On the x axis, it is in the form of P(x, 0).
	//At the origin, it is in the form of P(0, 0).
	if XValue > 0 && YValue > 0 {
		fmt.Printf("The point P(%d, %d) is in the 1st region of the x - y coordinate plane.\n", XValue, YValue)
		//Calculates the number of spaces required to the right of the y axis
		//to mark point P in line with the x value.
		for i := 0; i < XValue; i++ {
			X14Str = strconv.Itoa(i + 1)
			X14Length += len(X14Str)
		}
		NumberOfSpacesForPoint = X14Length + XValue*5 - 1 - len(YValueStr)
	} else if XValue < 0 && YValue > 0 {
		fmt.Printf("The point P(%d, %d) is in the 2nd region of the x - y coordinate plane.\n", XValue, YValue)
	} else if XValue < 0 && YValue < 0 {
		fmt.Printf("The point P(%d, %d) is in the 3rd region of the x - y coordinate plane.\n", XValue, YValue)
	} else if XValue > 0 && YValue < 0 {
		fmt.Printf("The point P(%d, %d) is in the 4th region of the x - y coordinate plane.\n", XValue, YValue)
		//Calculates the number of spaces required to the right of the y axis
		//to mark point P in line with the x value.
		for i := 0; i < XValue; i++ {
			X14Str = strconv.Itoa(i + 1)
			X14Length += len(X14Str)
		}
		NumberOfSpacesForPoint = X14Length + XValue*5 - 1 - len(YValueStr)
	} else if XValue == 0 && YValue != 0 {
		fmt.Printf("The point P(%d, %d) is on the y axis.", XValue, YValue)
	} else if XValue != 0 && YValue == 0 {
		fmt.Printf("The point P(%d, %d) is on the x axis.", XValue, YValue)
	} else if XValue == 0 && YValue == 0 {
		fmt.Printf("The point P(%d, %d) is on the origin.", XValue, YValue)
	}
	fmt.Print("\n\n")

	//It calculates the number of spaces that must be present before writing
	//the value of each point to be shown on the y axis.
	for i := 0; i < XNegative; i++ {
		XNegativeStr = strconv.Itoa(X[i])
		NNumberLengthX += len(XNegativeStr)
	}
	NumberOfInitialSpaces := NNumberLengthX + XNegative*4

	//It forms the positive part of the y axis.
	for i := 0; i < YPositive; i++ {
		for j := 0; j < NumberOfInitialSpaces; j++ {
			fmt.Print(" ")
		}
		fmt.Print("+")
		fmt.Print(YLength - YNegative - 1 - i)
		//It checks whether the y value of point P is above the x axis.
		//If the y value is above the x axis, it marks point P.
		if YValue == YLength - YNegative - 1 - i {
			for j := 0; j < NumberOfSpacesForPoint; j++ {
				fmt.Print(" ")
			}
			//Point P is in region 1.
			fmt.Println("*")
		} else {
			fmt.Print("\n")
		}
		for j := 0; j < NumberOfInitialSpaces; j++ {
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
		for j := 0; j < NumberOfInitialSpaces; j++ {
			fmt.Print(" ")
		}
		fmt.Println("|")
		for j := 0; j < NumberOfInitialSpaces; j++ {
			fmt.Print(" ")
		}
		fmt.Print(Y[YNegative - 1 - i])
		//It checks whether the y value of point P is above the x axis.
		//If the y value is above the x axis, it marks point P.
		if YValue == Y[YNegative - 1 - i] {
			for j := 0; j < NumberOfSpacesForPoint; j++ {
				fmt.Print(" ")
			}
			//Point P is in region 4.
			fmt.Println("*")
		} else {
			fmt.Print("\n")
		}
	}
}
