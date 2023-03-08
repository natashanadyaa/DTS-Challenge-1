package main

import "fmt"

func main(){
	var x = 21
	fmt.Printf("%d\n", x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%%\n")
	var j bool = true
	fmt.Printf("%t\n", j)
	fmt.Printf("%d\n", j)
	fmt.Println("\u042F")
	var b = 21
	fmt.Printf("%d\n",b)
	var d = 25
	fmt.Printf("%d\n",d)
	e := 0xf
    fmt.Printf("%x\n", e)
	f := 0xF13
    fmt.Printf("%X\n", f)
	fmt.Println("U+042F\n")
	var k float64 = 123.456
	fmt.Printf("%.6f\n", k)
	fmt.Printf("%e\n", k)
}