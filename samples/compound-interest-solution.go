package main

import (
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

func compute_compound_interest(principal float64, rate float64, t_years float64, n_yearly_compounded float64) float64 {
	return principal * math.Pow((1+rate/n_yearly_compounded), (n_yearly_compounded*t_years))
}

func handler(w http.ResponseWriter, r *http.Request) {
	values, _ := url.ParseQuery(r.URL.Path[1:])
	if len(values) < 4 {
		return
	}
	p, _ := strconv.ParseFloat(values["p"][0], 64)
	rate, _ := strconv.ParseFloat(values["r"][0], 64)
	t, _ := strconv.ParseFloat(values["t"][0], 64)
	n, _ := strconv.ParseFloat(values["n"][0], 64)
	a := compute_compound_interest(p, rate, t, n)
	fmt.Fprintf(w, "%.02f", a)
}

func main() {
	fmt.Println("Waiting...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9090", nil)
}
