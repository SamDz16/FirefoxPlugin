package RelaxBuisness

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Struct fields must be capitalized, if not they cannot be exported
type Query struct {
	Query   string
	Parents []string
}

type Sparql struct {
	Results Results `json:"results"`
}

type Results struct {
	Bindings []Bindings `json:"bindings"`
}

type Bindings struct{}

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
	r.Header.Add("Accept", "application/sparql-results+json")

	res, err := client.Do(r)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return []byte{}
	}

	// Read the response body
	dataBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return dataBody
}

func GetQueryTripplePatterns(initialQuery Query) []string {
	// fmt.Println("Executing GetQueryTripplePatterns() function ...")

	triplePatternsStr := ""
	start := false
	for _, ch := range initialQuery.Query {

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
			triplePatternsStr += char
		}
	}

	tps := strings.Split(triplePatternsStr, " . ")
	var resultTPs []string
	for _, tp := range tps {
		resultTPs = append(resultTPs, strings.TrimSpace(tp))
	}

	return resultTPs
}

func GenerateLevelTripplePatterns(triplePatterns []string, level int) []string {
	// fmt.Println("Executing GenerateLevelTripplePatterns() function ...")

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

func MakeQueries(tripplePatterns []string, queries *[]Query, K int) {
	// fmt.Println("Executing MakeQueries() function ...")
	K++
	k := strconv.Itoa(K)

	for _, t := range tripplePatterns {
		var q Query
		q.Query = "select * where { " + strings.TrimSpace(t) + " } limit " + k
		*queries = append(*queries, q)
	}
}

func MakeLattice(initialQuery Query, queries *[]Query, K int) {
	// fmt.Println("Executing MakeLattice() function ...")

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

	MakeQueries(allTripplePatterns, queries, K)
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
	// fmt.Println("Executing Base() function ...")

	qs := *queries

	for i := 0; i < len(*queries); i++ {
		if i == 0 {
			// The first element has no parents - root
			qs[i].Parents = []string{}
		} else if i == len(qs)-1 {
			// The last element is the empty query
			for _, qTemp := range GetQueryTripplePatterns(qs[0]) {
				qs[i].Parents = append(qs[i].Parents, qTemp)
			}
		} else {
			//  The rest of the elements : {1, 2, ..., len(queries)-2}
			for j := 0; j < i; j++ {
				res := IsDirectParent(qs[j], qs[i])
				if res {
					qs[i].Parents = append(qs[i].Parents, strings.Join(GetQueryTripplePatterns(qs[j])[:], " . "))
				}
			}
		}
	}
}

func ContainsKey(queries *map[*Query]bool, q Query) bool {
	// fmt.Println("Executing ContainsKey() function ...")

	for k := range *queries {
		qTemp := *k
		if qTemp.Query == q.Query {
			return true
		}
	}
	return false
}

func FindQuery(queries []Query, query Query) (int, bool) {
	// fmt.Println("Executing FindQuery() function ...")

	for i, q := range queries {
		if q.Query == query.Query {
			return i, true
		}
	}
	return -1, false
}

func RemoveQuery(listQueries []Query, index int) []Query {
	// fmt.Println("Executing RemoveQuery() function ...")

	var newListQueries []Query
	for j, q := range listQueries {
		if j != index {
			newListQueries = append(newListQueries, q)
		}
	}

	return newListQueries
}

func TpExecuteSPARQLQuery(requestUrl string, sparqlQuery string) int {
	// fmt.Println("Executing TpExecuteSPARQLQuery() function ...")

	// defer wg.Done()

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
		return -1
	}

	// Read the response body
	dataBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var results Sparql
	err = json.Unmarshal([]byte(dataBody), &results)

	if err != nil {
		return -1
	}

	return len(results.Results.Bindings)
}

