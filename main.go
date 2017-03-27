package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	api "./graphql"
)

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := api.ExecuteQuery(r.URL.Query()["query"][0])
		json.NewEncoder(w).Encode(result)
	})
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	fmt.Println("Now server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
