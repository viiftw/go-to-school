package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

var db []string
var mu sync.Mutex

var tmpl = `
<form action="/save">
  Message: <input name="message" type="text"><br>
  <input type="submit" value="Submit">
</form>
%v
`

func saveHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	r.ParseForm()
	messages, ok := r.Form["message"]
	if !ok {
		http.Error(w, "missing message", 500)
	}

	db = append(db, messages[0])

	http.Redirect(w, r, "/", 301)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "0")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var sb strings.Builder
	sb.WriteString("<ul>")
	for _, message := range db {
		sb.WriteString("<li>" + message + "</li>")
	}
	sb.WriteString("</ul>")

	fmt.Fprintf(w, tmpl, sb.String())
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/save", saveHandler)
	log.Fatal(http.ListenAndServe(":8980", nil))
}
