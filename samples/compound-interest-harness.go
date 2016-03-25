package main

import (
	"fmt"
	"math" // power function
	"net/http"
	"net/url" // parse a simple request
	"strconv" // convert a string to float
)

func compute_compound_interest(principal float64, rate float64, t_years float64, n_yearly_compounded float64) float64 {
	return //p * math.Pow(1+r/n),(n*t))
}

func handler(w http.ResponseWriter, r *http.Request) {
	values, _ := url.ParseQuery(r.URL.Path[1:])
	// Btw, How many elements must be in this map?
	// x,_ := strconv.ParseFloat(values["x"][0], 64)
	//fmt.Fprintf(w, "%.02f", a)
}

func main() {
	fmt.Println("Waiting...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9090", nil)
}
