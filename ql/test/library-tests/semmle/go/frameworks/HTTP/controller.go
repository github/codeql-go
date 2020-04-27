package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerExample() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("r.Host", r.Host)
		fmt.Println("r.RequestURI", r.RequestURI)

		fmt.Println("r.BasicAuth()")
		user, password, _ := r.BasicAuth()
		fmt.Println(fmt.Sprintf("username %q, password %q", user, password))
		fmt.Println("---")
	})
	fmt.Println("running server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
