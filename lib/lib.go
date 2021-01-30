package hello

import (
	"fmt"
)

var defaultName string
var defaultLastName = "牛"
var (
	defaultFirstName  string
	defaultMiddleName       = "豪豪"
	b                 int64 = 1234567823425434134
)

const (
	faviourite = "豪豪"
	a          = 12345678912345678912345689
)

// Person is a person
type Person struct {
	FirstName string
	LastName  string
	age       uint16
}

// Age returns the age of a person
func Age(p Person) int {
	return int(p.age)
}

// Age returns the age of a person
func (p Person) Age() int {
	return int(p.age)
}

// DefaultName returns the defaultName
func DefaultName() string {
	return defaultName
}

// SetDefaultName sets the defaultName
func SetDefaultName(s string) {
	defaultName = s
}

// World says hello to s
func World(s string) {
	fmt.Printf("hello, %s\n", s)
}

// Girl says hello to s
func Girl(s string) {
	fmt.Printf("hello, pretty %s\n", s)
}
