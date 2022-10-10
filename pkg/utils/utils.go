package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const randNumberUrl = "https://codechallenge.boohma.com/random"

type RandNumber struct {
	RandomNumber int `json:"random_number"`
}

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func GetRandomNumber() RandNumber {
	response, err := http.Get(randNumberUrl)

	if err != nil {
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject RandNumber
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(responseObject)
	fmt.Println(responseObject.RandomNumber)
	return responseObject
}
