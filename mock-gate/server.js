'use strict';

const express = require('express');

// Constants
const PORT = 8080;
const HOST = '0.0.0.0';

// App
const app = express();
app.use(express.json());       // to support JSON-encoded bodies
app.use(express.urlencoded()); // to support URL-encoded bodies


var server = require('http').createServer(app);
var io = require('socket.io')(server);

var gateClient

io.on('connection', (client) => {
  gateClient = client
});


app.get('/', (req, res) => {
  res.send('Hello world\n');
})

app.post('/', (req, res) => {
  res.status(200).end()
  if (gateClient != null) {
    gateClient.emit("gate", true)
  }
})

server.listen(PORT);
console.log(`Running on http://${HOST}:${PORT}`);