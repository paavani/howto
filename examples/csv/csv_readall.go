package main
import(
  "encoding/csv"
  "fmt"
  "strings"
  "log"
)
func main() {
  in := `Name,Price
         Carrot,"50"
         Icecream,"30"` //don't put the gap between data and `  it will give error
    
      r :=csv.NewReader(strings.NewReader(in))
    
  		record,err := r.ReadAll() //ReadAll reads all data at a time,no need to write EOF and for loop

  		if err != nil {
  			log.Println(err)
  		}
    // In single line display the data
  		fmt.Println(record)
	
  }