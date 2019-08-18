import { BookRecord } from '@/types/Book'

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
  useMockBooksMap: boolean
}
// ______________________________________________________
//
// getters
export interface G {
  isLoggedIn: boolean
  isFirstLanding: boolean
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
  searchBooksByISBN: { isbn: string }
  searchBooks: { search: string }
  addNewBook: { book: BookRecord }
  updateBook: { book: BookRecord }
  deleteBook: { id: string }
}
export type RA = {
  [K in keyof A]: A[K]
}
