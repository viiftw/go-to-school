package main

import (
	"fmt"
	"log"
	"net/http"
)

const content = `

<html>
   <head>
       <script>
          window.onload = function() {
             var params = new URLSearchParams(window.location.search);
             p = document.getElementById("content")
             p.innerHTML = params.get("message")
	     };
       </script>
   </head>
   <body>
       <p id="content"></p>
   </body>
</html>
`

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "0")
	fmt.Fprintf(w, content)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8980", nil))
}
