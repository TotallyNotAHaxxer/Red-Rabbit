package main

import (
  "fmt"
  "net/http"
)

func main() {
  fs := http.FileServer(http.Dir("./static"))
  http.Handle("/", fs)

  fmt.Printf("Listening on localhost:80...")
	
  err := http.ListenAndServe("localhost:80", nil)

  if err != nil {
    panic(err.Error())
  }
}