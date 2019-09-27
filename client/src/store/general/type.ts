import { BookRecord, ReadHistory, BookDetail, BookStats } from '@/types/Book'

type ViewType = 'mobile' | 'desktop'
type ViewTypeClass = 'is-mobile' | 'is-desktop'
type ModalTransitionClass = 'translate-y' | 'translate-x'

// ______________________________________________________
//
// state
export interface S {
  userId: string
  locale: string
  viewType: ViewType
  showMobileTopBar: boolean
  showMobileTabBar: boolean
  showDesktopNav: boolean
  booksMap: Record<string, BookRecord>
  readHistoriesMap: Record<string, ReadHistory[]>
  bookStatsArray: BookStats[]
  useMockBooksMap: boolean
}
// ______________________________________________________
//
// getters
export interface G {
  getUserId: string
  getLocale: string
  viewTypeClass: ViewTypeClass
  modalTransitionClass: ModalTransitionClass
  books: BookRecord[]
  tsundokuBooks: BookRecord[]
  kidokuBooks: BookRecord[]
  getBookById: (bookId: string) => BookRecord | undefined
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
  setShowMobileTopBar: boolean
  setShowMobileTabBar: boolean
  setShowDesktopNav: boolean
  setBooksMap: BookRecord[]
  setReadHistoriesMap: BookDetail
  setBookStats: BookStats[]
  updateBook: { book: BookRecord }
}
export type RM = {
  [K in keyof M]: M[K]
}
// ______________________________________________________
//
// actions
export interface A {
  getMyBooks: {}
  getBookDetail: { id: string }
  searchBooksByISBN: { isbn: string }
  searchBooks: { search: string }
  addNewBook: { book: BookRecord }
  updateBook: { book: BookRecord }
  deleteBook: { id: string }
  getAllBookStats: {}
}
export type RA = {
  [K in keyof A]: A[K]
}
