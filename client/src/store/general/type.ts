import {
  BookSimpleRecord,
  BookRecord,
  ReadHistory,
  BookDetail,
  BookStats
} from '@/types/Book'

type ViewType = 'mobile' | 'desktop'
type ViewTypeClass = 'is-mobile' | 'is-desktop'
type ModalTransitionClass = 'translate-y' | 'translate-x'
type ThemeType = 'light' | 'dark'
type ThemeSetting = 'light' | 'dark' | 'auto'

// ______________________________________________________
//
// state
export interface S {
  userId: string
  locale: string
  viewType: ViewType
  currentTheme: ThemeType
  themeSetting: ThemeSetting
  userConfirmed: boolean
  userLogined: boolean
  userScreenName: string | undefined
  userIconUrl: string | undefined
  userCreatedAt: string | undefined
  booksLoaded: boolean
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
  setCurrentTheme: ThemeType
  setUserConfirmed: boolean
  setUserLogined: boolean
  setUserScreenName: string
  setUserIconUrl: string
  setUserCreatedAt: string
  setBooksLoaded: boolean
  setShowMobileTopBar: boolean
  setShowMobileTabBar: boolean
  setShowDesktopNav: boolean
  setBooksMap: BookRecord[]
  setReadHistoriesMap: BookDetail
  setBookStats: BookStats[]
  updateBook: { book: BookRecord }
  deleteBook: string
}
export type RM = {
  [K in keyof M]: M[K]
}
// ______________________________________________________
//
// actions
export interface A {
  whoAmI: {}
  userLogout: {}
  getMyBooks: {}
  getBookDetail: { id: string }
  searchBooksByISBN: { isbn: string }
  searchBooks: { search: string }
  addNewBook: { book: BookSimpleRecord }
  updateBook: { book: BookRecord }
  deleteBook: { id: string }
  getAllBookStats: {}
}
export type RA = {
  [K in keyof A]: A[K]
}
