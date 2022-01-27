package main
 
import (
	"fmt"
	"html/template"
	"net/http"
)
 
const (
	Port = ":5000"
)
 
func serveStatic(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, nil)
}
 
func main() {
	http.HandleFunc("/", serveStatic)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.ListenAndServe(Port, nil)
}