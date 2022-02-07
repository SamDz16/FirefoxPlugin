// ##################################################################################################################################################################### //
// WASM SECTION - LOAD WASM
const go = new Go();

WebAssembly.instantiateStreaming(
	fetch('http://localhost:3000/fetch.wasm'),
	go.importObject
).then((result) => {
	go.run(result.instance);
});

// #################################################################################### PART 1 ################################################################################# //
// Create a div element to contain label and input
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

// Grab the Execute Query Button from "https://dbpedia.org/sparql"
const executeQuery = document.querySelector('#run');

// And then append the div element right before the execute query button
executeQuery.insertAdjacentElement('beforeBegin', isFailingDiv);

// #################################################################################### PART 2 ################################################################################# //
let resultsInput = document.createElement('input');
let rootDiv = document.createElement('div');
const sparqlForm = document.querySelector('#sparql_form');

// Create a global variable to hold true or fale in order to know whether a user has thecked the checkbox or not
let luisAlgorithmsChecked = false;

// GET TRIGGERED WHENEVER THE USER CLICKS ON THE 'EXECUTE QUERY' BUTTON
const isFailingAlgorithm = async (e) => {
	e.preventDefault();

	await baseAlgorithm()

	let results = [];
	var isfailing = 0;

	// Grab the results
	let endpoint = 'https://dbpedia.org/sparql';

	// let leoQuery = 'SELECT * WHERE { ?athlete  rdfs:label  "Lionel Messi"@en ; dbo:number  ?number }'
	let query = document.querySelector('#query').value
	try {

		// CALLING THE ISFAILING ALGORITHM
		isfailing = await isFailing(endpoint, query)

		// CALLING THE EXECUTEQPARQLALGORITHM
		const response = await executeSPARQLQuery(endpoint, query);

		// FROM SPARQL RESULTS TO XML
		const getXMLData = async (response) => {
			const str = await response.text();
			const data = await new window.DOMParser().parseFromString(str, 'text/xml');
			const results = data.getElementsByTagName('uri');
			isfailing = results.length === 0 ? 1 : 0
			return results
		}

		// Response is in XML format
		results = await getXMLData(response)
		
		// TEST VALE OF ISFAILING
		if (isfailing === 1) {

			// There is no results
			document.querySelector('#options').innerHTML += `<h1 class="text-center text-danger">isFailing returns: <b>${isfailing}</b></h1>`;
		} else {

			// There is at least one result
			document.querySelector('#options').innerHTML += `<h1 class="text-center text-success">isFailing returns: <b>${isfailing}</b></h1>`;

			// Create a div element to contain label and input
			const resultsDiv = document.createElement('div');

			// Creat the label element
			const resultsLabel = document.createElement('label');
			resultsLabel.setAttribute('for', 'resultsLabel');
			resultsLabel.append('Do you want to display results ?');
			resultsLabel.style.marginLeft = '10px';

			// Create the checkbox element
			resultsInput.setAttribute('type', 'checkbox');
			resultsInput.setAttribute('id', 'resultsLabel');
			resultsInput.setAttribute('name', 'resultsLabel');
			
			// Apprnd the two elements to the div element
			resultsDiv.append(resultsInput, resultsLabel);

			// And then append the div element right before the execute query button
			document.querySelector('fieldset').append(resultsDiv);

			// By default, the results are hidden
			rootDiv.setAttribute('style', 'display: none');

			rootDiv.id = 'rootResults';
			rootDiv.innerHTML = `<h2>Results</h2> <hr>`;
			rootDiv.classList.add('text-center');

			const resultDiv = document.createElement('div');
			rootDiv.append(resultDiv);

			resultDiv.setAttribute('v-for', 'result in results');
			resultDiv.innerHTML = `
				<p><a :href=result.textContent>{{ result.textContent }}</a></p>
			`;
			document
				.querySelector('#options')
				.insertAdjacentElement('beforebegin', rootDiv);

			// Create Vue component
			new Vue({
				el: '#rootResults',
				data: () => {
					return { results };
				},
			});
		}
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
});

resultsInput.addEventListener('change', () => {
	if (resultsInput.checked) {
		document
			.querySelector('#rootResults')
			.setAttribute('style', 'display: block');
	} else {
		document
			.querySelector('#rootResults')
			.setAttribute('style', 'display: none');
	}
});

// const baseAlgorithm = async () => {
// 	// Generate all possible queries from initial query : lattice
// 	const queries = AllQueries("SELECT * WHERE { ?fp <http://example.com/type> <http://example.com/FullProfessor> . ?fp <http://example.com/age> ?a . ?fp <http://example.com/nationality> ?n . ?fp <http://example.com/teacherOf> ?c }")

// 	const nbs = []

// 	for (query of queries) {
// 		const params = new URLSearchParams()
// 		params.append("query", query)

// 		const response = await axios.post("http://localhost:3030/base", params)
// 		nbs.push(response.data.results.bindings.length)
// 	}

// 	const NBs = new Uint8Array(nbs)

// 	const obj = await Base("SELECT * WHERE { ?fp <http://example.com/type> <http://example.com/FullProfessor> . ?fp <http://example.com/age> ?a . ?fp <http://example.com/nationality> ?n . ?fp <http://example.com/teacherOf> ?c }", 3, NBs)
// 	console.log("result ", obj);
// }

const baseAlgorithm = async () => {
	// Generate all possible queries from initial query : lattice
	const queries = AllQueries("SELECT * WHERE { ?fp <http://example.com/type> <http://example.com/FullProfessor> . ?fp <http://example.com/age> ?a . ?fp <http://example.com/nationality> ?n . ?fp <http://example.com/teacherOf> ?c }")

	const nbs = []

	for (query of queries) {
		// const params = new URLSearchParams()
		// params.append("query", query)

		// const response = await axios.post("http://localhost:3030/base", params)
		const nb = await 
		nbs.push(response.data.results.bindings.length)
	}

	const NBs = new Uint8Array(nbs)

	const obj = await Base("SELECT * WHERE { ?fp <http://example.com/type> <http://example.com/FullProfessor> . ?fp <http://example.com/age> ?a . ?fp <http://example.com/nationality> ?n . ?fp <http://example.com/teacherOf> ?c }", 3, NBs)
	console.log("result ", obj);
}