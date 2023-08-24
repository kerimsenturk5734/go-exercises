package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//fmt.Scan()
	/*
			Scan scans text read from standard input, storing successive
		 	space-separated values into successive arguments. Newlines count
		 	as space. It returns the number of items successfully scanned.
			If that is less than the number of arguments, err will report why.
	*/
	var d, e, f string

	fmt.Println("Please enter value")
	valCount, err := fmt.Scan(&d, &e, &f)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Values: %v, %v, %v  Count: %v \n", d, e, f, valCount)

	}
	//fmt.Scanln()
	/*
			Scanln is similar to Scan, but stops scanning at a newline and
		 	after the final item there must be a newline or EOF.
	*/

	var a, b, c string

	fmt.Println("Please enter 3 value")
	valCount1, err1 := fmt.Scanln(&a, &b, &c)

	if err1 != nil {
		log.Fatal(err1)
	} else {
		fmt.Printf("Values: %v, %v, %v  Count: %v \n", a, b, c, valCount1)

	}

	//bufio.Reader()
	fmt.Println("Please enter a value")
	reader := bufio.NewReader(os.Stdin)
	line, isPrefix, err2 := reader.ReadLine() //
	// line, isPrefix, err2 := reader.ReadString('\n') it read until come \n character

	if err2 != nil {
		log.Fatal(err2)
	} else {
		fmt.Printf("Input line is: %s, isPrefix: %v", line, isPrefix)
	}

	//bufio.Scannner()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err3 := scanner.Err()
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Printf("read line: %s\n", scanner.Text())

	scanner.Scan()
	fmt.Printf("read line %s\n", scanner.Text())
}
