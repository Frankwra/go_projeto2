package main

import "fmt"

func main() {

	for inicial := 1; inicial <= 100; inicial++ {
		if inicial%3 == 0 {
			fmt.Println(inicial)
		}
	}
}
