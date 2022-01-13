// ##################################################################################################################################################################### //
// WASM SECTION - LOAD WASM
const go = new Go();

WebAssembly.instantiateStreaming(
	fetch('http://localhost:9090/fetch.wasm'),
	go.importObject
).then((result) => {
	go.run(result.instance);
});

// #################################################################################### PART 1 ################################################################################# //
// Create a div element to contain label and input
const div = document.createElement('div');
// Creat the label element
const label = document.createElement('label');
label.setAttribute('for', 'algo');
label.append('Do you want to use Louise algorithms ?');
label.style.marginLeft = '10px';
// Create the checkbox element
const input = document.createElement('input');
input.setAttribute('type', 'checkbox');
input.setAttribute('id', 'algo');
input.setAttribute('name', 'algo');
// Apprnd the two elements to the div element
div.append(input, label);
// Grab the Execute Query Button from "https://dbpedia.org/sparql"
const executeQuery = document.querySelector('#run');
// And then append the div element right before the execute query button
executeQuery.insertAdjacentElement('beforeBegin', div);

// #################################################################################### PART 2 ################################################################################# //

// JSON PLACEHOLDER API
const BASE_URL = 'https://jsonplaceholder.typicode.com/users?_limit=5';
let users = [];

// loadUsers from JSONPLACEHOLDER API
const loadUsers = async (url) => {
	const response = await fetch(url);
	const users = await response.json();
	return users;
};

// LOAD USERS AND SAVE THEM TO GLOBAL VARIABLE 'USERS'
window.onload = async () => {
	users = await loadUsers(BASE_URL);

	// GET 'TOODS' DATA FROM LOCAL SERVER
	const btn = document.createElement('button');
	const hr = document.createElement('hr');
	btn.id = 'todos-btn';
	btn.classList.add('btn', 'btn-outline-primary', 'm-2');
	btn.textContent = 'Load todos from local server';
	document.querySelector('#main').append(hr, btn);
	btn.addEventListener('click', async () => {
		const response = await fetch('http://localhost:3000/api/todos');
		const todos = await response.json();
		console.log(todos);
		const rootDiv = document.createElement('div');
		rootDiv.id = 'rootTodos';
		rootDiv.innerHTML = `<h2>List of Todos fetched from local server</h2> <hr>`;
		rootDiv.classList.add('text-center');
		const todoDiv = document.createElement('div');
		rootDiv.append(todoDiv);
		todoDiv.setAttribute('v-for', 'todo in todos');
		todoDiv.innerHTML = `
	        <h3>Todo title: {{ todo.title }}</h3>
	        <p>Completed: {{ todo.completed }}</p>
	    `;
		btn.insertAdjacentElement('afterend', rootDiv);
		// Create Vue component
		new Vue({
			el: '#rootTodos',
			data: () => {
				return { todos };
			},
		});
	});

	const btnFetch = document.createElement('button');
	const hrFetch = document.createElement('hr');
	btnFetch.classList.add('btn', 'btn-outline-primary', 'm-2');
	btnFetch.textContent = 'GO &WASM with HTTP Requests';
	document.querySelector('#main').append(hrFetch, btnFetch);

	btnFetch.addEventListener('click', async () => {
		let endpoint = 'https://dbpedia.org/sparql';
		let query = 'select distinct ?Concept where {[] a ?Concept} LIMIT 100';
		try {
			const response = await fetchDBPedia(endpoint, query);

			// Response is in XML format
			const str = await response.text();
			const data = await new window.DOMParser().parseFromString(
				str,
				'text/xml'
			);
			const results = data.getElementsByTagName('uri');
			document.querySelector('#main').innerHTML += '<h1>The results are :</h1>';
			let res = '';
			for (result of results) {
				res += result.textContent + ' | ';
			}
			document.querySelector('#main').append(res);
		} catch (err) {
			console.error('Caught exception', err);
		}
	});
};

// Create a global variable to hold true or fale in order to know whether a user has thecked the checkbox or not
let luisAlgorithmsChecked = false;
// GET TRIGGERED WHENEVER THE USER CLICKS ON THE 'EXECUTE QUERY' BUTTON
const louiseAlgorithm = async (e) => {
	e.preventDefault();

	const rootDiv = document.createElement('div');
	rootDiv.id = 'rootUsers';
	rootDiv.innerHTML = `<h2>List of users</h2> <hr>`;
	rootDiv.classList.add('text-center');

	const userDiv = document.createElement('div');
	rootDiv.append(userDiv);

	userDiv.setAttribute('v-for', 'user in users');
	userDiv.innerHTML = `
        <h3>{{ user.name }} {{ user.username }}</h3>
        <p>{{ user.email }}</p>
    `;
	document
		.querySelector('#options')
		.insertAdjacentElement('beforebegin', rootDiv);

	console.log(users);
	// Create Vue component
	new Vue({
		el: '#rootUsers',
		data: () => {
			return { users };
		},
	});
};
// Add event listener to track whether the user would like to use the Luis algorithms or not
// and modify a global variable
input.addEventListener('change', () => {
	luisAlgorithmsChecked = input.checked;

	const sparqlForm = document.querySelector('#sparql_form');

	if (luisAlgorithmsChecked) {
		sparqlForm.addEventListener('submit', louiseAlgorithm);
	} else {
		sparqlForm.removeEventListener('submit', louiseAlgorithm, {
			passive: true,
		});
	}
});
