package main

import (
	"net/http"
	"fmt"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,"<h1>HelloWorld</h1>")
	})
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		t:=time.Now()
		timeStr:=fmt.Sprintf("{\"time\":\"%s\"}",t)
		w.Write([]byte(timeStr))
	})
	http.ListenAndServe(":8080",nil)
}
