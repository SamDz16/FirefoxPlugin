# Trigger Louise Algorithms - Chrome/Firefox Extension
This is a chrome/firefow extensions which gets triggered whenever user heads to the `https://dbpedia.org/sparql` enpoint. It adds to the Document Object Model (DOM) of the actual page a checkbox that user can click and enable in order to change the default form submission. For now all what it does is to grab somme data from the famous `jsonplaceholder` API and display 5 users. Another button `Load todos from local server` is added to the bottom of the page to load s5 todos from ocal server deployed by nodejs. Finally, a last button is added to the end of the webpage `GO & WASM with HTTP Request` where when clicked it makes an HTTP POST request to the `https://dbpedia.org/sparql` with this body : `select distinct ?Concept where {[] a ?Concept} LIMIT 100` for now.

## Installation
In order to enable correctely this extension into Firefox, you have to have some prerequisites:
* Have Chrome/Firefox installed in your machine (quite obvious xD)
* Load your extension into your browser
* Go and install the GO programming language at https://go.dev (I am using the 1.16.12 version)
* Make sure to have node and npm installed
* Make sure to add the `C:\Program Files\Go\bin` to your PATH environment variable

## Step 1: Compile your GO file to WebAssembly
The bellowed command is executed in the `wasm` folder in order to compile the `main.go` to get the wasm equivalent `fetch.wasm` in the `wasm/assets` folder.

```shell
cd wasm && GOOS=js GOARCH=wasm go build -o ./assets/fetch.wasm ./main.go
```

## Step 2: Install npm packages
In order to launch your node express server to serve your ressources, you have to have node modules installed inside your folder, for that, you have to execute this command:
```shell
npm i nodemon -g && cd server && npm install && npm start
```
Note: i've made an npm script called *start, it helps to quickly launch the server using nodemon

## Step 3: Launch the server
In order to launch the server to serve the `assets` folder (`index.html`), run the bellow command in the `/wasm//server` folder:
```shell
cd wasm/server && go run server.go
```

## Step 4: Open the endpoint
Open your Chrome/Firefox browser and head to this endpoint `https://dbpedia.org/sparql`. Enjoy!