func Base(q string, K int, endpoint string) ([]string, []string, int, string, string) {
	// fmt.Println("Executing Base() function ...")
	// Initialisations
	// ##################################################################################################################################################################### //
	// ##########################################################           INITIALIZE ALGO         ######################################################################## //
	// ##################################################################################################################################################################### //
	var initialQuery Query
	initialQuery.Query = q
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
	start := time.Now()
	MakeLattice(initialQuery, &listQueries, K)
	makeLatticeTime := time.Since(start)

	SetSuperQueries(&listQueries)

	start1 := time.Now()
	for len(listQueries) != 0 {

		// First element of the list
		qTemp := listQueries[0]

		// Remove the first element from the list
		listQueries = listQueries[1:]

		var Nb int

		if qTemp.Query != "select * where {  } limit "+strconv.Itoa(K+1) {

			dataBody := ExecuteSPARQLQuery(endpoint, qTemp.Query)
			var s Sparql
			json.Unmarshal([]byte(dataBody), &s)

			Nb = len(s.Results.Bindings)
			executedQueries[&qTemp] = Nb
		} else {
			Nb = 1
		}

		// Get Direct Super Queries Of 'qTemp'
		var superQueries []Query
		MakeQueries(qTemp.Parents, &superQueries, K)

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
			if parentsFIS && qTemp.Query != " " {
				*listXSS = append(*listXSS, qTemp)
			}
		}

	}
	executingQueriesTime := time.Since(start1)
	var xss []string
	for _, x := range *listXSS {
		xss = append(xss, x.Query)
	}
	var mfis []string
	for _, m := range *listMFIS {
		mfis = append(mfis, m.Query)
	}
	return xss, mfis, len(executedQueries), makeLatticeTime.String(), executingQueriesTime.String()
}

func BFS(q string, K int, endpoint string) ([]string, []string, int, string, string) {
	// fmt.Println("Executing Base() function ...")
	// Initialisations
	// ##################################################################################################################################################################### //
	// ##########################################################           INITIALIZE ALGO         ######################################################################## //
	// ##################################################################################################################################################################### //
	var initialQuery Query
	initialQuery.Query = q
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
	start := time.Now()
	MakeLattice(initialQuery, &listQueries, K)
	makeLatticeTime := time.Since(start)

	SetSuperQueries(&listQueries)

	start1 := time.Now()
	for len(listQueries) != 0 {

		// First element of the list
		qTemp := listQueries[0]

		// Remove the first element from the list
		listQueries = listQueries[1:]

		// NO QUERY EXECUTED HERE

		// Get Direct Super Queries Of 'qTemp'
		var superQueries []Query
		MakeQueries(qTemp.Parents, &superQueries, K)

		parentsFIS := true

		i := 0

		for parentsFIS && i < len(superQueries) {
			superQuery := superQueries[i]
			if !ContainsKey(&listFIS, superQuery) {
				parentsFIS = false
			}
			i++
		}
		if parentsFIS {
			// We execute query
			var Nb int

			if qTemp.Query != "select * where {  } limit "+strconv.Itoa(K+1) {

				dataBody := ExecuteSPARQLQuery(endpoint, qTemp.Query)
				var s Sparql
				json.Unmarshal([]byte(dataBody), &s)

				Nb = len(s.Results.Bindings)
				executedQueries[&qTemp] = Nb
			} else {
				Nb = 1
			}

			if Nb > K {
				// Query qTemp fails

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

			} else {
				// qTemp has succeded
				if qTemp.Query != " " {
					*listXSS = append(*listXSS, qTemp)
				}
			}
		}

	}
	executingQueriesTime := time.Since(start1)
	var xss []string
	for _, x := range *listXSS {
		xss = append(xss, x.Query)
	}
	var mfis []string
	for _, m := range *listMFIS {
		mfis = append(mfis, m.Query)
	}
	return xss, mfis, len(executedQueries), makeLatticeTime.String(), executingQueriesTime.String()
}

func GetQueryVariables(query Query) []string {
	start := false
	tp := ""
	existed := false
	var variables []string

	for _, ch := range query.Query {
		char := string(ch)

		if char == "?" {
			start = true
			continue
		} else if char == " " {
			if start {
				for _, variable := range variables {
					if variable == tp {
						existed = true
						break
					}
				}
				if !existed {
					variables = append(variables, tp)
				}
				existed = false
				tp = ""
			}
			start = false
			continue
		}

		if start {
			tp += char
		}
	}

	return variables
}

