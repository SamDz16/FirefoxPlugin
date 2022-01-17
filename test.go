package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	result := getData("https://dbpedia.org/sparql", "sel distinct ?Concept where {[] a ?Concept} LIMIT 100")
	println(result)
}

func getData(requestUrl string, sparqlQuery string) int {
	data := url.Values{}
	data.Set("query", sparqlQuery)

	u, _ := url.ParseRequestURI(requestUrl)
	urlStr := u.String()

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(r)
	if err != nil {
		// Handle errors: reject the Promise if we have an error
		println("Error")
	}
	defer res.Body.Close()

	// Read the response body
	dataBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// Handle errors here too
		println("ERROR")
	}
	return len(string(dataBody))
}
