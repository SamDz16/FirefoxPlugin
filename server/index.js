const express = require("express")
const axios = require("axios")
const cors = require("cors")
const app = express()

app.use(express.json())
app.use(
    cors({
        origin: "https://dbpedia.org"
    })
)

app.get("/", (req, res) => {
    res.send("You are in the home local server. To get the todos, go to /api/todos")
})

app.get("/wasm", (req, res) => {
    res.send("hi")
})

app.get("/api/todos", async (req,res) => {
    const {data: todos} = await axios.get("https://jsonplaceholder.typicode.com/todos?_limit=5")
    // console.log(posts)
    res.send(todos)
})

app.listen(3000, () => { console.log(`App is up and running on : http://localhost:3000`)})

