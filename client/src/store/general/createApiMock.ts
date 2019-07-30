import MockAdapter from 'axios-mock-adapter'
import { AxiosInstance } from 'axios';

const bookRecords = [
  {
    "id": "",
    "isbn": "9784800274892",
    "title": "響け! ユーフォニアム北宇治高校吹奏学部、波乱の第二楽章前編",
    "author": [
      "武田綾乃"
    ],
    "totalPages": 387,
    "price": 630,
    "caption": "",
    "publisher": "宝島社",
    "coverImageUrl": "http://books.google.com/books/content?id=Q6H-swEACAAJ\u0026printsec=frontcover\u0026img=1\u0026zoom=1\u0026source=gbs_api",
    "readPages": 0,
    "memo": ""
  },
  {
    "id": "",
    "isbn": "9784800274915",
    "title": "響け! ユーフォニアム北宇治高校吹奏学部、波乱の第二楽章後編",
    "author": [
      "武田綾乃"
    ],
    "totalPages": 387,
    "price": 630,
    "caption": "",
    "publisher": "宝島社",
    "coverImageUrl": "http://books.google.com/books/content?id=4BlLswEACAAJ\u0026printsec=frontcover\u0026img=1\u0026zoom=1\u0026source=gbs_api",
    "readPages": 0,
    "memo": ""
  },
  {
    "id": "",
    "isbn": "9784101800042",
    "title": "いなくなれ、群青",
    "author": [
      "河野裕"
    ],
    "totalPages": 318,
    "price": 590,
    "caption": "",
    "publisher": "新潮社",
    "coverImageUrl": "http://books.google.com/books/content?id=Q17FoQEACAAJ\u0026printsec=frontcover\u0026img=1\u0026zoom=1\u0026source=gbs_api",
    "readPages": 0,
    "memo": ""
  },
  {
    "id": "",
    "isbn": "9784832249905",
    "title": "ご注文はうさぎですか? 7",
    "author": [
      "Koi"
    ],
    "totalPages": 118,
    "price": 819,
    "caption": "",
    "publisher": "芳文社",
    "coverImageUrl": "",
    "readPages": 0,
    "memo": ""
  },
  {
    "id": "",
    "isbn": "9784839941062",
    "title": "プログラミングコンテストチャレンジブック [第2版]",
    "author": [
      "秋葉拓哉",
      "岩田陽一",
      "北川宜稔"
    ],
    "totalPages": 367,
    "price": 0,
    "caption": "",
    "publisher": "マイナビ出版",
    "coverImageUrl": "http://books.google.com/books/content?id=s40pvgbtOZ8C\u0026printsec=frontcover\u0026img=1\u0026zoom=1\u0026edge=curl\u0026source=gbs_api",
    "readPages": 0,
    "memo": ""
  },
]

export const createApiMock = (client: AxiosInstance) => {
  const mock = new MockAdapter(client);

  bookRecords.forEach(book => {
    mock.onGet('api/search/isbn', { params: { isbn: book.isbn } }).reply(200, [ book ])
  })

}