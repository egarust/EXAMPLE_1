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
	fmt.Println("hello world")
}
