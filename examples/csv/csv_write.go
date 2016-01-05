package main

import (
    "os"
    "log"
    "encoding/csv"
)

func main() {
    file, err := os.Create("output.csv")
    if err != nil {
        log.Fatal("Cannot create file ", err)
    }
    defer file.Close()
		var data = [][]string{{"Name","Price"},{"carrot", "50"}, {"Icecream", "30"}}

    writer := csv.NewWriter(file)

    for _, value := range data {
        err := writer.Write(value)
        if err != nil {
            log.Fatal("Cannot write to file ", err)
						return
        }
    }
		writer.Flush()
}
