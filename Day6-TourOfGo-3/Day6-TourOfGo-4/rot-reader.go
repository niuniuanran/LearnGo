package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr rot13Reader) Read(bytes []byte) (int, error) {
	temp := make([]byte, len(bytes), cap(bytes))
	num, err := rr.r.Read(temp)
	if err != nil {
		return 0, err
	}

	for i, c := range temp {
		if c <= 'Z' && c >= 'A' {
			bytes[i] = (c-'A'+13)%26 + 'A'
		} else if c <= 'z' && c >= 'a' {
			bytes[i] = (c-'a'+13)%26 + 'a'
		} else {
			bytes[i] = c
		}
	}
	return num, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
