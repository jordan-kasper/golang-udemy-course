package main

import "fmt"

func main(){
	a,b,c,d,e,f := 0,1,2,3,4,5

	fmt.Printf("The binary value for %d is %b\n and hexidecimal value is %#x\n", a, a, a)
	fmt.Printf("The binary value for %d is %b\n and hexidecimal value is %#x\n", b, b, b)
	fmt.Printf("The binary value for %d is %b\n and hexidecimal value is %#x\n", c, c, c)
	fmt.Printf("The binary value for %d is %b\n and hexidecimal value is %#x\n", d, d, d)
	fmt.Printf("The binary value for %d is %b\n and hexidecimal value is %#x\n", e, e, e)
	fmt.Printf("The binary value for %d is %b\n and hexidecimal value is %#x\n", f, f, f)

	num := 255	
	fmt.Printf("Integer in hex: %x\n", num)

	num1 := 10 
	fmt.Printf("%b\n", num1)
}