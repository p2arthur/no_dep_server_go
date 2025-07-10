package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/calculate", handleAdjustedValue)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", mux)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Welcome to the Investment Calculator!")
}

// /adjusted?principal=1000&rate=5.5&years=10
func handleAdjustedValue(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	principal, _ := strconv.ParseFloat(q.Get("principal"), 64)
	rate, _ := strconv.ParseFloat(q.Get("rate"), 64)
	years, _ := strconv.ParseFloat(q.Get("years"), 64)

	fmt.Println("Adjusted Value Calculation", principal, rate, years)

	expectedReturnRate := 5.5

	value := calculateInvestment(principal, rate, years, expectedReturnRate)
	fmt.Fprintf(w, "Adjusted Future Value: %.2f\n", value)
}

func calculateInvestment(principal float64, rate float64, years float64, expectedReturnRate float64) float64 {

	futureValue := principal * math.Pow(float64(1+expectedReturnRate/100), float64(years))
	futureAdjustedValue := futureValue / math.Pow(1+rate/100, years)

	fmt.Println("Investment Calculator", futureValue, futureAdjustedValue)
	return futureAdjustedValue
}
