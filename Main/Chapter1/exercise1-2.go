//Modify the echo program to print the index and value of the each of its arguments, one per line.
package main

import (
	"fmt"
	"os"
)

func main() {

	for idx, arg := range os.Args[0:] {
		fmt.Print(idx)
		fmt.Print(" ")
		fmt.Println(arg)

	}

}
