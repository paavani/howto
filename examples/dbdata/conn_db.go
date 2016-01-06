package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db, err := sql.Open("postgres", "user=pavan   password =pavan*** port=5433  host=localhost dbname = mydb sslmode =disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("create table if not exists hello(wishes varchar(10))")
	if err != nil {
		log.Fatal(err)

	//defer db.Exec( "drop table hello") 
		}
		
		log.Println("table created")

}


