package main

import "fmt"

func main(){
	i := 42
	f := 3.142
	g :=  0.867 + 0.5i 

	fmt.Printf("%v is of Type %T\n", i, i) 
	fmt.Printf("%v is of Type %T\n", f, f)
	fmt.Printf("%v is of Type %T\n",  g, g)
}
