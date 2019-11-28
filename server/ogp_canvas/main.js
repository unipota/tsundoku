const { registerFont, createCanvas } = require("canvas");
const fs = require("fs");
const path = require("path");
const args = [...process.argv].slice(2);

if (args.length !== 4) throw Error("4 args required");
const [shareId, booksCount, tsundokuPrice, kidokuPrice] = args;
const canvasWidth = 1200;
const canvasHeight = 630;

try {
  registerFont(path.resolve(__dirname, "./font/MPLUSRounded1c-Bold.ttf"), {
    family: "MplusRounded"
  });
} catch (err) {
  console.log(err);
}

const canvas = createCanvas(canvasWidth, canvasHeight);
const ctx = canvas.getContext("2d");

ctx.fillStyle = "white";
ctx.fillRect(0, 0, canvasWidth, canvasHeight);

ctx.font = '48px "MplusRounded"';
ctx.fillStyle = "black";
ctx.fillText(booksCount, 0, 100);
ctx.fillText(tsundokuPrice, 0, 200);
ctx.fillText(kidokuPrice, 0, 300);

const pngStream = canvas.createPNGStream();
const screenShotsPath = path.resolve(__dirname, "screenshots");
if (!fs.existsSync(screenShotsPath)) {
  fs.mkdirSync(screenShotsPath);
}
const out = fs.createWriteStream(path.resolve(screenShotsPath, shareId));
pngStream.pipe(out);
out.on("finish", () => {
  out.close();
});
