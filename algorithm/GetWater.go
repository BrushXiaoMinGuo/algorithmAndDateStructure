package main

import "fmt"

var (
	a int
	b int
)

func main() {
	a = 0
	b = 0
	n := 0
	for {
		n++
		if got := fillB(3,5,4); got{
			break
		}

	}
	fmt.Println(n)

}

func fillB(capacityA, capacityB, expected int) bool{

	// fill b
	a += capacityA
	if a+b < capacityB {
		b += a
	} else {
		b = a + b - capacityB
	}

	// empty a
	a = 0

	// print a and b
	fmt.Printf("a capacity is:%d, b capacity is:%d \n", a, b)

	// check b
	if b == expected {
		fmt.Println("success")
		return true
	}
	return false
}
