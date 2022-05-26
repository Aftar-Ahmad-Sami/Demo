package main

import "fmt"

func problem() {
	loop := true
	var number int
	for loop {
		fmt.Print("Enter the number = \t")
		fmt.Scanln(&number)

		result, state := half(number)
		if state == true {
			fmt.Println("The number is even and it's half is ", result)
		} else {
			fmt.Println("The number is odd and it's half is ", result)
		}

		fmt.Print("\n\nDo you want to continue? (0=NO,1=YES)\t")
		fmt.Scanln(&number)
		if number == 0 {
			loop = false
		}

	}
}

func half(number int) (int, bool) {
	halfNumber := number / 2
	if number%2 == 0 {
		return halfNumber, true
	} else {
		return halfNumber, false
	}
}
