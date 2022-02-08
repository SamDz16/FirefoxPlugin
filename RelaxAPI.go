package main

import (
	"Relaxbuisness/RelaxBuisness"

	"syscall/js"
)

func executeSPARQLQuery() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		requestUrl := args[0].String()
		sparqlQuery := args[1].String()

		// Handler for the Promise
		// We need to return a Promise because HTTP requests are blocking in Go
		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]

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

		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]

			go func() {
				// res := RelaxBuisness.IsFailing(requestUrl, sparqlQuery)
				var res int
				dataBody := RelaxBuisness.ExecuteSPARQLQuery(requestUrl, sparqlQuery)

				if len(dataBody) == 0 {
					res = 1
				} else {

					res = 0
				}

				// Retourne the result {0, 1}
				resolve.Invoke(res)
			}()

			return nil
		})

		// Create and return the Promise object
		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}

func Base() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// Parameters of the base algorithm
		//1.  initial query : string
		initialQuery := args[0].String()

		//2.  The constant K : integer
		K := args[1].Int()

		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]

			go func() {
				xss, mfis, nbr := RelaxBuisness.Base(initialQuery, K)

				var resXSS []interface{}
				for _, q := range xss {
					resXSS = append(resXSS, q)
				}

				// List of MFIS
				var resMFIS []interface{}
				for _, q := range mfis {
					resMFIS = append(resMFIS, q)
				}

				// Number of results
				var resGlobal []interface{}

				// Encapsulate everything into an object
				resGlobal = append(resGlobal, resXSS)
				resGlobal = append(resGlobal, resMFIS)
				resGlobal = append(resGlobal, nbr)

				// Return value of the Base Algorithm
				// return resGlobal
				resolve.Invoke(resGlobal)
			}()
			// The handler of a Promise doesn't return any value
			return nil
		})
		// Create and return the Promise object
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
