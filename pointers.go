package main

import "fmt"


func main() {

	i, j := 42, 2701

	p := &i // point to i
	fmt.Println(*p) // read i through the pointers

	*p = 21 // set i through the pointers
	fmt.Println(i)

	p = &j // point to j

	*p = *p/37 // divide j through the pointers
	fmt.Println(j) // see the new value of j
}
