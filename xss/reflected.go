package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Furthermore, some browsers are removing these heuristic-based XSS protections, for example if you are running Chrome 78 or above, you don’t need to include the w.Header().Set("X-XSS-Protection", "0") line for this attack to work.
	// XSS Auditor (removed)
	// https://www.chromestatus.com/feature/5021976655560704
	w.Header().Set("X-XSS-Protection", "0")

	messages, ok := r.URL.Query()["message"]
	if !ok {
		messages = []string{"hello, world"}
	}
	fmt.Fprintf(w, "<html><p>%v</p></html>", messages[0])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8980", nil))
}

// http://localhost:8980/?message=%3Cscript%3Ealert(1)%3C/script%3E
