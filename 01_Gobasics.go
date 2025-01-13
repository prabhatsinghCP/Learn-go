// Lets start learning about go now....
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {

	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n ", v, lim)
	}
	return lim
}

func main() {
	sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Printf("sum is %v\n and tupe of sun is %T\n", sum, sum)

	//init and post statements are optionals...
	for sum < 100 {
		sum++
	}
	fmt.Println("sum is", sum)

	//infinity loop ->once we remove conditional part our loop run infinitly
	// for {
	// }

	if true {
		fmt.Print("singh is king ")
	}

	//important point about if and else statements....
	pow(3, 2, 10)
	pow(3, 4, 20)

}
