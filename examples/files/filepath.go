package main

import(
	"fmt"
	"path/filepath"
	)

func main(){
	paths := []string {
		"/home/example.js",
		"/home/dai/npm.odt",
		"/home/goplay/src/github.com/pavan/data.md", // don't forget to give comma at the end of the line
		}

	for _,p := range paths{

		dir,file := filepath.Split(p)
		fmt.Printf("Input: %q\n\tdir %q\n\tfile %q\n",p,dir,file)
	}
}