func RemoveQueryTriplePattern(query Query, tp string, K int) Query {
	queryTPs := GetQueryTripplePatterns(query)
	var newQueryTPs []string

	for _, queryTP := range queryTPs {
		if strings.TrimSpace(queryTP) != strings.TrimSpace(tp) {
			newQueryTPs = append(newQueryTPs, strings.TrimSpace(queryTP))
		}
	}

	// Make a new query with these triple patterns
	var q Query
	q.Query = "select * where { "
	for i, tp := range newQueryTPs {
		if i == len(newQueryTPs)-1 {
			q.Query += strings.TrimSpace(tp)
		} else {
			q.Query += strings.TrimSpace(tp) + " . "
		}
	}
	q.Query += " } limit " + strconv.Itoa(K+1)

	return q
}

func ExistQuery(queries []Query, query Query) bool {
	for _, q := range queries {
		if q.Query == query.Query {
			return true
		}
	}
	return false
}

func Var(q string, K int, endpoint string) ([]string, []string, int, string, string) {
	// fmt.Println("Executing Base() function ...")
	// Initialisations
	// ##################################################################################################################################################################### //
	// ##########################################################           INITIALIZE ALGO         ######################################################################## //
	// ##################################################################################################################################################################### //
	var initialQuery Query
	initialQuery.Query = q
	// List Queries
	var listQueries []Query
	// Executed Queries : contains for each qury, the number of the results
	var executedQueries map[*Query]int = make(map[*Query]int)
	// List FIS : all rthe queries that fail
	var listFIS map[*Query]bool = make(map[*Query]bool)

	listXSS := &[]Query{}
	listMFIS := &[]Query{}

	var notToExecuteQueries []Query
	// ##################################################################################################################################################################### //
	// ##########################################################              RUN ALGO             ######################################################################## //
	// ##################################################################################################################################################################### //
	start := time.Now()
	MakeLattice(initialQuery, &listQueries, K)
	makeLatticeTime := time.Since(start)

	SetSuperQueries(&listQueries)

	start1 := time.Now()
	for len(listQueries) != 0 {
		// First element of the list
		qTemp := listQueries[0]

		// Remove the first element from the list
		listQueries = listQueries[1:]

		// NO QUERY EXECUTED HERE

		// Get Direct Super Queries Of 'qTemp'
		var superQueries []Query
		MakeQueries(qTemp.Parents, &superQueries, K)

		parentsFIS := true

		i := 0
		for parentsFIS && i < len(superQueries) {
			superQuery := superQueries[i]
			if !ContainsKey(&listFIS, superQuery) {
				parentsFIS = false
			}
			i++
		}

		if parentsFIS {
			if !ExistQuery(notToExecuteQueries, qTemp) {

				// We execute query
				var Nb int

				if qTemp.Query != "select * where {  } limit "+strconv.Itoa(K+1) {

					dataBody := ExecuteSPARQLQuery(endpoint, qTemp.Query)
					var s Sparql
					json.Unmarshal([]byte(dataBody), &s)

					Nb = len(s.Results.Bindings)
					executedQueries[&qTemp] = Nb
				} else {
					Nb = 1
				}

				if Nb > K {
					// Query qTemp fails : FIS

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

					// Var property
					// Get qTemp Triple Patterns
					qTempTPs := GetQueryTripplePatterns(qTemp)
					for _, tp := range qTempTPs {
						// subQ is {qtemp - tp}
						subQ := RemoveQueryTriplePattern(qTemp, tp, K)
						if len(GetQueryVariables(subQ)) == len(GetQueryVariables(qTemp)) {
							// executedQueries[&subQ] = K + 1
							notToExecuteQueries = append(notToExecuteQueries, subQ)
						}
					}
				} else {
					// qTemp has succeded
					if qTemp.Query != " " {
						*listXSS = append(*listXSS, qTemp)
					}
				}
			}
		}
	}
	executingQueriesTime := time.Since(start1)
	var xss []string
	for _, x := range *listXSS {
		xss = append(xss, x.Query)
	}
	var mfis []string
	for _, m := range *listMFIS {
		mfis = append(mfis, m.Query)
	}

	return xss, mfis, len(executedQueries), makeLatticeTime.String(), executingQueriesTime.String()
}

func ExtractTPProperties(tp string) (string, string, string) {
	tp = strings.TrimSpace(tp)
	props := strings.Split(tp, " ")
	return props[0], props[1], props[2]
}

