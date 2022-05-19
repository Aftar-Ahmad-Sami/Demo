package main

import "fmt"

func main() {

	//	fmt.Println("... Ready for Test Run? ....")
	readyTxt := 3
	for readyTxt > 0 {
		fmt.Println("Loading...")
		readyTxt--
	}

	// var n int
	// fmt.Print("Enter the size of array:\t")
	// fmt.Scanln(&n)

	n := 5
	fmt.Print("\nEnter the elements of array:\t")
	var ary [5]int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &ary[i])
	}
	fmt.Print("\nThe Elements are:\t")
	for i := 0; i < n; i++ {
		fmt.Print(ary[i], " ")
	}

}

//--------------Loops and Conditional Statements---------------------//
// var num int
// //	var state int
// fmt.Print("Enter Upper Bound:\t")
// fmt.Scanln(&num)

// fmt.Println()
// for i := 1; i <= num; i++ {
// 	// if i%6 == 0 {
// 	// 	state = 6
// 	// } else if i%2 == 0 {
// 	// 	state = 2
// 	// } else if i%3 == 0 {
// 	// 	state = 3
// 	// } else {
// 	// 	state = 0
// 	// }

// 	// switch state {
// 	// case 6:
// 	// 	fmt.Println(i, "\t- Divisible by 2 and 3")
// 	// case 2:
// 	// 	fmt.Println(i, "\t- Divisible by 2")
// 	// case 3:
// 	// 	fmt.Println(i, "\t- Divisible by 3")
// 	// default:
// 	// 	fmt.Println(i, "\t- Not Divisible by 2 or 3")
// 	// }

// 	//-------------------------------------------------//
// 	if i%6 == 0 {
// 		fmt.Println(i, "\t- Divisible by 2 and 3")
// 	} else if i%2 == 0 {
// 		fmt.Println(i, "\t- Divisible by 2")
// 	} else if i%3 == 0 {
// 		fmt.Println(i, "\t- Divisible by 3")
// 	} else {
// 		fmt.Println(i, "\t- Not Divisible by 2 or 3")
// 	}
// }
