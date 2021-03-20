package json

import "testing"

func TestMarshallPerson(t *testing.T) {
	_, err := MarshallPerson(Person{"yihao", "Wang"})
	if err != nil {
		t.Error("Got an error when marshalling person")
	}
}
