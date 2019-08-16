const puppeteer = require('puppeteer');

exports.screenshotHandler = function screenshotHandler(req, res) {
  const filename = req.query.filename;
  const tsundoku = req.query.tsundoku;
  const kidoku = req.query.kidoku;
  const count = req.query.count;

  screenshot(filename, tsundoku, kidoku, count);
  res.end(filename);
};

function screenshot(filename, tsundoku, kidoku, count) {
  const baseUrl = 'http://localhost:4000/html';
  const url =
    baseUrl +
    '?tsundoku=' +
    tsundoku +
    '&kidoku=' +
    kidoku +
    '&count=' +
    count;

  (async () => {
    const browser = await puppeteer.launch({
      args: [
        '--no-sandbox',
        '--disable-setuid-sandbox'
      ]
    });
    const page = await browser.newPage();
    await page.goto(url);
    const path = 'screenshots/' + filename;
    await page.screenshot({ path: path, fullPage: true });

    await browser.close();
    return path;
  })();
}
