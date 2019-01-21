package main
import (
  "log"
  "fmt"
  "net/http"
)
func hello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Hello World")
}

func main() {
  
  http.HandleFunc("/", hello)
  log.Printf("Listening on %s...\n", ":5000")

  if err := http.ListenAndServe(":5000", nil); err != nil {
    panic(err)
  }
}