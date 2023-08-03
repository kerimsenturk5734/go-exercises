package main

import "fmt"

func main() {

	localhost := "127.0.0.1"
	port := "8080"

	//fmt.Print()
	fmt.Println("--------------Print()--------------")
	{
		//You can print a lot of string in this func.
		//But it does not leave a gap between the string variables. You must do that manually.
		fmt.Print("Service is working on -->", localhost, ":", port)

		//It does not pass to below line after printing.
		fmt.Print(localhost, ":", port)

		//To pass to below line you can use escaping character \n
		fmt.Print("\n", localhost, ":", port, "\n")
	}

	//fmt.Println()
	fmt.Println("--------------Println()--------------")
	{
		//Println() leaves a gap between strings, and it passes the below line after printing.
		fmt.Println("Service is working on -->", localhost, ":", port)
	}

	//fmt.Printf()
	fmt.Println("--------------Printf()--------------")
	{
		//It does not pass to below line after printing.
		//Printf() waits a string parameter and what you want to print inside that string
		//You can format this string using some string formatters like %v, %T ...

		//%v presents value of variable
		fmt.Printf("Service is working on --> %v : %v \n", localhost, port)

		//%T presents type of variable
		fmt.Printf("Service is working on --> %T : %T \n", localhost, port)

		//%b presents binary equivalent of number variable
		fmt.Printf("Service is working on --> %b : %b \n", 53, 8080)

		//%X presents upper-case hexadecimal equivalent of number variable
		fmt.Printf("Service is working on --> %X : %X \n", 53, 8080)

		//%s presents string values of string variable, if you are printing a string you can use either %v and %s
		fmt.Printf("Service is working on --> %s : %v \n", localhost, 8080)

		//For more formatters check the https://pkg.go.dev/fmt
	}
}
