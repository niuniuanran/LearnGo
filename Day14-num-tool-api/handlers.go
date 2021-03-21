package main

import (
	"encoding/json"
	"fmt"
	"github.com/niuniuanran/LearnGo/Day14-num-tool-api/lib"
	"net/http"
	"strconv"
)

type numberResponse struct {
	Number  int   `json:"number"`
	Factors []int `json:"factors"`
	IsPrime bool  `json:"is-prime"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintln(w, "Invalid request")
		return
	}
	tool := r.FormValue("tool")
	num, err := strconv.ParseInt(r.FormValue("number"), 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintln(w, "Invalid number")
		return
	}
	number := int(num)
	switch tool {
	case "prime", "factors":
		isPrime := lib.IsPrime(number)
		response := numberResponse{number, lib.Factors(number), isPrime}
		result, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprintf(w, "Internal service error: failed to marshall response")
			return
		}
		_, _ = fmt.Fprintln(w, string(result))
	case "help":
		fallthrough
	default:
		_, _ = fmt.Fprintln(w, "number\n\tSpecify the number you are checking. Default: 0\n"+
			"tool\n\tSpecify the tool you want to use.\n"+
			"\t\tAvailable tools: \n"+
			"\t\tprime\n\t\t\t checks if the number is prime\n"+
			"\t\tfactors\n\t\t\t breaks the number into prime factors\n"+
			"\t\thelp\n\t\t\t see help instructions")
	}
}
