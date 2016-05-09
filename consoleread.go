//socketloop 
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	c1 := make(chan string, 1)

	go func() {
	consolereader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")

	input, err := consolereader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Println(input)
	c1 <- input

	}()

	select {
		case name := <-c1:
			fmt.Println(name)

		case <-time.After(time.Second * 5):
			fmt.Println("Timeout")

	}
}
