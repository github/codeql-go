package main

import (
	"net/http"
	"time"
)


func main() {
	var c = http.Cookie{
		Name:     "test3",
		Value:    "value3",
		Expires:  time.Now().Add(111 * time.Second),
		HttpOnly: false,
	}
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name: "test",
			Value: "value1",
			Expires: time.Now().Add(111 * time.Second),
			HttpOnly: true,
		})
		http.SetCookie(w, &http.Cookie{
			Name: "test2",
			Value: "value2",
			Expires: time.Now().Add(222 * time.Second),
			HttpOnly: false,
		})
		http.SetCookie(w, &http.Cookie{
			Name: "test4",
			Value: "value4",
			Expires: time.Now().Add(222 * time.Second),
		})
		http.SetCookie(w, &c)
		w.Write([]byte("111111"))
	})
	http.ListenAndServe("127.0.0.1:60000", nil)
}