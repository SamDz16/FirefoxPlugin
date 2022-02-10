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

		let results = [];
		var isfailing = 0;

		// let leoQuery = 'SELECT * WHERE { ?athlete  rdfs:label  "Lionel Messi"@en ; dbo:number  ?number }'
		let query = document.querySelector('#query').value;
		try {
			// Call the base Algorithm in DB Pedia
			await baseAlgorithm(query, 100, "https://dbpedia.org/sparql");

			// CALLING THE ISFAILING ALGORITHM
			isfailing = await isFailing("https://dbpedia.org/sparql", query);

			// CALLING THE EXECUTEQPARQLALGORITHM
			const response = await executeSPARQLQuery("https://dbpedia.org/sparql", query, 101);

			// FROM SPARQL RESULTS TO XML
			const getXMLData = async (response) => {
				const str = await response.text();
				const data = await new window.DOMParser().parseFromString(
					str,
					'text/xml'
				);
				const results = data.getElementsByTagName('uri');
				return results;
			};

			// Response is in XML format
			results = await getXMLData(response);

			// TEST VALE OF ISFAILING
			if (isfailing === 1) {
				// There is no results
				document.querySelector(
					'#options'
				).innerHTML += `<h1 class="text-center text-danger">isFailing returns: <b>${isfailing}</b></h1>`;
			} else {
				// There is at least one result
				document.querySelector(
					'#options'
				).innerHTML += `<h1 class="text-center text-success">isFailing returns: <b>${isfailing}</b></h1>`;

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
}

// Full GO
const baseAlgorithm = async (query, K, endpoint) => {

	const initialQuery = query

	const resGlobal = await Base(initialQuery, K, endpoint);
	console.log('LIST OF XSSs : ');
	console.log(resGlobal[0]);

	console.log('LIST OF MFISs : ');
	console.log(resGlobal[1]);

	console.log('NUMBER OF EXECUTED QUERIES: ');
	console.log(resGlobal[2]);
};

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

		// Number of executeed queries
		const nb = document.createElement("h1")
		nb.classList.add("text-center")
		nb.innerHTML = `<h1>3. Number of executed Queries: </h1>`
		nb.innerHTML += resGlobal[2]
		document.querySelector("#results-block").insertAdjacentElement("afterend", nb)

		// List MFIS
		const rootMFIS = document.createElement("div")
		rootMFIS.classList.add("text-center", "text-danger")
		rootMFIS.id = "rootMFIS"
		rootMFIS.innerHTML += "<h1>2. Liste des MFIS:</h1>"

		const mfisp = document.createElement("p")
		rootMFIS.append(mfisp)
		mfisp.setAttribute("v-for", "mfis in listMFIS")
		mfisp.innerHTML = `{{mfis}}`

		document.querySelector("#results-block").insertAdjacentElement("afterend", rootMFIS)

		// Create Vue element for MFIS list
		new Vue({
			el: '#rootMFIS',
			data: () => {
				if (resGlobal[1].length === 0) {
					return { listMFIS: [`empty`] }
				} else {
					return { listMFIS: resGlobal[1] };
				}
			},
		});

		// List XSS
		const rootXSS = document.createElement("div")
		rootXSS.classList.add("text-center", "text-success")
		rootXSS.id = "rootXSS"
		rootXSS.innerHTML += "<h1>1. Liste des XSS:</h1>"

		const xssp = document.createElement("p")
		rootXSS.append(xssp)
		xssp.setAttribute("v-for", "xss in listXSS")
		xssp.innerHTML = `{{xss}}`

		document.querySelector("#results-block").insertAdjacentElement("afterend", rootXSS)

		// Create Vue element for XSS list
		new Vue({
			el: '#rootXSS',
			data: () => {
				if (resGlobal[0] === 0) {
					return { listXSS: [`empty`]};
				} else {
					return { listXSS: resGlobal[0] };
				}
			},
		});
	})
}