package main
 
 import(
	"encoding/csv"
	"io"
	"log"
	"fmt"
	"strings"

	)


 func main(){
	
	in := `Name,Price
	carrot,50
	ice-cream,30` //Don't give the space to the data and ` symbol,it gives the error
	r := csv.NewReader(strings.NewReader(in))
	
	for{
	
	record,err :=r.Read()// read one by one record
	if err== io.EOF{//if u don't specify EOF as break it will be infinite loop
	break 
	}
	if err!=nil{
	log.Println(err)
	}
	fmt.Println(record)
	}
}