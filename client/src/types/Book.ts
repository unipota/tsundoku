export interface BookRecord {
  id: string,
  isbn: string,
  title: string,
  author: string,
  totalPages: number,
  regularPrice: number, // 定価
  caption: string | null,
  publisher: string,
  coverImageUrl: string | null
}

export interface BookUserRecord extends BookRecord {
  readPages: number,
  memo: string | null,
  purchacedPrice: number, // 購入した価格
}
