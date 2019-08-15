// googleのトップページのスクショを取ってくる
const puppeteer = require('puppeteer');

(async () => {
  const browser = await puppeteer.launch();
  const page = await browser.newPage();
  await page.goto('http://localhost:3000/html');
  await page.screenshot({ path: 'screenshots/screenshot.png' });

  await browser.close();
})();