package main

import(
  "bitbucket.org/ekanna/datastore"
  _ "github.com/lib/pq"
  "database/sql"
  "encoding/json"
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

	  ItemId   string  `json:"itemid"` //to display data as a jsondata write like this 
	  ItemName string  `json:"itemname"`
	  Uom      string  `json:"uom"`
	  Amount   float64 `json:"amount"`

}

func getItemdata(){
	  var data []Itempost
	  err := ds.GetRows(&data,"pp.Itempost","") // for displaying all columns use "",o/w specify column names

	  if err != nil {
			logger.Println(err)
	   		return
	   		}
	   		  logger.Printf("%+v", data)
		
	 jsonData, err := json.Marshal(data)
	  if err != nil {
	    logger.Println(err)
	  }
	  fmt.Println(string(jsonData))

   }

func main(){

 	 getItemdata()

}

