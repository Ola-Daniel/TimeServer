//a simple http server the returns the current day and time in JSON format in GMT+1

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Time struct {
	Time string
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	t := Time{time.Now().Format(time.RFC1123)}
	b, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, string(b))
}	

