package main

import (
	"fmt"
	"net/http"
	"os"
)

// PORT=8080 go run env_config.go will listen on port 8080
func main() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
	}

	fmt.Fprint(res, "The homepage")
}
