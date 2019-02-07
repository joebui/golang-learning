package main

import (
  "fmt"
  "io/ioutil"
  "math/rand"
  "os"
  "strings"
  "time"
)

// Desk type and functions
type desk []string

func (d desk) print() {
  for _, card := range d {
    fmt.Println(card)
  }
}

func (d desk) toString() string {
  return strings.Join([]string(d), ",")
}

func (d desk) saveToFile(filename string) error {
  return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Desk type and functions

func newDesk() desk {
  cards := desk{}
  cardSuits := []string{"H", "D", "S", "C"}
  cardValues := []string{"ace", "1", "2", "3"}

  for _, suit := range cardSuits {
    for _, value := range cardValues {
      cards = append(cards, value+" of "+suit)
    }
  }

  return cards
}

func deal(d desk, handSize int) (desk, desk) {
  return d[:handSize], d[handSize:]
}

func newDeskFromFile(filename string) desk {
  bs, err := ioutil.ReadFile(filename)

  if err != nil {
    fmt.Println("Error: ", err)
    os.Exit(1)
  }

  s := strings.Split(string(bs), ",")

  return desk(s)
}

func (d desk) shuffle() {
  source := rand.NewSource(time.Now().UnixNano())
  r := rand.New(source)

  for i := range d {
    newPosition := r.Intn(len(d) - 1)
    d[i], d[newPosition] = d[newPosition], d[i]
  }
}
