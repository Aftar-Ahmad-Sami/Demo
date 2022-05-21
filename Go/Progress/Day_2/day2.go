package main

import "fmt"

func main() {

	fmt.Println()

}

//----------------------Maps-----------------------------------//
// elementSymbol := [3]string{"H", "He", "Li"}
// element := make(map[string]int)
// for i := 1; i <= 3; i++ {
// 	element[elementSymbol[i-1]] = i
// }

// fmt.Println(element)

// number, state := element["Heoo"]
// fmt.Println(number, state)

//-----------------------Slice-------------------------------//
// var n int
// fmt.Print("Enter the size of array:\t")
// fmt.Scanln(&n)

// fmt.Print("\nSlicing Array + Appending Array\n\n")
// //ary := make([]int, n)
// ary := make([]int, n, 10)
// ary = append(ary, 1, 2, 3)
// n += 3
// sliceAry := ary[:]
// //sliceAry:=ary[0:]
// //sliceAry:=ary[:n]

// for i := 0; i < n; i++ {
// 	fmt.Print(ary[i], " ")
// 	fmt.Println(sliceAry[i])
// }

// fmt.Println()
// fmt.Print("\nCopying Array\n\n")
// copyAry := make([]int, 4)
// copy(copyAry, ary)

// for _, val := range copyAry {
// 	fmt.Print(val, " ")
// }
// fmt.Println()

//----------------------------Arrays---------------------------------//
// 	n := 5
// 	arr := [5]int{1, 2, 3, 4, 5}
// 	fmt.Print("\nEnter the elements of array:\t")
// 	var ary [5]int
// 	for i := 0; i < n; i++ {
// 		fmt.Scanf("%d", &ary[i])
// 	}
// 	fmt.Print("\nThe Input Elements are:\t")
// 	for i := 0; i < n; i++ {
// 		fmt.Print(ary[i], " ")
// 	}
// 	fmt.Print("\nThe Predetermined Elements are:\t")
// 	for _, value := range arr {
// 		fmt.Print(value, " ")
// 	}

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
