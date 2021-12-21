package main

import (
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"net/http"
	"strings"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/hello/")
	fmt.Fprintf(w, "hello, %s! I'm server3. \n", name)
}

func main() {
	http.HandleFunc("/hello/", Hello)
	otelHandler := otelhttp.NewHandler(http.HandlerFunc(Hello), "Hello")
	http.ListenAndServe(":8083", otelHandler)
}
