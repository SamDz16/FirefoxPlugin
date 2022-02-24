// ##################################################################################################################################################################### //
// WASM SECTION - LOAD WASM
const go = new Go();
// https://www.lias-lab.fr/ftppublic/recherche/fetch.wasm
WebAssembly.instantiateStreaming(
	fetch('http://localhost:3000/fetch.wasm'),
	go.importObject
).then((result) => {
	go.run(result.instance);
});

// Utility functions
const insertToDOM = (data, numberOfRows, disable=true) => {
	const div = document.createElement("div")
	div.classList.add("form-group", "row")

	const label = document.createElement("label")
	label.setAttribute("for", data.id)
	label.textContent = data.text

	let textarea

	if (numberOfRows === 0) {
		textarea = document.createElement("input")
		textarea.classList.add("form-control", "form-control-sm")
		textarea.setAttribute("value", data.value)
		textarea.setAttribute("name", data.id)
		textarea.setAttribute("id", data.id)
		textarea.setAttribute("disabled", true)
	} else {
		textarea = document.createElement("textarea")
		textarea.classList.add("form-control")
		textarea.setAttribute("rows", numberOfRows)
		textarea.setAttribute("name", data.id)
		textarea.setAttribute("id", data.id)
		if (disable) {
			textarea.setAttribute("disabled", true)
		}
		
		if (typeof(data.value) === "object") {
			for (v of data.value) {
				textarea.textContent += v
				textarea.textContent += "\n\n"
			}
		} else {
			textarea.textContent = data.value
		}
	}

	div.append(label, document.createElement("br"), textarea)
	return div
}

const addTitleToDom = (element, data, heading = 'h1') => {
	const h1 = document.createElement(heading)
	h1.textContent = data.title

	element.append(h1)
}

// link to lias: https://www.lias-lab.fr/ftppublic/recherche/fetch.wasm

