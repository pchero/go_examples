package main

import "fmt"

func main() {
	fmt.Printf("%s\n", "\"test\"")

	fmt.Printf(`"%s"\n`, "test")
}
