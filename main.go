package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var FlagListenAddress = flag.String("listenAddress", ":3000", "address to listen http requests on (for example 0.0.0.0:80)")

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Returned response to", r.RemoteAddr)
		fmt.Fprintf(w, "{\"ip\": \"%s\"}", strings.Split(r.RemoteAddr, ":")[0])
	})

	log.Println("Listening on", *FlagListenAddress)
	err := http.ListenAndServe(*FlagListenAddress, nil)
	if err != nil {
		panic(err)
	}
}
