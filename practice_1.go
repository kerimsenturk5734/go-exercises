package main

import "fmt"

func main() {
	// Find the which words have even length in given array.
	words := []string{"asd", "asdasd", "asdasd"}
	for _, word := range words {
		if hasItEvenLength(word) {
			fmt.Printf("%v has even length\n", word)
		} else {
			fmt.Printf("%v has odd length\n", word)
		}
	}

	//get name of the months via indicated number by using switch-case
	for i := 1; i <= 12; i++ {
		fmt.Printf("%v. month is %v\n", i, getMonthNameWithNumber(i))
	}
}

func hasItEvenLength(str string) bool {
	return (len(str) % 2) == 0
}

func getMonthNameWithNumber(num int) string {
	switch num {
	case 1:
		return "January"
	case 2:
		return "February"
	case 3:
		return "March"
	case 4:
		return "April"
	case 5:
		return "May"
	case 6:
		return "June"
	case 7:
		return "July"
	case 8:
		return "August"
	case 9:
		return "September"
	case 10:
		return "October"
	case 11:
		return "November"
	case 12:
		return "December"
	default:
		return ""
	}
}