// #################################################################################### PART 1 ################################################################################# //
// Create a div element to contain label and input
if(location.href === "https://dbpedia.org/sparql") {

	const isFailingDiv = document.createElement('div');

	// Creat the label element
	const isFailingLabel = document.createElement('label');
	isFailingLabel.setAttribute('for', 'algo');
	isFailingLabel.append('Do you want to use isFailing ?');
	isFailingLabel.style.marginLeft = '10px';

	// Create the checkbox element
	const isFailingInput = document.createElement('input');
	isFailingInput.setAttribute('type', 'checkbox');
	isFailingInput.setAttribute('id', 'algo');
	isFailingInput.setAttribute('name', 'algo');

	// Apprnd the two elements to the div element
	isFailingDiv.append(isFailingInput, isFailingLabel);

	// Div that contains label and input for the threshold K
	const KDiv = document.createElement("div")
	KDiv.id = "K"

	// Label of threshold K
	const KLabel = document.createElement("label")
	KLabel.setAttribute('for', 'k');
	KLabel.append('Insert the threshold K');
	KLabel.style.marginLeft = '10px';
	KLabel.style.marginBottom = '10px';

	const KInput = document.createElement('input');
	KInput.setAttribute('type', 'number');
	KInput.setAttribute('step', '1');
	KInput.setAttribute('min', '0');
	KInput.setAttribute('value', 0);
	KInput.setAttribute('id', 'k');
	KInput.setAttribute('name', 'k');

	KDiv.append(KInput, KLabel)

	// Grab the Execute Query Button from "https://dbpedia.org/sparql"
	const executeQuery = document.querySelector('#run');

	// And then append the div element right before the execute query button
	executeQuery.insertAdjacentElement('beforeBegin', isFailingDiv);
	executeQuery.insertAdjacentElement('beforeBegin', KDiv);

	// #################################################################################### PART 2 ################################################################################# //
	// let resultsInput = document.createElement('input');
	const sparqlForm = document.querySelector('#sparql_form');

	// cardinalities
	let cardinalities = ""

	// GET TRIGGERED WHENEVER THE USER CLICKS ON THE 'EXECUTE QUERY' BUTTON
	const isFailingAlgorithm = async (e) => {
		e.preventDefault();

		let isfailing
		let resGlobal

		let query = document.querySelector('#query').value;
		try {
			// Call the base Algorithm in DB Pedia
			resGlobal = await Base(query, +KInput.value, "https://dbpedia.org/sparql");

			// CALLING THE ISFAILING ALGORITHM
			isfailing = await isFailing("https://dbpedia.org/sparql", query);

			document.querySelector("#options").append(document.createElement("hr"))
			addTitleToDom(document.querySelector("#options"), { title: "BASE ALGORITHM" })
			document.querySelector("#options").append(document.createElement("hr"))

			document.querySelector("#options").append(insertToDOM({id: "isfailing", value: isfailing, text: "IsFailing returns ?"}, 0))

			// Displaying the results
			document.querySelector("#options").append(insertToDOM({id: "xss", value: resGlobal[0], text: "List of XSS"}, 15))
			document.querySelector("#options").append(insertToDOM({id: "mfis", value: resGlobal[1], text: "List of MFIS"}, 15))

			addTitleToDom(document.querySelector("#options"), {title: "Statistics"}, "h3")
			document.querySelector("#options").append(insertToDOM({id: "nb", value: resGlobal[2], text: "Number of executed queries"}, 0))
			document.querySelector("#options").append(insertToDOM({id: "etMakeLattice", value: resGlobal[3], text: "Execution Time for Make Lattice algorithm"}, 0))
			document.querySelector("#options").append(insertToDOM({id: "etAllQueries", value: resGlobal[4], text: "Execution Time for All subqueries of the initial query"}, 0))


			// Call the BFS Algorithm in DB Pedia
			resGlobal = await BFS(query, +KInput.value, "https://dbpedia.org/sparql");

			// CALLING THE ISFAILING ALGORITHM
			isfailing = await isFailing("https://dbpedia.org/sparql", query);

			document.querySelector("#options").append(document.createElement("hr"))
			addTitleToDom(document.querySelector("#options"), { title: "BFS ALGORITHM" })
			document.querySelector("#options").append(document.createElement("hr"))

			document.querySelector("#options").append(insertToDOM({id: "isfailing", value: isfailing, text: "IsFailing returns ?"}, 0))

			// Displaying the results
			document.querySelector("#options").append(insertToDOM({id: "xss", value: resGlobal[0], text: "List of XSS"}, 15))
			document.querySelector("#options").append(insertToDOM({id: "mfis", value: resGlobal[1], text: "List of MFIS"}, 15))

			addTitleToDom(document.querySelector("#options"), {title: "Statistics"}, "h3")
			document.querySelector("#options").append(insertToDOM({id: "nb", value: resGlobal[2], text: "Number of executed queries"}, 0))
			document.querySelector("#options").append(insertToDOM({id: "etMakeLattice", value: resGlobal[3], text: "Execution Time for Make Lattice algorithm"}, 0))
			document.querySelector("#options").append(insertToDOM({id: "etAllQueries", value: resGlobal[4], text: "Execution Time for All subqueries of the initial query"}, 0))

			// Call the Var Algorithm in DB Pedia
			resGlobal = await Var(query, +KInput.value, "https://dbpedia.org/sparql");

			// CALLING THE ISFAILING ALGORITHM
			isfailing = await isFailing("https://dbpedia.org/sparql", query);

			document.querySelector("#options").append(document.createElement("hr"))
			addTitleToDom(document.querySelector("#options"), { title: "Var ALGORITHM" })
			document.querySelector("#options").append(document.createElement("hr"))

			document.querySelector("#options").append(insertToDOM({id: "isfailing", value: isfailing, text: "IsFailing returns ?"}, 0))

			// Displaying the results
			document.querySelector("#options").append(insertToDOM({id: "xss", value: resGlobal[0], text: "List of XSS"}, 15))
			document.querySelector("#options").append(insertToDOM({id: "mfis", value: resGlobal[1], text: "List of MFIS"}, 15))

			addTitleToDom(document.querySelector("#options"), {title: "Statistics"}, "h3")
			document.querySelector("#options").append(insertToDOM({id: "nb", value: resGlobal[2], text: "Number of executed queries"}, 0))
			document.querySelector("#options").append(insertToDOM({id: "etMakeLattice", value: resGlobal[3], text: "Execution Time for Make Lattice algorithm"}, 0))
			document.querySelector("#options").append(insertToDOM({id: "etAllQueries", value: resGlobal[4], text: "Execution Time for All subqueries of the initial query"}, 0))
			
		} catch (err) {
			console.error('Caught exception', err);
		}
	};

	// Add event listener to track whether the user would like to use the Luis algorithms or not
	// and modify a global variable
	isFailingInput.addEventListener('change', () => {
		if (isFailingInput.checked) {
			sparqlForm.addEventListener('submit', isFailingAlgorithm);
		} else {
			sparqlForm.removeEventListener('submit', isFailingAlgorithm, {
				passive: true,
			});
		}
	})

	// FULL
	document.querySelector("#footer").insertAdjacentElement("beforebegin", document.createElement("hr"))

	const h1Full = document.createElement("h1")
	h1Full.textContent = "Full ALGORITHM"
	document.querySelector("#footer").insertAdjacentElement("beforebegin", h1Full)

	document.querySelector("#footer").insertAdjacentElement("beforebegin", document.createElement("hr"))

	document.querySelector("#footer").insertAdjacentElement("beforebegin", insertToDOM({id: "cardinalities", value: "", text: "Cardinalities"}, 7, false))

	document.querySelector("#cardinalities").addEventListener('change', () => {
		cardinalities = document.querySelector("#cardinalities").value
	})

	// Launch Full Algo
	const FullAlgoBtn = document.createElement("button")
	FullAlgoBtn.classList.add("btn", "btn-lg", "btn-m-3", "btn-info")
	FullAlgoBtn.textContent = "Launch Full Algorithm"
	FullAlgoBtn.id = "full"
	document.querySelector("#footer").insertAdjacentElement("beforebegin", FullAlgoBtn)

	document.querySelector("#full").addEventListener("click", async () => {
		cardinalities = cardinalities.split("\n")
		var cards = ""
		for (let cardinality of cardinalities) {
			cardinality = cardinality.split(/[(),]+/)
			cards += cardinality[2] + "-" + cardinality[3] + " "
		}

		// Grab the query
		let query = document.querySelector('#query').value;

		let resGlobal = await Full(query, +KInput.value, "https://dbpedia.org/sparql", cards) 
		// CALLING THE ISFAILING ALGORITHM
		isfailing = await isFailing("https://dbpedia.org/sparql", query);

		document.querySelector("#footer").insertAdjacentElement("beforebegin", insertToDOM({id: "isfailing", value: isfailing, text: "IsFailing returns ?"}, 0))

		// Displaying the results
		document.querySelector("#footer").insertAdjacentElement("beforebegin", insertToDOM({id: "xss", value: resGlobal[0], text: "List of XSS"}, 15))
		document.querySelector("#footer").insertAdjacentElement("beforebegin", insertToDOM({id: "mfis", value: resGlobal[1], text: "List of MFIS"}, 15))

		const h3Full = document.createElement("h3")
		h3Full.textContent = "Statistics"
		document.querySelector("#footer").insertAdjacentElement("beforebegin", h3Full)

		document.querySelector("#footer").insertAdjacentElement("beforebegin", insertToDOM({id: "nb", value: resGlobal[2], text: "Number of executed queries"}, 0))
		document.querySelector("#footer").insertAdjacentElement("beforebegin", insertToDOM({id: "etMakeLattice", value: resGlobal[3], text: "Execution Time for Make Lattice algorithm"}, 0))
		document.querySelector("#footer").insertAdjacentElement("beforebegin", insertToDOM({id: "etAllQueries", value: resGlobal[4], text: "Execution Time for All subqueries of the initial query"}, 0))
	})
}

