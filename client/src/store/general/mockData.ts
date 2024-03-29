import { BookRecord } from '@/types/Book'

export const mockBooksMap: Record<string, BookRecord> = {
  mock0: {
    id: 'mock0',
    isbn: '9784101800042',
    title: 'いなくなれ、群青',
    author: ['河野裕'],
    totalPages: 318,
    price: 637,
    caption:
      '１１月１９日午前６時４２分、僕は彼女に再会した。誰よりも真っ直ぐで、正しく、凛々しい少女、真辺由宇。あるはずのない出会いは、安定していた僕の高校生活を一変させる。奇妙な島。連続落書き事件。そこに秘められた謎…。僕はどうして、ここにいるのか。彼女はなぜ、ここに来たのか。やがて明かされる真相は、僕らの青春に残酷な現実を突きつける。「階段島」シリーズ、開幕。',
    publisher: '新潮社',
    coverImageUrl:
      'https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/0042/9784101800042.jpg?_ex=200x200',
    readPages: 200,
    memo: '',
    createdAt: '',
    updatedAt: ''
  },
  mock1: {
    id: 'mock1',
    isbn: '9784832249905',
    title: 'ご注文はうさぎですか？　7',
    author: ['Koi'],
    totalPages: 119,
    price: 884,
    caption: null,
    publisher: '',
    coverImageUrl:
      'https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/9905/9784832249905.jpg?_ex=200x200',
    readPages: 119,
    memo: '',
    createdAt: '',
    updatedAt: ''
  },
  mock2: {
    id: 'mock2',
    isbn: '9784839941062',
    title:
      'プログラミングコンテストチャレンジブック第2版 問題解決のアルゴリズム活用力とコーディングテクニッ',
    author: ['秋葉拓哉', '岩田陽一'],
    totalPages: 367,
    price: 3542,
    caption:
      'プログラミングコンテストの問題を通してアルゴリズムのしくみや考え方を楽しく習得。世界トップレベルの著者たちがコンテストで得た知識やノウハウを難易度別にまとめました。現役プログラマだけでなくプログラマを目指している方にもぜひ読んでいたただきたい１冊。',
    publisher: 'マイナビ出版',
    coverImageUrl:
      'https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/1062/9784839941062.jpg?_ex=200x200',
    readPages: 20,
    memo: '',
    createdAt: '',
    updatedAt: ''
  },
  mock3: {
    id: 'mock3',
    isbn: '9784839941062',
    title:
      'プログラミングコンテストチャレンジブック第2版 問題解決のアルゴリズム活用力とコーディングテクニッ',
    author: ['秋葉拓哉', '岩田陽一'],
    totalPages: 367,
    price: 3542,
    caption:
      'プログラミングコンテストの問題を通してアルゴリズムのしくみや考え方を楽しく習得。世界トップレベルの著者たちがコンテストで得た知識やノウハウを難易度別にまとめました。現役プログラマだけでなくプログラマを目指している方にもぜひ読んでいたただきたい１冊。',
    publisher: 'マイナビ出版',
    coverImageUrl:
      'https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/1062/9784839941062.jpg?_ex=200x200',
    readPages: 20,
    memo: '',
    createdAt: '',
    updatedAt: ''
  }
}
