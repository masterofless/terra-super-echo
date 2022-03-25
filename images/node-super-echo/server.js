'use strict';

const express = require('express');
const PORT = 8080;
const HOST = '0.0.0.0';

const app = express();
app.get('/', (req, res) => {
  res.send('Hello World');
});

app.get('/utterances', (req, res) => {
  const responseData = {
    utterance: "Hello, This is me shouting for joy."
  }
  const jsonContent = JSON.stringify(responseData);
  res.end(jsonContent);
});

app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);
