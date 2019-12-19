package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	port := flag.Int("port", 8080, "Specify the port number, default 8080")
	strPort := fmt.Sprintf("%s%d", ":", *port)
	fmt.Println("Starting up... on ", strPort)

	http.Handle("/theatron/", http.StripPrefix("/theatron/", http.FileServer(http.Dir("."))))
	err := http.ListenAndServe(strPort, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Shutting down...")

}
