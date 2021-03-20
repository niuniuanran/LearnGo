package prime

import "testing"

var tests = []struct{
	name string
	number int
	expected bool
}{
	{"-1 should not be prime", -1, false},
	{"1 should not be prime", 1, false},
	{"2 should be prime", 2, true},
	{"3 should be prime", 3, true},
	{"2017 should be prime", 3, true},
}

func TestIsPrime(t *testing.T) {
	for _, tt := range tests {
		result := IsPrime(tt.number)
		if result != tt.expected {
			t.Error(tt.name, "- Did not get expected result")
		}
	}
}