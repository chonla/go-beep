package main

import "fmt"

var outFn = fmt.Print

func main() {
	outFn("\a")
}
