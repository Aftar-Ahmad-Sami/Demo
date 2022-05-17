package main

import "fmt"

func main() {

	//	fmt.Println("... Ready for Test Run? ....")
	var num int
	fmt.Print("Enter Upper Bound:\t")
	fmt.Scanln(&num)

	fmt.Println()
	for i := 1; i <= num; i++ {
		if i%2 == 0 && i%3 == 0 {
			fmt.Println(i, "\t- Divisible by 2 and 3")
		} else if i%2 == 0 {
			fmt.Println(i, "\t- Divisible by 2")
		} else {
			fmt.Println(i, "\t- Divisible by 3")
		}
	}

}
