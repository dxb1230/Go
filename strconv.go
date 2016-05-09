package main

import ( 
	"strconv"
)

func main() {

	i, err := strconv.Atoi("-42")

	if err  { 
		println(err)
	} else {
		 println(i)
	}
}
