package main

import "fmt"

func main() {
	for i := 60; i < 121; i++ {
		fmt.Printf("%03d \t %08b \t %#X \t %q\n", i, i, i, i)
	}
}
