package main

import(
	"log"
	"net/http"
	"html/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, r)
}

func main() {
	http.HandleFunc("/", index)

	log.Println("Server running on port 8000...")
	http.ListenAndServe(":8000", nil )
}
