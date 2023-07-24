package main

import (
	"fmt"
	"strconv"
)

// all functions takes an array(slice of numbers to add to each other and return the value of sum)

func sum(arr []float64) float64 {
	// take the first element and store it in valueint and add it to the intial value of sum then add the second element to valueint and add it to the new value of sum and so on for the next elements

	var sum float64
	for _, valueInt := range arr {
		sum += valueInt
	}
	return sum
}

func sub(arr []float64) float64 {
	// we intialize subtract as the first element in the slice we make then subtract the second element from it and return it to subtract so we have the new value that will be subtracted fromt the other numbers
	// we add the other values in the slice by making i increases in every iteration
	i := 1
	var subtract float64 = arr[0]
	for {
		subtract -= arr[i]
		if i+1 == len(arr) {
			break
		} else {
			i += 1
		}
	}
	return subtract

}

func multi(arr []float64) float64 {
	// take the first element and store it in valueint and multiply it with the intial value of multi then multiply the second element to valueint and multiply it with the new value of multi and so on for the next elements
	var multi float64 = 1
	for _, valueInt := range arr {
		multi *= valueInt
	}
	return multi
}

func div(arr []float64) float64 {
	// we intialize res as the first element in the slice we make then divide it by the second element and return it to res so we have the new value that will be divided by the other numbers
	// we add the other values in the slice by making i increases in every iteration
	var res float64 = arr[0]
	i := 1
	for {
		res = res / arr[i]
		if i+1 == len(arr) {
			break
		} else {
			i += 1
		}
	}
	return res
}

func main() {

	// create a map that its keys of type string and vlaues of type function that takes a slice as input float64 and return float64
	// assign each operator to their function so when entered they call their function
	m := map[string]func(arr []float64) float64{"+": sum, "-": sub, "*": multi, "/": div}

	arr := make([]float64, 0)
	var operator, j string

	// ask the user to enter the operator needed for the operation and assign a variable v1 which access the map to call the function needed
	fmt.Printf("choose operator: + - * / : ")
	fmt.Scanln(&operator)
	v1 := m[operator]

	for {

		fmt.Printf("enter a number: ")
		fmt.Scanln(&j)

		// when end is entered the program ends and return the value of the function as result
		if j == "end" {
			fmt.Println("your result is: ", v1(arr))
			break
			// we convert the type of input taken to float64 so it can be used inside math operations and add the values we take from the user to the end of the slice using append
		} else {
			x, _ := strconv.ParseFloat(j, 64)
			arr = append(arr, x)
		}

	}
}
