package RelaxBuisnessTest

import (
	. "Relaxbuisness/RelaxBuisness"
	"testing"
)

// func TestGetQueryTripplePatterns(t *testing.T) {
// 	// Given
// 	var q Query
// 	q.Query = "select * where { t1 . t2 . t3 }"

// 	// When
// 	result := GetQueryTripplePatterns(q) // results should be: [t1, t2, t3]

// 	// Then
// 	expected := []string{"t1", "t2", "t3"}
// 	for i, tp := range result {
// 		if tp != expected[i] {
// 			t.Errorf("GetQueryTripplePatterns FAILD. Expected %v found %v\n", expected, result)
// 			return
// 		}
// 	}
// 	t.Logf("GetQueryTripplePatterns PASSED. Expected %v found %v\n", expected, result)
// }

// func TestGenerateLevelTripplePatterns(t *testing.T) {
// 	// Given
// 	tripplePatterns := []string{"t1", "t2", "t3", "t4"}
// 	level := 2

// 	// When
// 	result := GenerateLevelTripplePatterns(tripplePatterns, level) // results should be: [t1.t2, t1.t3, t1.t4, t2.t3, t2.t4, t3.t4]

// 	// Then
// 	expected := []string{"t1 . t2", "t1 . t3", "t1 . t4", "t2 . t3", "t2 . t4", "t3 . t4"}
// 	for i, tp := range result {
// 		if tp != expected[i] {
// 			t.Errorf("GenerateLevelTripplePatterns FAILD. Expected %v found %v\n", expected, result)
// 			return
// 		}
// 	}
// 	t.Logf("GenerateLevelTripplePatterns PASSED. Expected %v found %v\n", expected, result)
// }

// func TestMakeQueries(t *testing.T) {
// 	// Given
// 	tripplePatterns := []string{"t1 . t2", "t1", "t2", " "}
// 	var queries []Query

// 	// When
// 	MakeQueries(tripplePatterns, &queries, 3) // results should be: [select * where {t1.t2}, select * where {t1}, select * where {t2}, select * where { }]

// 	// Then
// 	var q1, q2, q3, q4 Query
// 	q1.Query = "select * where { t1 . t2 } limit 4"
// 	q2.Query = "select * where { t1 } limit 4"
// 	q3.Query = "select * where { t2 } limit 4"
// 	q4.Query = "select * where {  } limit 4"
// 	expected := []Query{q1, q2, q3, q4}
// 	for i, q := range queries {
// 		if q.Query != expected[i].Query {
// 			t.Errorf("MakeQueries FAILD. Expected %v found %v\n", expected, queries)
// 			return
// 		}
// 	}
// 	t.Logf("MakeQueries PASSED. Expected %v found %v\n", expected, queries)
// }

// func TestMakeLattice(t *testing.T) {
// 	// Given
// 	var initialQuery Query
// 	initialQuery.Query = "select * where { t1 . t2 }"

// 	var queries []Query

// 	// When
// 	MakeLattice(initialQuery, &queries, 3) // results should be: [select * where {t1.t2}, select * where {t1}, select * where {t2}, select * where { }]

// 	// Then
// 	var q1, q2, q3, q4 Query
// 	q1.Query = "select * where { t1 . t2 } limit 4"
// 	q2.Query = "select * where { t1 } limit 4"
// 	q3.Query = "select * where { t2 } limit 4"
// 	q4.Query = "select * where {  } limit 4"

// 	expected := []Query{q1, q2, q3, q4}
// 	for i, q := range queries {
// 		if q.Query != expected[i].Query {
// 			t.Errorf("MakeLattice FAILD. Expected %v found %v\n", expected, queries)
// 			return
// 		}
// 	}
// 	t.Logf("MakeLattice PASSED. Expected %v found %v\n", expected, queries)
// }

// func TestIsDirectParent(t *testing.T) {

// 	// Given
// 	var q1 Query
// 	var q2 Query

// 	// Potential parent
// 	q1.Query = "select * where {t1 . t2 . t3}"

// 	// Potential children
// 	q2.Query = "select * where {t1 . t2}"

// 	// When
// 	result := IsDirectParent(q1, q2)

// 	// Then
// 	expected := true

// 	if result {
// 		t.Logf("IsDirectParent PASSED. Expected %v found %v\n", expected, result)
// 	} else {
// 		t.Errorf("IsDirectParent FAILD. Expected %v found %v\n", expected, result)
// 		return
// 	}
// }

