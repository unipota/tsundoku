const h1 = document.getElementById('midashi');
h1.textContent = '変化後';

function ogpSetting(url) {
  const headDeta = document.head.children;

  for (let i = 0; i < headDeta.length; i++) {
    const propertyVal = headDeta[i].getAttribute('property');
    if (propertyVal != null) {
      if (propertyVal.indexOf('og:image') != -1) {
        headDeta[i].setAttribute('content', url);
      }
    }
  }
}

ogpSetting('変更後のURL2');
