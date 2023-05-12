package main

import (
    "flag"
    "fmt"
    "net/http"
    "strconv"
)

var port int

func test(w http.ResponseWriter, req *http.Request) {
    fmt.Println("Request in 127.0.0.1:" + strconv.Itoa(port))
	fmt.Fprintf(w, req.Host)
}

func main() {
    flag.IntVar(&port, "p", 8000, "Provide a port number")
    flag.Parse()

    http.HandleFunc("/", test)

	// Finally, we call the `ListenAndServe` with the port
	// and a handler. `nil` tells it to use the default
	// router we've just set up.
	http.ListenAndServe(":" + strconv.Itoa(port), nil)
}