// func TestSetSuperQueries(t *testing.T) {
// 	// Given
// 	var q1, q2, q3, q4 Query
// 	q1.Query = "select * where {t1 . t2}"
// 	q2.Query = "select * where {t1}"
// 	q3.Query = "select * where {t2}"
// 	q4.Query = "select * where { }"
// 	var queries []Query = []Query{q1, q2, q3, q4}

// 	// When
// 	SetSuperQueries(&queries) // results should be: [{select * where {t1.t2}, []}, {select * where {t1}, [t1.t2]}, {select * where {t2}, [t1.t2]}, {select * where { }, [t1, t2]}]

// 	// Then
// 	q1.Parents = []string{""}
// 	q2.Parents = []string{"t1 . t2"}
// 	q3.Parents = []string{"t1 . t2"}
// 	q4.Parents = []string{"t1", "t2"}

// 	expected := []Query{q1, q2, q3, q4}
// 	for i, q := range queries {
// 		if q.Query != expected[i].Query {
// 			t.Errorf("SetParents FAILD. Expected %v found %v\n", expected, queries)
// 			return
// 		}
// 		for j, p := range q.Parents {
// 			if p != expected[i].Parents[j] {
// 				t.Errorf("SetParents FAILD. Expected %v found %v\n", expected, queries)
// 				return
// 			}
// 		}
// 	}
// 	t.Logf("SetParents PASSED. Expected %v found %v\n", expected, queries)
// }

func TestContainsKey(t *testing.T) {

	// Given
	var q1, q2, q3, q4, q5 Query
	// q5 to test that it return false
	q5.Query = "select * where { t1 . t2 . t3 }"

	q1.Query = "select * where { t1 . t2 }"

	q2.Query = "select * where { t1 }"

	q3.Query = "select * where { t2 }"

	q4.Query = "select * where { }"

	var queries map[*Query]bool = make(map[*Query]bool)
	queries[&q1] = true
	queries[&q2] = true
	queries[&q3] = true
	queries[&q4] = true

	// When
	result := ContainsKey(&queries, q1)

	// Then
	expected := true

	if result == expected {
		t.Logf("ContainsKey PASSED. Expected %v found %v\n", expected, result)
		return
	} else {
		t.Errorf("ContainsKey FAILED. Expected %v found %v\n", expected, result)
	}
}

// func TestFindQuery(t *testing.T) {

// 	// Given
// 	var q1, q2, q3 Query

// 	q1.Query = "select * where {t1 . t2}"

// 	q2.Query = "select * where {t1}"

// 	q3.Query = "select * where {t2}"

// 	var queries []Query
// 	queries = append(queries, q1)
// 	queries = append(queries, q2)
// 	queries = append(queries, q3)

// 	// When
// 	index, found := FindQuery(queries, q3)

// 	// Then
// 	expectedIndex := 2
// 	expectedFound := true

// 	if expectedIndex == index && expectedFound == found {
// 		t.Logf("FindQuery PASSED. ExpectedIndex %v found %v - ExpectedFound %v found %v\n", expectedIndex, index, expectedFound, found)
// 		return
// 	} else {
// 		t.Errorf("FindQuery FAILED. ExpectedIndex %v found %v - ExpectedFound %v found %v\n", expectedIndex, index, expectedFound, found)
// 	}
// }

// func TestRemoveQuery(t *testing.T) {

// 	// Given
// 	var q1, q2, q3 Query

// 	q1.Query = "select * where {t1 . t2}"

// 	q2.Query = "select * where {t1}"

// 	q3.Query = "select * where {t2}"

// 	var queries []Query
// 	queries = append(queries, q1)
// 	queries = append(queries, q2)
// 	queries = append(queries, q3)

// 	// When
// 	newListQuery := RemoveQuery(queries, 0)

// 	expected := []Query{q2, q3}

// 	// Then
// 	for _, q := range expected {
// 		if q.Query == "select * where {t1 . t2}" {
// 			t.Errorf("RemoveQuery FAILED. Expected %v found %v\n", expected, newListQuery)
// 		}
// 	}
// 	t.Logf("RemoveQuery PASSED. Expected %v found %v\n", expected, newListQuery)
// }

// func TestGetQueryVariables(t *testing.T) {
// 	// Given
// 	var query Query
// 	query.Query = "SELECT * WHERE { ?subject <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://dbpedia.org/ontology/Film> . ?subject <http://dbpedia.org/ontology/starring> <http://dbpedia.org/resource/Sandra_Dee> . ?subject <http://dbpedia.org/ontology/starring> ?actors . ?subject <http://www.w3.org/2000/01/rdf-schema#comment> ?abstract . ?subject <http://www.w3.org/2000/01/rdf-schema#label> ?label . ?subject <http://dbpedia.org/ontology/releaseDate> ?released }"

