package main

import (
	"Relaxbuisness/RelaxBuisness"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"syscall/js"
)

type Query struct {
	query   string
	parents []string
}

type Sparql struct {
	Results *Results `xml:"results"`
}
type Results struct {
	Result []Result `xml:"result"`
}

type Result struct{}

func GetQueryTripplePatterns(initialQuery Query) []string {
	triplePatternsStr := ""
	start := false
	for _, ch := range initialQuery.query {

		char := string(ch)

		if char == "{" || char == "}" {
			if char == "{" {
				start = true
				continue
			} else {
				start = false
				continue
			}
		}

		if start {
			// if char == " " {
			// 	continue
			// }
			triplePatternsStr += char
		}
	}

	return strings.Split(triplePatternsStr, " . ")
}

func GenerateLevelTripplePatterns(triplePatterns []string, level int) []string {
	combinations := []string{}
	n := len(triplePatterns)

	indexes := []int{}
	for i := 0; i < level; i++ {
		indexes = append(indexes, i)
	}

	//  liste_combinaisons.append(tuple([e[index] for index in indices]))
	triple := ""
	for i, index := range indexes {
		if i == len(indexes)-1 {
			triple += triplePatterns[index]
		} else {
			triple += triplePatterns[index] + " . "
		}
	}
	combinations = append(combinations, triple)

	if level == 0 {
		return []string{" "}
	}

	if level == n {
		return combinations
	}

	i := level - 1

	for i != -1 {
		indexes[i] += 1

		for j := i + 1; j < level; j++ {
			indexes[j] = indexes[j-1] + 1
		}

		if indexes[i] == (n - level + i) {
			i -= 1
		} else {
			i = level - 1
		}

		temp := []string{}
		for _, index := range indexes {
			temp = append(temp, triplePatterns[index])
		}

		// fmt.Println("temp", temp)

		triple := ""
		for i, t := range temp {
			if i == len(temp)-1 {
				triple += t
			} else {
				triple += t + " . "
			}
		}

		combinations = append(combinations, triple)
	}
	return combinations
}

func MakeQueries(tripplePatterns []string, queries *[]Query) {
	for _, t := range tripplePatterns {
		var q Query
		q.query = "select * where {" + t + "}"
		*queries = append(*queries, q)
	}
}

func MakeLattice(initialQuery Query, queries *[]Query) {

	// Get all the tripple patterns
	tripplePatterns := GetQueryTripplePatterns(initialQuery)

	// triplePatternsNbr will contain the length of all the individual triple patterns of the initialQuery
	triplePatternsNbr := len(tripplePatterns)

	level := triplePatternsNbr
	var allTripplePatterns []string

	for i := 0; i < triplePatternsNbr+1; i++ {
		var temp []string = GenerateLevelTripplePatterns(tripplePatterns, level)
		for _, t := range temp {
			// q.parents = []string{}

			allTripplePatterns = append(allTripplePatterns, t)
		}
		level--
	}

	MakeQueries(allTripplePatterns, queries)
}

func IsDirectParent(q1, q2 Query) bool {
	tp1 := GetQueryTripplePatterns(q1)
	tp2 := GetQueryTripplePatterns(q2)

	if len(tp1) == len(tp2)+1 {
		matches := 0

		for _, q2 := range tp2 {
			for _, q1 := range tp1 {
				if q2 == q1 {
					matches++
					break
				}
			}
		}
		if matches == len(tp2) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func SetSuperQueries(queries *[]Query) {
	qs := *queries

	for i := 0; i < len(*queries); i++ {
		if i == 0 {
			// The first element has no parents - root
			qs[i].parents = []string{}
		} else if i == len(qs)-1 {
			// The last element is the empty query
			for _, qTemp := range GetQueryTripplePatterns(qs[0]) {
				qs[i].parents = append(qs[i].parents, qTemp)
			}
		} else {
			//  The rest of the elements : {1, 2, ..., len(queries)-2}
			for j := 0; j < i; j++ {
				res := IsDirectParent(qs[j], qs[i])
				if res {
					qs[i].parents = append(qs[i].parents, strings.Join(GetQueryTripplePatterns(qs[j])[:], " . "))

				}
			}
		}
	}
}

func ContainsKey(queries *map[*Query]bool, q Query) bool {

	for k := range *queries {
		qTemp := *k
		if qTemp.query == q.query {
			return true
		}
	}
	return false
}

func TpExecuteSPARQLQuery(requestUrl string, sparqlQuery string) int {
	// Make the HTTP request
	data := url.Values{}
	data.Set("query", sparqlQuery)

	u, _ := url.ParseRequestURI(requestUrl)
	urlStr := u.String()

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(r)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return 0
	}

	// Read the response body
	dataBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var XMLResponseData Sparql
	xml.Unmarshal([]byte(dataBody), &XMLResponseData)

	return len(*&XMLResponseData.Results.Result)
}

func FindQuery(queries []Query, query Query) (int, bool) {
	for i, q := range queries {
		if q.query == query.query {
			return i, true
		}
	}
	return -1, false
}

func RemoveQuery(listQueries []Query, index int) []Query {
	var newListQueries []Query
	for j, q := range listQueries {
		if j != index {
			newListQueries = append(newListQueries, q)
		}
	}

	return newListQueries
}

// Copyright (C) 2020 Alessandro Segala (ItalyPaleAle)
// License: MIT

// MyGoFunc fetches an external resource by making a HTTP request from Go
// The JavaScript method accepts one argument, which is the URL to request
func executeSPARQLQuery() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// Get the URL as argument
		// args[0] is a js.Value, so we need to get a string out of it
		requestUrl := args[0].String()
		sparqlQuery := args[1].String()

		// Handler for the Promise
		// We need to return a Promise because HTTP requests are blocking in Go
		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]
			// reject := args[1]

			// Run this code asynchronously
			go func() {

				dataBody := RelaxBuisness.ExecuteSPARQLQuery(requestUrl, sparqlQuery)

				// "dataBody" is a byte slice, so we need to convert it to a JS Uint8Array object
				arrayConstructor := js.Global().Get("Uint8Array")
				dataJS := arrayConstructor.New(len(dataBody))
				js.CopyBytesToJS(dataJS, dataBody)

				// Create a Response object and pass the data
				responseConstructor := js.Global().Get("Response")
				response := responseConstructor.New(dataJS)

				// Resolve the Promise
				resolve.Invoke(response)
			}()

			// The handler of a Promise doesn't return any value
			return nil
		})

		// Create and return the Promise object
		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}

