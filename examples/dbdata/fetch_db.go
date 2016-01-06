package main
import(
   "database/sql"
   "log"
   _"github.com/lib/pq"
)

var (
	empid int 
	ename string 
	)

	func main(){
	db,err := sql.Open("postgres","user=lakshmi-2826 password=postgres dbname=mydb port=5433 host=localhost sslmode=disable")
	
	//rows,err := db.Query("select empid,ename from employees where empid = $1",7934)  //for one selected row

        rows,err := db.Query("select empid,ename from employees ")  //for all rows of table displaying
	if err != nil {
	log.Fatal(err)
	}
  	
	defer rows.Close()
 
    for rows.Next() {
	err := rows.Scan(&empid, &ename)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(empid, ename)
	}
	err =rows.Err()
	if err != nil {
	log.Fatal(err)
	}
	}

