package main

import (
	"fmt"
	"math"
)

func main() {
	langs := []string { "Kotlin", "Go" , "Elixir" , "Phoenix",  "Reactjs",  "RxJava",  "RxAndroid"}
	fmt.Println("Hello world")
	fmt.Println("%v\n", math.Sqrt(2))

	for i :=0; i < len(langs) ; i++ {
		fmt.Printf("langs[%d] = %s\n", i , langs[i])

	}
}
