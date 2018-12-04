package main

import (
    "fmt"
    "log"
    "net/http"
)

func helloWorld(name string) string {
	if name == "" {
		name = "Mundo"
		// name = "World"
	}
	return fmt.Sprintf("Hello, %s!", name)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, helloWorld(r.URL.Path[1:]))
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
