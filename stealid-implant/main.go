package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	postBody, _ := json.Marshal(map[string]string{
		"country_code": "in",
		"ip":           "1.2.3.5",
		"creds":        "sdfds",
		"user":         "6sba2ueggucg514",
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://localhost:8090/api/collections/bots/records", "application/json", responseBody)
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
