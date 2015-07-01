
###print hello world on terminal

```go

	package main
 
	import (
		"fmt"
	)

	func main(){
	fmt.Println("hello world")
	}

```

	save the file with .go exetension
	eg:helloworld.go

### how to run

	$go run helloworld.go
	
### how to generate binary code file

	$go build helloworld.go

	after executing this command it will generate binary code file(helloworld).To see list of files

		$ls

		helloworld
		helloworld.go

### how to execute binary code file 

		$./helloworld  <--|

### print hello web on web page

```go

	package main

	import(
			//"io"
			"fmt"
			"net/http"
	   )

	func HelloServer (w http.ResponseWriter,req *http.Request){
			// io.WriteString(w,"hello world \n")
			fmt.Fprintf(w,"helloweb")
	 }
	 
	 func main(){
			http.HandleFunc("/",HelloServer)
			fmt.Println("server running on 9000")
			http.ListenAndServe(":9000",nil)
	 }


```
		save it as helloweb.go

		$go run helloweb.go

		open browser type localhost:9000 or 127.0.0.1:9000

 

