package main

import "fmt"

type contactInfo struct {
  email   string
  zipCode int
}

type person struct {
  firstName   string
  lastname    string
  contactInfo contactInfo
}

func main() {
  alex := person{firstName: "dien", lastname: "bui"}
  alex.print()

  fmt.Println("---")

  var alex2 person
  alex2.firstName = "Joe"
  alex2.lastname = "Bui"
  alex2.print()

  fmt.Println("---")

  jim := person{
    firstName: "Jim",
    lastname:  "Phil",
    contactInfo: contactInfo{
      zipCode: 12312,
      email:   "aaa@ddd.com",
    },
  }
  jim.print()
  jim.updateFirstName("Jimmy")
  jim.print()
}

// *pointer
func (p *person) updateFirstName(name string) {
  (*p).firstName = name
}

func (p person) print() {
  fmt.Printf("%+v\n", p)
}
