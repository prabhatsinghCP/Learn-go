// defer is very interesting concept in go
// how defer works:
// 1)the code starts exicuting step by step and when
// defer hit then that statement is insialised but not
// executed when surroundng gets exicuted then defer statement exicutes
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// func main() {
// 	defer fmt.Println("world")

// 	fmt.Print("Hello")
// }

//2)if there will be multiple defer statement ..
//they will follow stack property (LIFO) after executing surrounding statements..

// func main() {
// 	defer fmt.Print("Prabhat ")
// 	defer fmt.Print("am ")
// 	fmt.Printf("I ")

// }

//Reading the file from external files......
//3 where should we use defer ..
// while cleaning , closing something , release some data .etc

// func main() {
// 	file, err := os.Open("file.txt")

// 	if err != nil {
// 		fmt.Println("we got some error and error is :", err)
// 		return
// 	}

// 	defer file.Close()

// 	content, err := ioutil.ReadFile("file.txt")

// 	if err != nil {
// 		fmt.Println("The file does not read properly and err is :", err)
// 		return
// 	}

// 	fmt.Println("the content we got is :")
// 	fmt.Println(string(content))

// }

func main() {
	file, err := os.Open("file.txt")

	if err != nil {
		fmt.Println("Hey! I got some error please see it : ", err)
		return
	}

	defer file.Close()

	content, err := ioutil.ReadFile("file.txt")

	if err != nil {
		fmt.Println("Hey got some error ,please debug it : ", err)
		return
	}

	fileContent := string(content)

	lines := strings.Split(fileContent, "\n")

	if len(lines) < 2 {
		fmt.Println("file does not contain enough data")
		return
	}

	for i := 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			continue
		}
		fmt.Printf("%d th line is: %s\n", i+1, lines[i])
	}
}
