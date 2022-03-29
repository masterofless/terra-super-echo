'use strict';

const express = require('express');
const PORT = 8080;
const HOST = '0.0.0.0';

const app = express();
app.get('/', (req, res) => {
  res.send('Hello World');
});

app.get('/utterance', (req, res) => {
  var yodel = '';
  getUtterance(function(data) {
    const responseData = {
      utterance: data.yodel
    }
    res.end(JSON.stringify(responseData));
  });
});

app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);

var http = require('http');

function getUtterance(success) {
  var options = {
    host: 'terra-super-echo-yodeler',
    port: 8002,
    path: '/yodel',
    method: 'GET',
    headers: {}
  };

  var responseObject;
  var req = http.request(options, function(res) {
    res.setEncoding('utf-8');

    var responseString = '';

    res.on('data', function(data) {
      responseString += data;
    });

    res.on('end', function() {
      console.log(`Got response: ${responseString}`);
      responseObject = JSON.parse(responseString);
      success(responseObject);
    });
  });

  req.write('');
  req.end();
}
