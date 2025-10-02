package main

import (
	"fmt"
	"net/http"

	"github.com/lybpyn/zerswag"
)

func main() {
	handler := zerswag.New("/doc")

	// 添加根路径重定向到swagger.html
	http.HandleFunc("/doc/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/doc/" {
			http.Redirect(w, r, "/doc/swagger.html", http.StatusFound)
			return
		}
		handler.Route().Handler.ServeHTTP(w, r)
	})

	fmt.Println("Starting server on :8080")
	fmt.Println("Swagger UI available at: http://localhost:8080/doc/")
	fmt.Println("API docs available at: http://localhost:8080/doc/api-gate.api.json")

	http.ListenAndServe(":8080", nil)
}