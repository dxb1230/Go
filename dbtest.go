package main

import (
	"os"
	"strings"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
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

	//rows, err := db.Query("select unit_unique_identifier, unit_description from unit_master")

	for {

	rows, err := db.Query("select unit_identifier, raw_data from unit_in_data where unit_identifier like '%" + os.Args[1] + "%' order by id desc limit 12") 
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)

		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, strings.Split(name, "*"))
	}
	err = rows.Err()

	if err != nil {
		log.Fatal(err)

	}
		time.Sleep(5000 * time.Millisecond)
		println("---")
	}
}
