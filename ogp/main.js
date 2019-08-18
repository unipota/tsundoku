const express = require('express');
const app = express();
const bodyParser = require('body-parser');
const html = require('./router/html');
const ss = require('./router/screenshot');
const ejs = require('ejs');

const port = process.env.PORT || 4000;

app.use(
  bodyParser.urlencoded({
    extended: true
  })
);
app.use(bodyParser.json());
app.engine('ejs', ejs.renderFile);


app.get('/html', html.buildOgpHtmlHandler);
app.get('/ss', ss.screenshotHandler);


app.listen(port);
console.log('listen on port ' + port);
