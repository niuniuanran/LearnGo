package main

import (
	"fmt"
	hello "learngo/lib"
)

func main() {
	person := hello.Person{
		FirstName: "豪豪",
		LastName:  "王",
	}

	persons := [...]hello.Person{
		{},
		person,
		person,
	}

	var sayHelloToAll = func(people []hello.Person) {
		people[1].FirstName = "然然"
		for i := 0; i <= 2; i++ {
			fmt.Printf("Hello, %s %s\n", people[i].FirstName, people[i].LastName)
		}
	}

	var sayHelloToAll2 = func(people [3]hello.Person) {

		for _, p := range people {
			if p.FirstName != "" {
				fmt.Printf("Hello, %s %s\n", p.FirstName, p.LastName)
			}
		}
	}

	sayHelloToAll(persons[:])
	sayHelloToAll2(persons)

	fmt.Println("-------")

	fs()

	// for _, f := range actionGenerator() {
	// 	fmt.Println(f("Hi"))
	// }

}

// func actionGenerator() []func(string) string {
// 	persons := [...]hello.Person{
// 		{FirstName: "豪豪"},
// 		{FirstName: "牛牛"},
//    		{FirstName: "狗狗"},
// 	}
// 	functions := make([]func(string) string, 3, 3)
// 	for i, p := range persons {
// 		functions[i] = func(h string) string {
// 			return h + p.FirstName
// 		}
// 	}
// 	return functions
// }

func fs() {
	ps := []hello.Person{
		{FirstName: "豪豪"},
		{FirstName: "牛牛"},
		{FirstName: "狗狗"},
	}

	var res []func() string
	for _, p := range ps {
		func(p hello.Person) {
			res = append(res, func() string {
				return p.FirstName
			})
		}(p)
	}
	for _, f := range res {
		fmt.Println(f())
	}
}
