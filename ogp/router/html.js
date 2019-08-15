exports.buildHtmlHandler = function buildHtmlHandler(req, res) {
  const name = req.query.name;
  const tsundoku = req.query.tsundoku;
  const kidoku = req.query.kidoku;
  const count = req.query.count;

  const code = buildHtml(name, tsundoku, kidoku, count);
  res.writeHead(200, {
    'content-Type': 'text/html',
    'Content-Length': code.length,
    Expires: new Date().toUTCString()
  });
  res.end(code);
};

function buildHtml(name, tsundoku, kidoku, count) {
  const header = '';
  const body =
    '<p>name:' +
    name +
    '</p><p>tsundoku:' +
    tsundoku +
    '</p><p>kidoku:' +
    kidoku +
    '</p><p>count:' +
    count +
    '</p>';

  return (
    '<!DOCTYPE html>' +
    '<html><header>' +
    header +
    '</header><body>' +
    body +
    '</body></html>'
  );
}