// 	// When
// 	result := GetQueryVariables(query)

// 	// Then
// 	expected := 5

// 	if expected == len(result) {
// 		t.Logf("GetQueryVariables PASSED. Expected %v found %v \n", expected, result)
// 		return
// 	} else {
// 		t.Errorf("GetQueryVariables FAILED. Expected %v found %v \n", expected, result)
// 	}
// }

// func TestRemoveQueryTriplePattern(t *testing.T) {
// 	// Given
// 	var query Query
// 	query.Query = "select * where { t1 . t2 . t3 . t4 }"

// 	// When
// 	newQuery := RemoveQueryTriplePattern(query, "t3", 3)

// 	// Then
// 	var expectedQuery Query
// 	expectedQuery.Query = "select * where { t1 . t2 . t4 } limit 4"
// 	if expectedQuery.Query == newQuery.Query {
// 		t.Logf("RemoveQueryTriplePattern PASSED. Expected %v found %v \n", expectedQuery.Query, newQuery.Query)
// 		return
// 	} else {
// 		t.Errorf("RemoveQueryTriplePattern FAILED. Expected %v found %v \n", expectedQuery.Query, newQuery.Query)
// 	}
// }

// func TestExtractTPProperties(t *testing.T) {
// 	// Given
// 	tp := " ?subject <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://dbpedia.org/ontology/Film>"

// 	// when
// 	s, p, o := ExtractTPProperties(tp)

// 	// Then
// 	sExpected := "?subject"
// 	pExpected := "<http://www.w3.org/1999/02/22-rdf-syntax-ns#type>"
// 	oExpected := "<http://dbpedia.org/ontology/Film>"

// 	if s == sExpected && p == pExpected && o == oExpected {
// 		t.Logf("ExtractTPProperties PASSED. Expected %v %v %v found %v %v %v \n", sExpected, pExpected, oExpected, s, p, o)
// 		return
// 	} else {
// 		t.Errorf("ExtractTPProperties FAILED. Expected %v %v %v found %v %v %v \n", sExpected, pExpected, oExpected, s, p, o)
// 	}
// }

// func TestGetPredicates(t *testing.T) {
// 	// Given
// 	var q Query
// 	q.Query = "SELECT * WHERE { ?subject <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://dbpedia.org/ontology/Film> . ?subject <http://dbpedia.org/ontology/starring> <http://dbpedia.org/resource/Sandra_Dee> . ?subject <http://dbpedia.org/ontology/starring> ?actors . ?subject <http://www.w3.org/2000/01/rdf-schema#comment> ?abstract . ?subject <http://www.w3.org/2000/01/rdf-schema#label> ?label . ?subject <http://dbpedia.org/ontology/releaseDate> ?released }"

// 	// when
// 	result := GetPredicates(q)

// 	// Then
// 	expected := []string{"<http://www.w3.org/1999/02/22-rdf-syntax-ns#type>", "<http://dbpedia.org/ontology/starring>", "<http://www.w3.org/2000/01/rdf-schema#comment>", "<http://www.w3.org/2000/01/rdf-schema#label>", "<http://dbpedia.org/ontology/releaseDate>"}

// 	if len(result) == len(expected) {
// 		for i, p := range expected {
// 			if p != result[i] {
// 				t.Errorf("GetPredicates FAILED. Expected %v found %v \n", p, expected[i])
// 				break
// 			}
// 		}
// 		t.Logf("GetPredicates PASSED. Expected %v found %v \n", expected, result)
// 	} else {
// 		t.Errorf("GetPredicates FAILED. Expected %v found %v \n", len(expected), len(result))
// 	}
// }

// func TestExistString(t *testing.T) {
// 	// Given
// 	strs := []string{"<http://www.w3.org/1999/02/22-rdf-syntax-ns#type>", "<http://dbpedia.org/ontology/starring>", "<http://www.w3.org/2000/01/rdf-schema#comment>"}
// 	str := "<http://www.w3.org/2000/01/rdf-schema#comment>"

// 	// when
// 	result := ExistString(strs, str)

// 	// Then
// 	expected := true

// 	if result == expected {
// 		t.Logf("ExistString PASSED. Expected %v found %v \n", expected, result)
// 		return
// 	} else {
// 		t.Errorf("ExistString FAILED. Expected %v found %v \n", expected, result)
// 	}
// }
