package main

import(
  "bitbucket.org/ekanna/datastore"
  _ "github.com/lib/pq"
  "database/sql"
  "encoding/csv"
  "encoding/json"
  "log"
  "fmt"
  "os"
  "strings"
)

var(
  logger = log.New(os.Stdout, "datastore: ", log.Ldate+log.Ltime+log.Lshortfile)
  db     *sql.DB
  ds     datastore.DS

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

  ItemId   string  `json:"itemid"`//to display data as a jsondata write like this
  ItemName string  `json:"itemname"`
  Uom      string  `json:"uom"`
  Amount   string  `json:"amount"`

}

func main(){
  defer db.Close()
  var data []Itempost
  err := ds.GetRows(&data,"pp.Itempost","")

  if err != nil {
        logger.Println(err)
        return
      }
  //logger.Printf("%+v", data)

  jsonData, err := json.Marshal(data)
  if err != nil {
    logger.Println(err)
  }
  fmt.Println(string(jsonData))


 file, err := os.Create("out.csv")//create the output.csv file
    if err != nil {
        log.Fatal("Cannot create file ", err)
    }
    defer file.Close()
      // Write a csv file
      csvw := csv.NewWriter(file)
      var header []string
      k := "itemid,itemname,uom,amount"

      k2 := strings.Split(k, ",")
      header = append(header, k2...)
      csvw.Write(header)

     for _, item := range data{

         var record []string
          record = append(record,
          item.ItemId, item.ItemName, item.Uom, item.Amount)
        csvw.Write(record)
      }

      csvw.Flush()
      return
}
