package json

import (
	"encoding/json"
	"log"
)

type Person struct {
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
	p, _ := UnmarshallPerson(myJson)

	log.Printf("%v", p)

	result, _ := MarshallPerson(*p)
	log.Printf("%v", string(result))
}

func MarshallPerson(p Person) (string, error) {
	b, err := json.Marshal(p)
	return string(b), err
}

func UnmarshallPerson(personJson string) (*Person, error) {
	var unmarshalled Person
	err := json.Unmarshal([]byte(personJson), &unmarshalled)
	if err != nil {
		log.Println("Error while unmarshalling", err)
		return nil, err
	}
	return &unmarshalled, nil
}
