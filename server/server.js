const express = require("express")
const cors = require("cors")

const app = express();

app.use(
    cors({
        origin: "https://dbpedia.org"
    })
)

app.use(express.static("public", {
    setHeaders: (res, path) => {
        if(path.endsWith(".wasm")) {
            res.set("Content-Type", "application/wasm")
        }
    }
}))

app.listen(3000, () => console.log("server app and listening on http://localhost:3000"))