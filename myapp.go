package main
import (
		"fmt"
		"net/http"
		"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
		
		h, _ := os.Hostname()
		fmt.Fprintf(w, "222 Hi there, I am served from %s!", h)
}

func main() {

		http.HandleFunc("/", handler)
		http.ListenAndServe(":8484", nil) 
}
