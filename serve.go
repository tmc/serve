// program serve serves the current working directory over http on port 8080
package main

import (
	"net/http"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	println("http://localhost:8080/")
	panic(http.ListenAndServe(":8080", http.FileServer(http.Dir(wd))))
}
