package main

import (
	"Relaxbuisness/RelaxBuisness"
	"strings"
	"unsafe"

	"syscall/js"
)

type Query struct {
	query   string
	parents []string
}

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

func MakeQueries(tripplePatterns []string) []string {
	var res []string

	for _, t := range tripplePatterns {
		res = append(res, "select * where {"+t+"}")
	}
	return res
}

func MakeLattice(q string) []interface{} {

	var initialQuery Query
	initialQuery.query = q

	// Get all the tripple patterns
	tripplePatterns := GetQueryTripplePatterns(initialQuery)

	// triplePatternsNbr will contain the length of all the individual triple patterns of the initialQuery
	triplePatternsNbr := len(tripplePatterns)

	level := triplePatternsNbr
	var allTripplePatterns []string

	for i := 0; i < triplePatternsNbr+1; i++ {
		var temp []string = GenerateLevelTripplePatterns(tripplePatterns, level)
		for _, t := range temp {
			allTripplePatterns = append(allTripplePatterns, t)
		}
		level--
	}

	res := MakeQueries(allTripplePatterns)

	var ret []interface{}

	for _, q := range res {
		ret = append(ret, q)
	}

	return ret
}

func AllQueries() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		q := args[0].String()
		return MakeLattice(q)
	})
}

func IntToByte(num int) byte {
	size := int(unsafe.Sizeof(num))
	arr := make([]byte, size)
	for i := 0; i < size; i++ {
		byt := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&num)) + uintptr(i)))
		arr[i] = byt
	}
	return arr[0]
}

func Base() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		// Parameters of the base algorithm

		//1.  initial query : string
		initialQuery := args[0].String()

		//2.  The constant K : integer
		K := args[1].Int()
		// convert it to byte
		k := IntToByte(K)

		//3. NBS: convert JS array to Go slice in NBs
		NBs := make([]byte, args[2].Get("length").Int())
		_ = js.CopyBytesToGo(NBs, args[2])

		// Call the Base Algorithm
		xss, mfis, nbr := RelaxBuisness.Base(initialQuery, k, NBs)

		// List of xss
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
		return resGlobal
	})
}

func TpExecuteSPARQLQuery() js.Func {
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
				nb := RelaxBuisness.TpExecuteSPARQLQuery(requestUrl, sparqlQuery)

				// Resolve the Promise
				resolve.Invoke(nb)
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
	js.Global().Set("AllQueries", AllQueries())
	js.Global().Set("Base", Base())
	js.Global().Set("TpExecuteSPARQLQuery", TpExecuteSPARQLQuery())

	<-c
}
