package main 

import (
	"fmt"
	"net/http"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World")
}

func testHealth(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "OK")
}

func main(){
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", testHealth)

	fmt.Println("Server listening on port 8080...")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}