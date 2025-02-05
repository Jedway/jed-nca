package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
)

type NumberResponse struct {
	Number     int      `json:"number"`
	IsPrime    bool     `json:"is_prime"`
	IsPerfect  bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum   int      `json:"digit_sum"`
	FunFact    string   `json:"fun_fact"`
}

type ErrorResponse struct {
	Number string `json:"number"`
	Error  bool   `json:"error"`
}

func main() {
	http.HandleFunc("/api/classify-number", classifyNumberHandler)
	http.ListenAndServe(":8080", nil)
}

func classifyNumberHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Handle CORS

	numberParam := r.URL.Query().Get("number")
	number, err := strconv.Atoi(numberParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Number: numberParam, Error: true})
		return
	}

	isPrime := checkPrime(number)
	isPerfect := checkPerfect(number)
	digitSum := calculateDigitSum(number)
	properties := determineProperties(number, isPrime)
	funFact := getFunFact(number)

	response := NumberResponse{
		Number:     number,
		IsPrime:    isPrime,
		IsPerfect:  isPerfect,
		Properties: properties,
		DigitSum:   digitSum,
		FunFact:    funFact,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func checkPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func checkPerfect(n int) bool {
	if n < 2 {
		return false
	}
	sum := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}
	return sum == n
}

func calculateDigitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func determineProperties(n int, isPrime bool) []string {
	properties := []string{}
	if isArmstrong(n) {
		properties = append(properties, "armstrong")
	}
	if n%2 == 0 {
		properties = append(properties, "even")
	} else {
		properties = append(properties, "odd")
	}
	return properties
}

func isArmstrong(n int) bool {
	originalNumber := n
	sum := 0
	digits := len(strconv.Itoa(n))
	for n > 0 {
		digit := n % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		n /= 10
	}
	return sum == originalNumber
}

func getFunFact(n int) string {
	url := fmt.Sprintf("http://numbersapi.com/%d/math", n)
	resp, err := http.Get(url)
	if err != nil {
		return "No fun fact available."
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "No fun fact available."
	}
	return string(body)
}
