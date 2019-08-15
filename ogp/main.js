const express = require('express');
const app = express();
const bodyParser = require('body-parser');
const html = require('./router/html');
const ss = require('./router/scressnshot');

const port = process.env.PORT || 4000;

app.use(
  bodyParser.urlencoded({
    extended: true
  })
);
app.use(bodyParser.json());

app.use('/static', express.static('public'));

app.post('/ogp', function(req, res) {
  var userName = req.body.userName;
  res.json({
    userName: userName
  });
});

app.get('/html', html.buildHtmlHandler);

app.get('/ss', ss.screenshotHandler);

app.use(express.static('screenshots'));

app.listen(port);
console.log('listen on port ' + port);
