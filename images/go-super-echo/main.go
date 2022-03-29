package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "sync"
    "strconv"
)

var counter int
var mutex = &sync.Mutex{}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
    mutex.Lock()
    counter++
    fmt.Fprintf(w, strconv.Itoa(counter))
    mutex.Unlock()
}

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q, from Go!!", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/yodel", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Yodel-o-delay-hee-hoo from Go!")
    })

    http.HandleFunc("/increment", incrementCounter)

    log.Fatal(http.ListenAndServe(":8080", nil))

}