// ##########################################################################################################################################################################################
// ############################################################### IF I AM IN THE http://localhost:3030/base endpoint #######################################################################
// ##########################################################################################################################################################################################


if (location.href === "http://localhost:3030/dataset.html") {

	// Create Button to launch the Base Algorithm
	const btnQuery = document.createElement("btn")
	btnQuery.setAttribute("id", "btnQuery")
	btnQuery.classList.add("btn", "btn-danger", "btn-lg")
	btnQuery.textContent = "Query using Base Algorithm"

	const h2 = document.querySelector('h2');
	h2.insertAdjacentElement("afterend", btnQuery)

	btnQuery.addEventListener("click", async () => {
		const query = document.querySelector("pre").textContent
		const K = 3

		const resGlobal = await Base(query, K, "http://localhost:3030/base")

		const container = document.createElement("div")
		container.classList.add("class", "container")

		// Displaying the results
		container.append(insertToDOM({id: "xss", value: resGlobal[0], text: "List of XSS"}, 15))
		container.append(insertToDOM({id: "mfis", value: resGlobal[1], text: "List of MFIS"}, 15))

		container.append(addTitleToDom({title: "Statistics"}))
		container.append(insertToDOM({id: "nb", value: resGlobal[2], text: "Number of executed queries"}))
		container.append(insertToDOM({id: "etMakeLattice", value: resGlobal[3], text: "Execution Time for Make Lattice algorithm"}))
		container.append(insertToDOM({id: "etAllQueries", value: resGlobal[4], text: "Execution Time for All subqueries of the initial query"}))

		document.querySelector("body").append(container)
	})
}

