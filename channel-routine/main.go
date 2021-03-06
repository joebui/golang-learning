package main

import (
  "fmt"
  "net/http"
  "time"
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

  // Run infinitely
  for l := range channel {
    // Listen for data and execute.
    go (func(link string) {
      time.Sleep(5 * time.Second)
      checkLink(link, channel)
    })(l)
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
