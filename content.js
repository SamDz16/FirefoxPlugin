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

			document.querySelector("#options").append(insertToDOM({id: "isfailing", value: isfailing, text: "IsFailing returns ?"}))

			// Displaying the results
			document.querySelector("#options").append(insertToDOM({id: "xss", value: resGlobal[0], text: "List of XSS"}, 15))
			document.querySelector("#options").append(insertToDOM({id: "mfis", value: resGlobal[1], text: "List of MFIS"}, 15))

			addTitleToDom(document.querySelector("#options"), {title: "Statistics"}, "h3")
			document.querySelector("#options").append(insertToDOM({id: "nb", value: resGlobal[2], text: "Number of executed queries"}))
			document.querySelector("#options").append(insertToDOM({id: "etMakeLattice", value: resGlobal[3], text: "Execution Time for Make Lattice algorithm"}))
			document.querySelector("#options").append(insertToDOM({id: "etAllQueries", value: resGlobal[4], text: "Execution Time for All subqueries of the initial query"}))


			// Call the base Algorithm in DB Pedia
			resGlobal = await BFS(query, +KInput.value, "https://dbpedia.org/sparql");

			// CALLING THE ISFAILING ALGORITHM
			isfailing = await isFailing("https://dbpedia.org/sparql", query);

			document.querySelector("#options").append(document.createElement("hr"))
			addTitleToDom(document.querySelector("#options"), { title: "BFS ALGORITHM" })
			document.querySelector("#options").append(document.createElement("hr"))

			insertToDOM({id: "isfailing", value: isfailing, text: "IsFailing returns ?"})

			// Displaying the results
			document.querySelector("#options").append(insertToDOM({id: "xss", value: resGlobal[0], text: "List of XSS"}, 15))
			document.querySelector("#options").append(insertToDOM({id: "mfis", value: resGlobal[1], text: "List of MFIS"}, 15))

			addTitleToDom(document.querySelector("#options"), {title: "Statistics"}, "h3")
			document.querySelector("#options").append(insertToDOM({id: "nb", value: resGlobal[2], text: "Number of executed queries"}))
			document.querySelector("#options").append(insertToDOM({id: "etMakeLattice", value: resGlobal[3], text: "Execution Time for Make Lattice algorithm"}))
			document.querySelector("#options").append(insertToDOM({id: "etAllQueries", value: resGlobal[4], text: "Execution Time for All subqueries of the initial query"}))


			// Call the base Algorithm in DB Pedia
			resGlobal = await Var(query, +KInput.value, "https://dbpedia.org/sparql");

			// CALLING THE ISFAILING ALGORITHM
			isfailing = await isFailing("https://dbpedia.org/sparql", query);

			document.querySelector("#options").append(document.createElement("hr"))
			addTitleToDom(document.querySelector("#options"), { title: "Var ALGORITHM" })
			document.querySelector("#options").append(document.createElement("hr"))

			insertToDOM({id: "isfailing", value: isfailing, text: "IsFailing returns ?"})

			// Displaying the results
			document.querySelector("#options").append(insertToDOM({id: "xss", value: resGlobal[0], text: "List of XSS"}, 15))
			document.querySelector("#options").append(insertToDOM({id: "mfis", value: resGlobal[1], text: "List of MFIS"}, 15))

			addTitleToDom(document.querySelector("#options"), {title: "Statistics"}, "h3")
			document.querySelector("#options").append(insertToDOM({id: "nb", value: resGlobal[2], text: "Number of executed queries"}))
			document.querySelector("#options").append(insertToDOM({id: "etMakeLattice", value: resGlobal[3], text: "Execution Time for Make Lattice algorithm"}))
			document.querySelector("#options").append(insertToDOM({id: "etAllQueries", value: resGlobal[4], text: "Execution Time for All subqueries of the initial query"}))

			// insertStats(document.querySelector("#options"), resGlobal)
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

const insertToDOM = (data, numberOfRows = 0) => {
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
		textarea.setAttribute("disabled", true)
		
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