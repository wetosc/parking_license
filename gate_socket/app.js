'use strict';

const express = require('express');

// Constants
const PORT = 4999;
const HOST = '0.0.0.0';

// App
const app = express();
app.use(express.json());       // to support JSON-encoded bodies
app.use(express.urlencoded()); // to support URL-encoded bodies
app.use(express.static('www'))

var server = require('http').createServer(app);
var io = require('socket.io')(server);

// var clients

io.on('connection', (client) => {
    // gateClient = client
});


app.get('/', (req, res) => {
    //   res.send('Hello world\n');
    res.end()
})

app.post('/', (req, res) => {
    res.status(200).end()
    
    io.sockets.emit("gate", { open: true, plate: req.body.plate })
    
    setTimeout(() => {
        io.sockets.emit("gate", { open: false })
    }, 30 * 1000);
})

server.listen(PORT);
console.log(`Running on http://${HOST}:${PORT}`);