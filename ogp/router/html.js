exports.buildOgpHtmlHandler = function buildOgpHtmlHandler (req, res){
  const tsundoku = req.query.tsundoku;
  const kidoku = req.query.kidoku;
  const count = req.query.count;

  res.render('/tsundoku/ogp/public/ogpimage.ejs', {
    h1: "ツンドク",
    tsundoku: tsundoku,
    kidoku: kidoku,
    count: count
  });
}
