// Lets start learning about go now....
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func pow(x, n, lim float64) float64 {

	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		//variable inside the statement are accessable here as well
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

	//lets learn switch case...
	//break is by default applicable on each case ..
	//not only number can be used , other thing also can be used...

	// switch val := rand.Intn(3) + 1; val {
	// case 1:
	// 	fmt.Print("prabhat")

	// case 2:
	// 	fmt.Print("Singh")
	// case 3:
	// 	fmt.Print("Uber")
	// }

	//syntex -2
	val := rand.Intn(3) + 1
	switch val {
	case 1:
		fmt.Print("prabhat")

	case 2:
		fmt.Print("Singh")
	case 3:
		fmt.Print("Uber")
	}

	//switch statements without condition
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Print("cur time is less than 12")
	case t.Hour() < 15:
		fmt.Print("cur time is between 12 and 15")
	case t.Hour() < 24:
		fmt.Print("day is about to finish , wrap your task!")
	}

}
