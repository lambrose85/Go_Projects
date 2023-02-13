package main

import (
	"fmt"
	"os"
)

var shift = 5

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//for i, c := range "abc" {
	//	fmt.Println(i, " => ", string(c))
	//}

	dat, err := os.ReadFile("test.txt")
	check(err)
	fmt.Print(string(dat))
	encrypt(string(dat))
}
func encrypt(input string) {
	temp := []rune(input)
	for i, v := range input {
		fmt.Print("data: ", v, " ", i)
	}
	for i := 0; i < len(input); i++ {
		temp[i] = temp[i] + 5
	}
	fmt.Print("\n", temp)
	test := string(temp)
	fmt.Println(test)
}
func decrypt(input string) {

}
