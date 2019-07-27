import { BookRecord } from '@/types/Book'

type ViewType = 'mobile' | 'desktop'

// ______________________________________________________
//
// state
export interface S {
  userId: string
  locale: string
  viewType: ViewType
  booksMap: Record<string, BookRecord>
}
// ______________________________________________________
//
// getters
export interface G {
  getUserId: string
  getLocale: string
  books: BookRecord[]
  tsundokuBooks: BookRecord[]
  kidokuBooks: BookRecord[]
  getBookById: (bookId: string) => BookRecord
  tsundokuPrice: number
  kidokuPrice: number
}
// root getters has no namespace, so we can write root getter names like this
export type RG = {
  [K in keyof G]: G[K]
}
// ______________________________________________________
//
// mutations
export interface M {
  setUserId: string
  setLocale: string
  setViewType: ViewType
}
export type RM = {
  [K in keyof M]: M[K]
}
// ______________________________________________________
//
// actions
// eslint-disable-next-line
export interface A {
  searchBooksByISBN: { isbn: string }
}
// eslint-disable-next-line
export interface RA {
  searchBooksByISBN: A['searchBooksByISBN']
}
