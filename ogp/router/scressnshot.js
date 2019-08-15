const puppeteer = require('puppeteer');

exports.screenshotHandler = function screenshotHandler(req, res) {
  const filename = 'screenshot.png';
  const name = 'tester';
  const tsundoku = '5000';
  const kidoku = '1000';
  const count = '13';

  screenshot(filename, name, tsundoku, kidoku, count);
  res.end(filename);
};

function screenshot(filename, name, tsundoku, kidoku, count) {
  const baseUrl = 'http://localhost:4000/html';
  const url =
    baseUrl +
    '?name=' +
    name +
    '&tsundoku=' +
    tsundoku +
    '&kidoku=' +
    kidoku +
    '&count=' +
    count;

  (async () => {
    const browser = await puppeteer.launch({
      // executablePath: '/usr/bin/chromium-browser',
      args: [
        '--no-sandbox',
        '--disable-setuid-sandbox'
      ]
    });
    const page = await browser.newPage();
    await page.goto(url);
    const path = 'screenshots/' + filename;
    await page.screenshot({ path: path });

    await browser.close();
    return path;
  })();
}
