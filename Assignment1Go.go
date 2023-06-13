package main

import (
	"fmt"
	"sort"
)

// Program to find second largest number in array
func secondLargestNum() {
	arr := []int{2, 3, 5, 1, 4, 6, 8, 7, 9}
	sort.Ints(arr)
	if len(arr) <= 1 {
		fmt.Println("\n1).. No second largest number")
	} else {
		secondLargestNum := arr[(len(arr) - 2)]
		fmt.Println("\n1).. Second Largest Number : ", secondLargestNum)
	}
}

// Program to find if a number is PRIME or NOT
func primeNumber(a int) {
	flag := 1
	for i := 2; i < a; i++ {
		if a%i == 0 {
			flag = 0
		}
	}
	if flag == 1 {
		fmt.Println("\n2)..", a, "is a PRIME number")
	} else {
		fmt.Println("\n2)..", a, "is NON PRIME number")
	}
}

// Program to find the count of even, odd and zeros in an array
func countEvenOddZero() {
	array := []int{1, 2, 3, 4, 0, 0, 3, 2, 5, 4, 3, 1}
	count := [3]int{0, 0, 0} // index 0 for 0, index 1 for odd, index 2 for even
	for i := range array {
		if array[i] == 0 {
			count[0]++
		} else if array[i]%2 != 0 {
			count[1]++
		} else {
			count[2]++
		}
	}
	fmt.Println("\n3).. Counts of even, odd and zero")
	fmt.Println("\tZero : ", count[0])
	fmt.Println("\tOdd  : ", count[1])
	fmt.Println("\tEven : ", count[2])

}

// Program to find sum of upto "n" numbers in fibonacci series
func sumOfFibonaci(number int) {
	num0, num1 := 0, 1
	sum := num0 + num1
	for i := 2; i < number; i++ {
		num2 := num0 + num1
		sum += num2
		num0 = num1
		num1 = num2
	}
	fmt.Println("\n4).. Sum of Fibonaci is : ", sum)
}

func main() {
	//calling function to find second largest number
	secondLargestNum()

	//calling function to check prime or not
	primeNumber(70)

	//calling function to count even,odd and zero
	countEvenOddZero()

	//calling function to find sum of fibonaci series upto n
	sumOfFibonaci(10)
}
