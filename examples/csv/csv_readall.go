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
         Icecream,"30"` //don't give the gap between data and `  it will give error
    
      r :=csv.NewReader(strings.NewReader(in))
    
  		record,_:= r.ReadAll() 

      for  i:= range record{
      
      fmt.Println(record[i]) //it will print the values line by line
    }
  }