func isFailing() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		requestUrl := args[0].String()
		sparqlQuery := args[1].String()

		return RelaxBuisness.IsFailing(requestUrl, sparqlQuery)
	})
}

func Base() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		// initialQuery := args[0].String()
		K := args[1].Int()

		// handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// 	resolve := args[0]

		// 	go func() {
		// 		nbrExecutedQueries := RelaxBuisness.Base(initialQuery, K)

		// 		// Resolve the Promise
		// 		resolve.Invoke(nbrExecutedQueries)

		// 	}()

		// 	return nil
		// })
		// promiseConstructor := js.Global().Get("Promise")
		// return promiseConstructor.New(handler)
		// Initialisations
		// ##################################################################################################################################################################### //
		// ##########################################################           INITIALIZE ALGO         ######################################################################## //
		// ##################################################################################################################################################################### //

		var initialQuery Query
		initialQuery.query = args[0].String()

		// List Queries
		var listQueries []Query

		// Executed Queries : contains for each qury, the number of the results
		var executedQueries map[*Query]int = make(map[*Query]int)

		// List FIS : all rthe queries that fail
		var listFIS map[*Query]bool = make(map[*Query]bool)

		listXSS := &[]Query{}
		listMFIS := &[]Query{}

		// ##################################################################################################################################################################### //
		// ##########################################################              RUN ALGO             ######################################################################## //
		// ##################################################################################################################################################################### //

		MakeLattice(initialQuery, &listQueries)

		SetSuperQueries(&listQueries)
		var s int = 0

		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]
			reject := args[1]

			go func() {

				for len(listQueries) != 0 {

					// First element of the list
					qTemp := listQueries[0]

					// Remove the first element from the list
					listQueries = listQueries[1:]

					var Nb int

					// go func() {
					// Make HTTP request and save the results of the request in Nb
					data := url.Values{}
					data.Set("query", qTemp.query)

					u, _ := url.ParseRequestURI("http://localhost:3030/base")
					urlStr := u.String()

					client := &http.Client{}
					r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
					r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

					res, err := client.Do(r)
					if err != nil {
						log.Fatalln(err)
					}
					defer res.Body.Close()

					if res.StatusCode != 200 {
						reject.Invoke(0)
					}

					// Read the response body
					dataBody, err := ioutil.ReadAll(res.Body)
					if err != nil {
						log.Fatalln(err)
					}

					var XMLResponseData Sparql
					xml.Unmarshal([]byte(dataBody), &XMLResponseData)

					Nb = len(*&XMLResponseData.Results.Result)

					// Nb = TpExecuteSPARQLQuery("http://localhost:3030/base", qTemp.query)

					// add qTemp to executedQueries list with the number Nb of the results
					executedQueries[&qTemp] = Nb
					// }()

					// Get Direct Super Queries Of 'qTemp'
					var superQueries []Query
					MakeQueries(qTemp.parents, &superQueries)

					for _, mfis := range superQueries {
						fmt.Println(s, " - superqueries: ", mfis)
						s++
					}

					fmt.Println("executedQueries: ")
					for k, v := range executedQueries {
						fmt.Println("query: ", (*k).query, " nbr: ", v)
					}

					parentsFIS := true

					i := 0

					for parentsFIS && i < len(superQueries) {
						superQuery := superQueries[i]
						if !ContainsKey(&listFIS, superQuery) {
							parentsFIS = false
						}
						i++
					}

					if Nb > K {
						// Query qTemp fails
						if parentsFIS {
							// We remove all the superqueries of qTemp from listMFIS list
							for _, qSQ := range superQueries {
								index, found := FindQuery(*listMFIS, qSQ)
								if found {
									*listMFIS = RemoveQuery(*listMFIS, index)
								}
							}

							// Since the request qTemp has failed, we add it to the list of FIS
							listFIS[&qTemp] = true

							// qTemps is the new MFIS
							*listMFIS = append(*listMFIS, qTemp)
						}
					} else {
						// qTemp has succeded
						if parentsFIS && qTemp.query != " " {
							*listXSS = append(*listXSS, qTemp)
						}
					}
				}

				fmt.Println("list XSS: ", *listXSS)
				fmt.Println("list MFIS: ", *listMFIS)

				resolve.Invoke(len(executedQueries))
			}()

			return nil
		})
		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}

func main() {
	c := make(chan int)

	js.Global().Set("executeSPARQLQuery", executeSPARQLQuery())
	js.Global().Set("isFailing", isFailing())
	js.Global().Set("Base", Base())

	<-c
}
