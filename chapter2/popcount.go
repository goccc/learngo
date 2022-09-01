package main

import "fmt"

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		//fmt.Printf("i = %d, pc[%d] = %d, pc[%d] = %d, byte(%d&1) =  %d\n",
		//i, i, pc[i], i/2, pc[i/2], i, byte(i&1))
	}
	fmt.Println(1 / 2)
	fmt.Println(3 & 1)

	fmt.Println(2 & 1)
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
func Pop(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1)
		fmt.Printf("x=%b\nx-1=%b\n", x, x-1)
	}
	return n
}

func main() {
	Pop(212)

}
