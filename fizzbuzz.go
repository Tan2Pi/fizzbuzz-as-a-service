package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"encoding/json"
)

type FizzBuzz struct {
	Start int `json:"start"`
	End int `json:"end"`
	FizzNum int `json:"fizzNum"`
	BuzzNum int `json:"buzzNum"`
	Results []string `json:"results"`
}

func fizzBuzz(w http.ResponseWriter, r *http.Request) {
	log.Println("here")
	var f FizzBuzz
	err := json.NewDecoder(r.Body).Decode(&f)
	fmt.Println(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	for i := f.Start; i <= f.End; i++ {
		if i % f.FizzNum == 0 && i % f.BuzzNum == 0{
			f.Results = append(f.Results, "FizzBuzz")
		} else if i % f.FizzNum == 0 {
			f.Results = append(f.Results, "Fizz")
		} else if i % f.BuzzNum == 0 {
			f.Results = append(f.Results, "Buzz")
		} else {
			f.Results = append(f.Results, strconv.Itoa(i))
		}
	}
	fmt.Println(f)
	json.NewEncoder(w).Encode(f)
}

func fizzBuzzAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received")
	fizzBuzzRequest := FizzBuzz{
		Start: 3,
		End: 100,
		FizzNum: 3,
		BuzzNum: 5,
		Results: make([]string, 0),
	}
	request, err := json.Marshal(fizzBuzzRequest)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post("http://localhost:10000/fizzbuzz", "application/json", bytes.NewBuffer(request))
	if err != nil {
		fmt.Println(err.Error())
		fmt.Fprintln(w, err.Error())
	}
	defer resp.Body.Close()

	var fizzBuzzResponse FizzBuzz
	json.NewDecoder(resp.Body).Decode(&fizzBuzzResponse)

	fmt.Println(fizzBuzzResponse)
	for _, r := range fizzBuzzResponse.Results {
		fmt.Fprintln(w, r)
	}
}

func handleRequests() {
	http.HandleFunc("/api", fizzBuzzAPI)
	http.HandleFunc("/fizzbuzz", fizzBuzz)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}