# Trigger Louise Algorithms - Firefox Extension
This is a firefow extensions which gets triggered whenever user heads to the `https://dbpedia.org/sparql` enpoint. It adds to the Document Object Model (DOM) of the actual page a checkbox that user can click and enable in order to change the default form submission. For now all what it does is to grab somme data from the famous `jsonplaceholder` API and display 5 users. Another button `Load todos from local server` is added to the bottom of the page to load s5 todos from ocal server deployed by nodejs.

## Installation
In order to enable correctely this extension into Firefox, you have to have some prerequisites:
* Have Firefox installed in your machine (quite obvious xD)
* Head to this URL `about:debugging#/runtime/this-firefox` and load the `manifest.json`
* Go and install the GO programming language at https://go.dev
* Install TinyGO: follow this [installation guide](https://tinygo.org/getting-started/install/) to get tinygo installed into your machine
* Make sure to have node and npm installed
* Make sure to add the `C:\Program Files\Go\bin` and the `C:\tinygo\bin` to your PATH environment variable

## Step 1: Compile your GO file to WebAssembly
The bellowed command is executed in the root folder in order to compile the `main.go` to get the wasm equivalent `main.wasm` in the `assets` folder.
Note that tinygo is used to get ***tiny*** wasm file size. In my case it tooks only **181KB** for the `main.wasm` file size wherease if we use the built in go way it generates wasm file size with almost **2MB** of size.

```shell
tinygo build -o ./assets/main.wasm -target wasm ./main.go
```

## Step 2: Install npm packages
In order to launch your node express server to serve your ressources, you have to have node modules installed inside your folder, for that, you have to execute this command:
```shell
npm i nodemon -g && cd server && npm install && npm start
```
Note: i've made an npm script called *start, it helps to quickly launch the server using nodemon

## Step 3: Launch the server
In order to launch the server to serve the `assets` folder (`index.html`), run the bellow command in the `./server` folder:
```shell
cd wasm/server && go run server.go
```

## Step 4: Open the endpoint
Open your Firefox browser and head to this endpoint `https://dbpedia.org/sparql`. Enjoy!
