package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
)

type response1 struct {
    Yodel string `json:"yodel"`
}

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from GO!!!")
    })

    http.HandleFunc("/yodel", func(w http.ResponseWriter, r *http.Request){
        answer := &response1{Yodel: "Yodelay Hee Who!!!!!"}
        json, _ := json.Marshal(answer)
        fmt.Fprintf(w, string(json))
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}
