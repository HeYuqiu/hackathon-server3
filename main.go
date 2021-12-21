package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s! I'm server3. \n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/hello/:name", Hello)
	log.Fatal(http.ListenAndServe(":8083", router))
}
