package main

import (
	"fmt"
)

func main() {
	var a,b int
	a = 105
	b = 224
	for{
		if a,b = educlidean(a,b); a==0||b==0{
			break
		}
	}
	if a == 0 {
		fmt.Printf("max gxd is %d",b)
	}else {
		fmt.Printf("max gxd is %d",a)
	}

}

func educlidean(a int, b int) (int, int) {

	var r int
	if a < b {
		r = b - a
		b = r
	} else {
		r = a - b
		a = r
	}
	return a, b
}