func GetPredicates(q Query) []string {
	insert := true
	var predicates []string
	qTPs := GetQueryTripplePatterns(q)
	for _, tp := range qTPs {
		_, p, _ := ExtractTPProperties(tp)

		for _, pred := range predicates {
			if pred == p {
				insert = false
			}
		}

		if insert {
			predicates = append(predicates, p)
		}

		insert = true
	}

	return predicates
}

func ExistString(strs []string, str string) bool {
	str = strings.TrimSpace(str)
	for _, s := range strs {
		s = strings.TrimSpace(s)
		if s == str {
			return true
		}
	}

	return false
}

func Full(q string, K int, endpoint string, strCardsArray []string) ([]string, []string, int, string, string) {
	// fmt.Println("Executing Base() function ...")
	// Initialisations
	// ##################################################################################################################################################################### //
	// ##########################################################           INITIALIZE ALGO         ######################################################################## //
	// ##################################################################################################################################################################### //
	var initialQuery Query
	initialQuery.Query = q
	// List Queries
	var listQueries []Query
	// Executed Queries : contains for each qury, the number of the results
	var executedQueries map[*Query]int = make(map[*Query]int)
	// List FIS : all rthe queries that fail
	var listFIS map[*Query]bool = make(map[*Query]bool)

	listXSS := &[]Query{}
	listMFIS := &[]Query{}

	var notToExecuteQueries []Query

	var predicateCards map[string]string = make(map[string]string)
	predicates := GetPredicates(initialQuery)

	for i, p := range predicates {
		predicateCards[p] = strCardsArray[i]
	}

	// ##################################################################################################################################################################### //
	// ##########################################################              RUN ALGO             ######################################################################## //
	// ##################################################################################################################################################################### //
	start := time.Now()
	MakeLattice(initialQuery, &listQueries, K)
	makeLatticeTime := time.Since(start)

	SetSuperQueries(&listQueries)

	start1 := time.Now()
	for len(listQueries) != 0 {
		// First element of the list
		qTemp := listQueries[0]

		// Remove the first element from the list
		listQueries = listQueries[1:]

		// NO QUERY EXECUTED HERE

		// Get Direct Super Queries Of 'qTemp'
		var superQueries []Query
		MakeQueries(qTemp.Parents, &superQueries, K)

		parentsFIS := true

		i := 0
		for parentsFIS && i < len(superQueries) {
			superQuery := superQueries[i]
			if !ContainsKey(&listFIS, superQuery) {
				parentsFIS = false
			}
			i++
		}

		if parentsFIS {
			if !ExistQuery(notToExecuteQueries, qTemp) {

				// We execute query
				var Nb int

				if qTemp.Query != "select * where {  } limit "+strconv.Itoa(K+1) {

					dataBody := ExecuteSPARQLQuery(endpoint, qTemp.Query)
					var s Sparql
					json.Unmarshal([]byte(dataBody), &s)

					Nb = len(s.Results.Bindings)
					executedQueries[&qTemp] = Nb
				} else {
					Nb = 1
				}

				if Nb > K {
					// Query qTemp fails : FIS

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

					// Var property
					// Get qTemp Triple Patterns
					qTempTPs := GetQueryTripplePatterns(qTemp)
					for _, tp := range qTempTPs {
						// subQ is {qtemp - tp}
						subQ := RemoveQueryTriplePattern(qTemp, tp, K)
						if len(GetQueryVariables(subQ)) == len(GetQueryVariables(qTemp)) {
							// executedQueries[&subQ] = K + 1
							notToExecuteQueries = append(notToExecuteQueries, subQ)
						} else {
							// Full
							s, p, _ := ExtractTPProperties(tp)
							card, ok := predicateCards[p]
							if ok {
								fCard, _ := strconv.ParseFloat(card, 64)
								if fCard == 1 && ExistString(GetQueryVariables(subQ), s) {
									notToExecuteQueries = append(notToExecuteQueries, subQ)
								}
							}
						}
					}
				} else {
					// qTemp has succeded
					if qTemp.Query != " " {
						*listXSS = append(*listXSS, qTemp)
					}
				}
			}
		}
	}
	executingQueriesTime := time.Since(start1)
	var xss []string
	for _, x := range *listXSS {
		xss = append(xss, x.Query)
	}
	var mfis []string
	for _, m := range *listMFIS {
		mfis = append(mfis, m.Query)
	}

	return xss, mfis, len(executedQueries), makeLatticeTime.String(), executingQueriesTime.String()
}
