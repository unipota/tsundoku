export interface BookRecord {
  id: string
  isbn: string
  title: string
  author: string[]
  totalPages: number
  price: number // 定価
  caption: string | null
  publisher: string
  coverImageUrl: string | null
  readPages: number
  memo: string | null
  createdAt: string
  updatedAt: string
}

export type BookSimpleRecord = Omit<BookRecord, 'createdAt' | 'updatedAt'>

export interface ReadHistory {
  readPages: number
  createdAt: string
}

export interface BookDetail extends BookRecord {
  readHistories: ReadHistory[]
}

export interface BookStats {
  id: string
  title: string
  totalPages: number
  price: number
  readHistories: ReadHistory[]
}
