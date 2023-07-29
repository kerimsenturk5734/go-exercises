package main

import (
	"fmt"
	"strconv"
)

func main() {
	//intToFloat conversion
	fmt.Println("---------------Integer To Float-------------")
	var price float32 = 10.5
	var off int = 10

	fmt.Printf("Value: %v, Type: %T \n", price, price)
	fmt.Printf("Value: %v, Type: %T \n", off, off)

	newPrice := price + ((price * float32(off)) / 100)

	fmt.Printf("Value: %v, Type: %T \n", newPrice, newPrice)

	//floatToInt Conversion
	fmt.Println("---------------Float To Integer-------------")
	var height float32 = 1.90
	var weight int = 85

	fmt.Printf("Value: %v, Type: %T \n", height, height)
	fmt.Printf("Value: %v, Type: %T \n", weight, weight)

	bodyIndex := weight / int(height*height)

	fmt.Printf("Value: %v, Type: %T \n", bodyIndex, bodyIndex)

	//int16-int8 conversion
	fmt.Println("---------------int8 To int16-------------")
	var val1 int8 = 127
	var val2 int16 = 425

	fmt.Printf("Value: %v, Type: %T \n", val1, val1)
	fmt.Printf("Value: %v, Type: %T \n", val2, val2)

	newVal1 := int16(val1) + val2
	fmt.Printf("Value: %v, Type: %T \n", newVal1, newVal1)

	fmt.Println("---------------int16 To int8-------------")
	var val3 int8 = 127
	var val4 int16 = 5

	fmt.Printf("Value: %v, Type: %T \n", val3, val3)
	fmt.Printf("Value: %v, Type: %T \n", val4, val4)

	newVal2 := val3 - int8(val4) //If val4 is in range of int8 there is not a problem
	fmt.Printf("Value: %v, Type: %T \n", newVal2, newVal2)

	//If val4 is out range of int8 the value returning to start value -128 of int8
	val4 = 128
	fmt.Printf("Value: %v, Type: %T \n", int8(val4), int8(val4))

	//We can not use int(string) or string(int) expression for conversion between int string.
	//int-string conversion
	fmt.Println("---------------int To string-------------")
	var val5 int = 24
	var str2 string = strconv.Itoa(val5) + "abcd"

	fmt.Printf("Value: %v, Type: %T \n", str2, str2)

	fmt.Println("---------------string To int-------------")
	var str3 string = "10"
	//var str3 string = "a10" --> Invalid syntax error
	//var str3 string = "10abcd" --> Invalid syntax error

	var convertedStr, err = strconv.Atoi(str3)

	fmt.Printf("Value: %v, Type: %T \n", convertedStr, convertedStr)
	fmt.Printf("Value: %v, Type: %T \n", err, err)

	fmt.Println("---------------ASCII Value-------------")
	//To get string value from integer ASCII value
	var asciiJ int32 = 106
	var strJ = string(asciiJ) + string(asciiJ) // It gets integer value and assign its ASCII equivalent to strJ

	fmt.Printf("Value: %v, Type: %T \n", strJ, strJ)

}