class Query {
	constructor(query, parents) {
		this.query = query
		this.parents = parents
	}
}

const getQueryTriplePatterns = ({query}) => {
	let start = false
	let res = ""
	for (let i = 0; i < query.length; i++) {
		if (query[i] == "{") {
			start = true
			continue
		}

		if (query[i] == "}") {
			start = false
			continue
		}

		if (start) {
			res += query[i]
		}
	}
	res = res.trim()
	return res.split(" . ")
}

const makeLattice = (query) => {

}

const algoBase = (query, k, endpoint) => {
	const mfis = []
	const xss = []
	const fis = []
	const executedQueries = []
	const nb = 0

	const q = new Query("SELECT * WHERE { ?subject <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://dbpedia.org/ontology/Film> . ?subject <http://dbpedia.org/ontology/starring> <http://dbpedia.org/resource/Sandra_Dee> . ?subject <http://dbpedia.org/ontology/starring> ?actors . ?subject <http://www.w3.org/2000/01/rdf-schema#comment> ?abstract . ?subject <http://www.w3.org/2000/01/rdf-schema#label> ?label . ?subject <http://dbpedia.org/ontology/releaseDate> ?released }", [])
	const res = getQueryTriplePatterns(q)
	console.log(res)

	function k_combinations(set, k) {
		var i, j, combs, head, tailcombs;
		
		// There is no way to take e.g. sets of 5 elements from
		// a set of 4.
		if (k > set.length || k <= 0) {
			return [];
		}
		
		// K-sized set has only one K-sized subset.
		if (k == set.length) {
			return [set];
		}
		
		// There is N 1-sized subsets in a N-sized set.
		if (k == 1) {
			combs = [];
			for (i = 0; i < set.length; i++) {
				combs.push([set[i]]);
			}
			return combs;
		}
		
		// Assert {1 < k < set.length}
		
		// Algorithm description:
		// To get k-combinations of a set, we want to join each element
		// with all (k-1)-combinations of the other elements. The set of
		// these k-sized sets would be the desired result. However, as we
		// represent sets with lists, we need to take duplicates into
		// account. To avoid producing duplicates and also unnecessary
		// computing, we use the following approach: each element i
		// divides the list into three: the preceding elements, the
		// current element i, and the subsequent elements. For the first
		// element, the list of preceding elements is empty. For element i,
		// we compute the (k-1)-computations of the subsequent elements,
		// join each with the element i, and store the joined to the set of
		// computed k-combinations. We do not need to take the preceding
		// elements into account, because they have already been the i:th
		// element so they are already computed and stored. When the length
		// of the subsequent list drops below (k-1), we cannot find any
		// (k-1)-combs, hence the upper limit for the iteration:
		combs = [];
		for (i = 0; i < set.length - k + 1; i++) {
			// head is a list that includes only our current element.
			head = set.slice(i, i + 1);
			// We take smaller combinations from the subsequent elements
			tailcombs = k_combinations(set.slice(i + 1), k - 1);
			// For each (k-1)-combination we join it with the current
			// and store it to the set of k-combinations.
			for (j = 0; j < tailcombs.length; j++) {
				combs.push(head.concat(tailcombs[j]));
			}
		}
		return combs;
	}

	var set = ["t1", "t2", "t3", "t4"]

	var k, i, combs, k_combs;
	combs = [];
	
	// Calculate all non-empty k-combinations
	for (k = 1; k <= set.length; k++) {
		k_combs = k_combinations(set, k);
		for (i = 0; i < k_combs.length; i++) {
			combs.push(k_combs[i]);
		}
	}
	console.log(combs.reverse());

}

algoBase()
