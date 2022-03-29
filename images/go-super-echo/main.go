package main

import (
    "fmt"
    "log"
    "time"
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
        t := time.Now()
        s := fmt.Sprintf("Yodelay Hee Who from go at %s @ %v!!!!!", t.Format(time.Kitchen), t.Location())
        answer := &response1{Yodel: s}
        json, _ := json.Marshal(answer)
        fmt.Fprintf(w, string(json))
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}
