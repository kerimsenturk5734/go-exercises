package main

import "fmt"

func main() {
	var x int = 25
	var y int = 7

	//Plus
	fmt.Println(x + y)

	//Minus
	fmt.Println(x - y)

	//Divide
	fmt.Println(x / y)

	//Multiplication
	fmt.Println(x * y)

	//Remainder
	fmt.Println(x % y)

	//Increment
	{
		//This operation just using as postfix
		x++
		y++
	}

	//Decrement
	{
		//This operation just using as postfix
		x--
		y--

		/*
			In Golang, you can't use increment and decrement operations in expressions.
			Because these operations are statement not expression

			x = x + y --> statement
			x + y --> expression
			x++ --> statement
			y-- --> statement

			That is why there is no prefix operation.

			You can find more details in --> https://go.dev/doc/faq#inc_dec.
		*/
	}

	//Logical AND operation for int values
	{
		//It converts the operands to binary and compare them with AND operator.
		/*
				0001 1001 --> 25
				0000 0111 --> 7
			AND____________________
				0000 0000 --> 0
		*/
		fmt.Println(x & y)

		//Logical OR operation for int values
	}
	//It converts the operands to binary and compare them with OR operator.
	{
		/*
				0001 1001 --> 25
				0000 0111 --> 7
			OR____________________
				0001 1111 --> 31
		*/
		fmt.Println(x | y)
	}

	// How to get power of a number? Let's code a function to do that!
	sub := 3
	expo := 4
	thirdPowOfTwo, err := pow(sub, expo)

	fmt.Printf("%v. power of %v is : %v, Error: %v\n", expo, sub, thirdPowOfTwo, err)

	// How to get average of numbers? Let's code a function to do that!
	avg := avg(1, 2, 4, 456, 123, 34, 67, 82, 34, 13)

	fmt.Printf("Average : %v", avg)
}

func pow(sub int, expo int) (int, error) {

	if expo == 0 {
		if sub == 0 {
			return 0, fmt.Errorf("0 power of 0 is undefined")
		}
		return 1, error(nil)
	}

	expo--
	ret, _ := pow(sub, expo)
	return sub * ret, error(nil)
}

func avg(nums ...int) float32 {
	var sum int
	for _, num := range nums {
		sum += num
	}

	return float32(sum) / float32(len(nums))
}
