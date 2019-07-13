export interface BookRecord {
  id: string
  isbn: string
  title: string
  author: string
  totalPages: number
  price: number // 定価
  caption: string | null
  publisher: string
  coverImageUrl: string | null
  readPages: number
  memo: string | null
}
