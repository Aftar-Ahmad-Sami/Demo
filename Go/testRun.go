package main

import "fmt"

func main() {

	// fmt.Println("... Ready for Test Run? ....")
	// readyTxt := 3
	// for readyTxt > 0 {
	// 	fmt.Println("Loading...")
	// 	readyTxt--
	// }
	// fmt.Println("")

	//-------------------Functions-----------------------//

	var num int
	fmt.Print("\nEnter the size of array:\t")
	fmt.Scanln(&num)

	fmt.Print("\nEnter the elements of array:\t")
	ary := make([]float64, num)
	for i := 0; i < num; i++ {
		fmt.Scanf("%f", &ary[i])
	}

	//average := averageFunc(ary)
	average, totalSum := averageAndTotal(ary)
	fmt.Println("The total summation:=\t", totalSum)
	fmt.Println("The average :=\t", average)

	fmt.Println("Adding integers 1,2,3,4,5 = \t", add(1, 2, 3, 4, 5))
}

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
