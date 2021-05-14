// Created by: haokai
// Created on: May 2021
// This program displays, "Convert-Fahrenheit-to-Celsius"

package main

import (
	"fmt"

	"github.com/leekchan/accounting"
)

func main() {
	// This function displays currency
	accountingFormater := accounting.Accounting{Precision: 2}
	// This function does addition
	var fahrenheit float64
	var celsius float64
	// input
	fmt.Println("This program is about Convert-Fahrenheit-to-Celsius program.")
	fmt.Println()
	fmt.Print("Enter the number of fahrenheit( ℉ ): ")
	fmt.Scanln(&fahrenheit)
	// process
	celsius = (fahrenheit - 32) * 5 / 9
	// output
	fmt.Println(accountingFormater.FormatMoney(fahrenheit), " ℉ =", accountingFormater.FormatMoney(celsius), "℃")
	fmt.Println("\nDone.")
}
