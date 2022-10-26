package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("p", 8080, "port")
	flag.Parse()
	lis := fmt.Sprintf("0.0.0.0:%d", *port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, This is for benchmark!"))
	})
	fmt.Printf("server listen on %s", lis)
	log.Fatal(http.ListenAndServe(lis, nil))
}
