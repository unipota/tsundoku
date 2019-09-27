export interface BookSimpleRecord {
  id: string
  isbn: string
  title: string
  author: string[]
  price: number
  caption: string | null
  publisher: string
  coverImageUrl: string | null
}

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

export interface ReadHistory {
  readPages: number
  createdAt: string
}

export interface BookDetail {
  id: string
  isbn: string
  title: string
  author: string[]
  totalPages: number
  price: number // 定価
  caption: string | null
  publisher: string
  coverImageUrl: string | null
  readHistories: ReadHistory[]
  memo: string | null
  createdAt: string
  updatedAt: string
}

export interface BookStats {
  id: string
  title: string
  totalPages: number
  price: number
  readHistories: ReadHistory[]
}
