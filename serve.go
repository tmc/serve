// program serve serves the current working directory over http on port 8080
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	fmt.Printf("http://localhost:%v/\n", port)
	panic(http.ListenAndServe(fmt.Sprintf(":%v", port), http.FileServer(http.Dir(wd))))
}
