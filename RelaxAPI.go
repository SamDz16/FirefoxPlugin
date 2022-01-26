package main

import (
	"buisness/RelaxBuisness"
	"syscall/js"
)

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

				// "data" is a byte slice, so we need to convert it to a JS Uint8Array object
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

func main() {
	c := make(chan int)

	js.Global().Set("executeSPARQLQuery", executeSPARQLQuery())
	js.Global().Set("isFailing", isFailing())

	<-c

}
