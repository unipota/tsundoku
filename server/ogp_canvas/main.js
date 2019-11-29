const { registerFont, createCanvas, loadImage } = require("canvas");
const fs = require("fs");
const path = require("path");
const args = [...process.argv].slice(2);

if (args.length !== 4) throw Error("4 args required");
const [shareId, booksCount, tsundokuPrice, kidokuPrice] = args;
const canvasWidth = 1200;
const canvasHeight = 630;
const textBlack = "#4B4B4B";
const tsundokuRed = "#E3402A";
const kidokuBlue = "#2488D0";
// const toukeiBlackBg = "#e0e0e0";
// const tsundokuRedBg = "#ffdad5";
// const kidokuBlueBg = "#c3d4f0";

// Load Asset
const loadAssets = async () => {
  try {
    registerFont(path.resolve(__dirname, "font/MPLUSRounded1c-Bold.ttf"), {
      family: "MplusRounded"
    });
  } catch (err) {
    console.log(err);
  }

  const assets = {};
  assets.template = await loadImage(
    path.resolve(__dirname, "assets", "template.svg")
  );

  return assets;
};

// Draw Canvas
const draw = async () => {
  const assets = await loadAssets();
  const canvas = createCanvas(canvasWidth, canvasHeight);
  const ctx = canvas.getContext("2d");
  ctx.fillStyle = "white";
  ctx.fillRect(0, 0, canvasWidth, canvasHeight);

  ctx.drawImage(assets.template, 0, 0);

  ctx.font = '58px "MplusRounded"';
  ctx.fillStyle = textBlack;
  const booksTextMetrics = ctx.measureText(booksCount);
  ctx.fillText(booksCount, 124, 96 + 60);
  ctx.font = '24px "MplusRounded"';
  ctx.fillText("冊", 124 + booksTextMetrics.width + 12, 131 + 24);

  ctx.font = '28px "MplusRounded"';
  ctx.fillStyle = tsundokuRed;
  const tsundokuTextMetrics = ctx.measureText("¥");
  ctx.fillText("¥", 667, 131 + 24);
  ctx.font = '58px "MplusRounded"';
  ctx.fillText(
    Number(tsundokuPrice).toLocaleString("ja-JP"),
    667 + tsundokuTextMetrics.width + 12,
    96 + 60
  );

  ctx.font = '28px "MplusRounded"';
  ctx.fillStyle = kidokuBlue;
  const kidokuTextMetrics = ctx.measureText("¥");
  ctx.fillText("¥", 124, 402 + 24);
  ctx.font = '58px "MplusRounded"';
  ctx.fillText(
    Number(kidokuPrice).toLocaleString("ja-JP"),
    124 + kidokuTextMetrics.width + 12,
    368 + 60
  );

  return canvas;
};

// Output ImageFile
draw().then(canvas => {
  const pngStream = canvas.createPNGStream();
  const screenShotsPath = path.resolve(__dirname, "screenshots");
  if (!fs.existsSync(screenShotsPath)) {
    fs.mkdirSync(screenShotsPath);
  }
  const out = fs.createWriteStream(
    path.resolve(screenShotsPath, `${shareId}.png`)
  );
  pngStream.pipe(out);
  out.on("finish", () => {
    out.close();
  });
});
