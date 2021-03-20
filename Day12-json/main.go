package main

import (
	"encoding/json"
	"log"
)

type person struct {
	// To make the field "marshallable", the field has to be exported.
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	myJson := `
		{
			"first_name": "anran", 
			"last_name": "niu"
		}
	`
	var unmarshelled person
	// If person has no exported fields, there will be a warning here saying
	// "struct doesn't have any exported fields, nor custom marshaling"
	// and the unmarshalled object will be empty.
	err := json.Unmarshal([]byte(myJson), &unmarshelled)
	if err != nil {
		log.Println("Error while unmarshalling", err)
		return
	}

	log.Printf("%v", unmarshelled)

	result, _ := json.Marshal(unmarshelled)
	log.Printf("%v", string(result))
}
