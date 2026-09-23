package main

import "fmt"

const (
	varApp = 10
)

func printverapp(version int8) {
	fmt.Println(version)
}

func main() {
	printverapp(varApp)
	for i := 0; i < 5; i++ {
		fmt.Println("hello world")
	}

}
