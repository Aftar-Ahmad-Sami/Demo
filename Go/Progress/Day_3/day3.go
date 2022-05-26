package main

import "fmt"

func main() {
	fmt.Println("Hello....Nice to meet you !!!")
	state:= true
	for state{
		var n int
		fmt.Print("Do you want to go through Problems? (1 = YES, 2 = NO)\t")
		fmt.Scanln(&n)
		if(n==1) {
			fmt.Print("\nLoading......\n\n")
			problem()
		}
		state = false
	}
}

/**
func add(args ...int) int {
	total := 0
	for _, value := range args {
		total += value
	}
	return total
}

func averageAndTotal(ary []float64) (float64, float64) {
	totalSum := 0.0
	for _, value := range ary {
		totalSum += value
	}
	avg := totalSum / float64(len(ary))
	return avg, totalSum
}

func averageFunc(ary []float64) float64 {
	totalSum := 0.0
	for _, value := range ary {
		totalSum += value
	}
	avg := totalSum / float64(len(ary))
	return avg
}
**/

////////////////////////////////////////////////////////////////////////////////////////

// func main() {

// 	//-------------------Functions-----------------------//

// 	var num int
// 	fmt.Print("\nEnter the size of array:\t")
// 	fmt.Scanln(&num)

// 	fmt.Print("\nEnter the elements of array:\t")
// 	ary := make([]float64, num)
// 	for i := 0; i < num; i++ {
// 		fmt.Scanf("%f", &ary[i])
// 	}

// 	//average := averageFunc(ary)
// 	average, totalSum := averageAndTotal(ary)
// 	fmt.Println("The total summation:=\t", totalSum)
// 	fmt.Println("The average :=\t", average)

// 	fmt.Println("Adding integers 1,2,3,4,5 = \t", add(1, 2, 3, 4, 5))
// }

// //-------------------Closure Functions-----------------------//

// addition := func(x, y int) int {
// 	return x + y
// }
// fmt.Println(addition(1, 1))

// //------------------------Defer---------------------------//
// func first() {
// 	fmt.Println("1st")
// }
// func second() {
// 	fmt.Println("2nd")
// }
// func main() {
// 	defer second()
// 	first()
// }

// /**This has 3 advantages: (1) it keeps our Close call near
// our Open call so its easier to understand, (2) if our func-
// tion had multiple return statements (perhaps one in
// an if and one in an else) Close will happen before
// both of them and (3) deferred functions are run even if
// a run-time panic occurs.**/

// //--------------------------Panic and Recover-------------------------//
// func main() {
// 	defer func() {
// 		str := recover()
// 		fmt.Println(str)
// 	}()
// 	panic("PANIC")
// }
