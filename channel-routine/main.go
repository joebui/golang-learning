package main

import (
  "fmt"
  "net/http"
)

func main() {
  links := []string{
    "http://google.com",
    "http://github.com",
    "http://facebook.com",
  }

  channel := make(chan string)

  for _, link := range links {
    go checkLink(link, channel)
  }

  for l := range channel {
    // Listen for data.
    // fmt.Println(<-channel)
    go checkLink(l, channel)
  }
}

func checkLink(l string, c chan string) {
  _, err := http.Get(l)
  if err != nil {
    fmt.Println("Error")
    c <- l
    return
  }

  fmt.Println(l, "works")
  c <- l
}
