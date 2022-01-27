package main

import (
    "html/template"
    "fmt"
    "net/http"
)

const (
    Port = ":3000"
)
 
func serveStatic(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("template.html", "partial1.html")
    if err != nil {
        fmt.Println(err)
    }
    t.Execute(w, nil)
}
 
func main() {
    http.HandleFunc("/", serveStatic)
    // http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
    http.ListenAndServe(Port, nil)
}