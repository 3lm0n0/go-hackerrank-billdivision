package main

import "fmt"

func main() {
	fmt.Println(bonAppetit())
}

func bonAppetit(bill []int32, k int32, b int32) {
	sum := int32(0)
	for i := int32(0); i < int32(len(bill)); i++ {
		if i != k {
			sum = sum + bill[i]
		}
	}
	sum = sum / 2
	if sum == b {
		fmt.Println("Bon Appetit")
		return
	}
	fmt.Println(b - sum)
	return
}

// sample input
// 4 1
// 3 10 2 9
// 12 --> (if this value is equal to Anna's portion value return "Bon Appetit")

// sample output
// 5

// 3 + 2 + 9 = 14
// 14 / 2 = 7
// 12 - 7 = 5
