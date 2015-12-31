package main

import(
  "bitbucket.org/ekanna/datastore"
  _ "github.com/lib/pq"
  "database/sql"
  "log"
  "fmt"
  "os"
  )

var(
  logger = log.New(os.Stdout, "datastore: ", log.Ldate+log.Ltime+log.Lshortfile)
  db  *sql.DB
  ds  datastore.DS
)

func init() {
	var err error
	db, err = sql.Open("postgres", "user=lakshman password=lakshman  host=localhost  dbname=mydb port=5433 sslmode=disable")
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	ds = datastore.DS{Db: db}
}

type Itempost struct{

	  ItemId   string  
	  ItemName string  
	  Uom      string  
	  Amount   float64 
}

func getItemdata(){
	  var data []Itempost
	  err := ds.GetRows(&data,"pp.Itempost","")

	  if err != nil {
			logger.Println(err)
	   		return
	   		}
	   		  logger.Printf("%+v", data)
		
   }

func main(){

 	 getItemdata()

}

