// Testing functions of the tempconv package
package main

import (
	"fmt"

	//"gopl.io/ch2/tempconv"
	"gopl.io/ch2/popcount"
)

func main() {

	k := tempconv.Kelvin(3)
	fmt.Printf("%s = %s\n", k, tempconv.KToC(k))
	k = tempconv.Kelvin(3)
	fmt.Printf("%s = %s\n", k, tempconv.KToF(k))
	c := tempconv.Celsius(3)
	fmt.Printf("%s = %s\n", c, tempconv.CToK(c))
	c = tempconv.Celsius(3)
	fmt.Printf("%s = %s\n", c, tempconv.CToF(c))
	f := tempconv.Fahrenheit(3)
	fmt.Printf("%s = %s\n", f, tempconv.FToC(f))
	f = tempconv.Fahrenheit(3)
	fmt.Printf("%s = %s\n", f, tempconv.FToK(f))

	fmt.Printf("%d\n", popcount.PopCount(18446744073709551615))
	fmt.Printf("%d\n", popcount.PopCount2(18446744073709551615))

}
