package main

import (
	"os"
	"strings"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"fmt"
	"bufio"
	"strconv"
)
func main(){

	db, err := sql.Open("mysql", "seroot:admin@tcp(www.sesolarpumps.com:3306)/se_data_monitor")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


	var (
		id int
		name string
	)

	var filter string;

	//rows, err := db.Query("select unit_unique_identifier, unit_description from unit_master")

	c1 := make(chan string, 1)
	for {
		fmt.Println("")
		fmt.Println("Lateset raw data view with unit_id filter (c) 2016 SunEdison Energy India ")
		fmt.Println("Implemented in Go Lang by Preveen Padmanabhan <dxb1230@gmail.com>")
		fmt.Println("-------------------------------------------------------------------------")
	go func() {
		fmt.Print("Filter: ")
		consolereader := bufio.NewReader(os.Stdin)
		input, err := consolereader.ReadString('\n')
			if(err!=nil) {
			}
		c1 <- strings.Replace(input, "\n", "", 1)
	}()

	select {

		case filter = <-c1:
			
		case <- time.After(time.Second * 5):
			//filter = ""
			fmt.Println(" " + filter)
	}
	if strings.Compare(filter , "q") == 0  {	
		fmt.Println("Good Bye! have a nice Day!")
		os.Exit(0)
	}

	rows, err := db.Query("select unit_identifier, raw_data from unit_in_data where unit_identifier like '%" + filter + "%' order by id desc limit 12") 
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)

		if err != nil {
			log.Fatal(err)
		}
		name = strings.Replace(name, "#RD", "", 1)
		name = strings.Replace(name, strconv.Itoa(id), "", 1)
		log.Println(id, strings.Split(name, "*"))
		log.Println()
	}
	err = rows.Err()

	if err != nil {
		log.Fatal(err)

	}
	//	time.Sleep(5000 * time.Millisecond)
		println("---")
	}

}
