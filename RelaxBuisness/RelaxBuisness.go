package RelaxBuisness

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// To be exported function must sytart with a capital letter
func ExecuteSPARQLQuery(requestUrl string, sparqlQuery string) []byte {
	// Make the HTTP request
	// res, err := http.DefaultClient.Get(requestUrl)
	data := url.Values{}
	data.Set("query", sparqlQuery)

	u, _ := url.ParseRequestURI(requestUrl)
	urlStr := u.String()

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(r)
	if err != nil {
		println("ERROR")
	}
	defer res.Body.Close()

	// Read the response body
	dataBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		println("ERROR")
	}

	return dataBody
}

func IsFailing(requestUrl string, sparqlQuery string) int {

	var dataBody []byte

	go func() {
		dataBody = ExecuteSPARQLQuery(requestUrl, sparqlQuery)
	}()

	if len(dataBody) == 0 {
		return 1
	}

	return 0
}
