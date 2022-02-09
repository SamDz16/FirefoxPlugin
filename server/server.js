const express = require("express")
const cors = require("cors")

const app = express();

app.use(
    cors({
        origin: "*"
    })
)

app.use(express.static("public"))

app.listen(3000, () => console.log("server app and listening on http://localhost:3000"))