package main

import (
	"flag"
	"os"
	"net/http"
	"fmt"
	"log"
	"strings"
)

var (
	listenPort   int
	redirectCode int
)

func init() {
	flagset := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	flagset.IntVar(&listenPort, "port", 8080, "The port to listen on")
	flagset.IntVar(&redirectCode, "status", 301, "The HTTP status code to send")
	flagset.Parse(os.Args[1:])
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		host := strings.Split(r.Host, ":")[0]
		path := r.URL.Path
		query := ""
		if len(r.URL.RawQuery) > 0 {
			query = fmt.Sprint("?", r.URL.RawQuery)
		}
		redirectURI := fmt.Sprint("https://", host, path, query)
		http.Redirect(w, r, redirectURI, redirectCode)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", listenPort), nil))
}
