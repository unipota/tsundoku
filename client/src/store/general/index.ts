/* eslint-disable @typescript-eslint/explicit-function-return-type */
/* tslint:disable:no-shadowed-variable */

import Vue from 'vue'
import { Getters, Mutations, Actions } from 'vuex'
import { S, G, M, A } from './type'
import i18n from '@/i18n'
import api from './api'
import { BookSimpleRecord, BookRecord, BookStats } from '@/types/Book'
import { mockBooksMap } from './mockData'
import { tsundokuPrice } from '@/utils/tsundoku'
// ______________________________________________________
//
export const state = (): S => ({
  userId: '',
  locale: 'ja',
  viewType: 'desktop',
  currentTheme: 'light',
  themeSetting: 'auto',
  userConfirmed: false,
  userLogined: false,
  userScreenName: undefined,
  userIconUrl: undefined,
  userCreatedAt: undefined,
  booksLoaded: false,
  showMobileTopBar: true,
  showMobileTabBar: true,
  showDesktopNav: true,
  booksMap: {},
  readHistoriesMap: {},
  bookStatsArray: [],
  scrollLock: false,
  useMockBooksMap: false // 開発用
})
// ______________________________________________________
//
export const getters: Getters<S, G> = {
  getUserId(state) {
    return state.userId
  },
  getLocale(state) {
    return state.locale
  },
  viewTypeClass(state) {
    switch (state.viewType) {
      case 'mobile':
        return 'is-mobile'
      case 'desktop':
        return 'is-desktop'
    }
  },
  modalTransitionClass(state) {
    switch (state.viewType) {
      case 'mobile':
        return 'translate-y'
      case 'desktop':
        return 'translate-x'
    }
  },
  books(state): BookRecord[] {
    let books: BookRecord[]
    if (state.useMockBooksMap) {
      books = Object.values(mockBooksMap)
    } else {
      books = Object.values(state.booksMap)
    }
    return books
  },
  tsundokuBooks(_, getters) {
    return getters.books.filter(book => {
      if (book.totalPages === 0) {
        return true
      }
      return book.readPages < book.totalPages
    })
  },
  kidokuBooks(_, getters) {
    return getters.books.filter(book => {
      if (book.totalPages === 0) {
        return false
      }
      return book.readPages >= book.totalPages
    })
  },
  getBookById(_, getters) {
    return bookId => getters.books.find(book => book.id === bookId)
  },
  tsundokuPrice(_, getters) {
    return getters.tsundokuBooks.reduce(
      (sum, book) =>
        sum + tsundokuPrice(book.readPages, book.totalPages, book.price),
      0
    )
  },
  kidokuPrice(_, getters) {
    return getters.kidokuBooks.reduce((sum, book) => sum + book.price, 0)
  }
}

// ______________________________________________________
//
export const mutations: Mutations<S, M> = {
  setUserId(state, userId): void {
    state.userId = userId
  },
  setLocale(state, locale): void {
    state.locale = locale
    i18n.locale = locale
  },
  setViewType(state, viewType): void {
    state.viewType = viewType
  },
  setCurrentTheme(state, themeType): void {
    state.currentTheme = themeType
  },
  setUserConfirmed(state, userConfirmed): void {
    state.userConfirmed = userConfirmed
  },
  setUserLogined(state, userLogined): void {
    state.userLogined = userLogined
  },
  setUserScreenName(state, userScreenName): void {
    state.userScreenName = userScreenName
  },
  setUserIconUrl(state, userIconUrl): void {
    state.userIconUrl = userIconUrl
  },
  setUserCreatedAt(state, userCreatedAt): void {
    state.userCreatedAt = userCreatedAt
  },
  setBooksLoaded(state, booksLoaded): void {
    state.booksLoaded = booksLoaded
  },
  setShowMobileTopBar(state, showMobileTopBar): void {
    state.showMobileTopBar = showMobileTopBar
  },
  setShowMobileTabBar(state, showMobileTabBar): void {
    state.showMobileTabBar = showMobileTabBar
  },
  setShowDesktopNav(state, showDesktopNav): void {
    state.showDesktopNav = showDesktopNav
  },
  setBooksMap(state, booksArray) {
    state.booksMap = {}
    booksArray.forEach((book: BookRecord) => {
      state.booksMap = { ...state.booksMap, [book.id]: book }
    })
  },
  setReadHistoriesMap(state, bookDetail) {
    state.readHistoriesMap = {
      ...state.readHistoriesMap,
      [bookDetail.id]: bookDetail.readHistories
    }
  },
  setBookStats(state, bookStatsArray) {
    state.bookStatsArray = bookStatsArray
  },
  updateBook(state, { book }) {
    state.booksMap[book.id] = book
    state.booksMap = { ...state.booksMap }
  },
  setScrollLock(state, scrollLock) {
    state.scrollLock = scrollLock
  },
  deleteBook(state, bookId) {
    Vue.delete(state.booksMap, bookId)
  }
}
// ______________________________________________________
//
export const actions: Actions<S, A, G, M> = {
  whoAmI({ commit }): Promise<void> {
    return new Promise((resolve, reject) => {
      api
        .whoAmI()
        .then(result => {
          if (result.data) {
            const { logined, screenName, iconUrl, createdAt } = result.data
            commit('setUserLogined', logined)
            commit('setUserScreenName', screenName)
            commit('setUserIconUrl', iconUrl)
            commit('setUserCreatedAt', createdAt)
            commit('setUserConfirmed', true)
            resolve()
          }
        })
        .catch(err => {
          resolve()
        })
    })
  },
  userLogout(): Promise<void> {
    return new Promise((resolve, reject) => {
      api
        .userLogout()
        .then(result => {
          resolve()
          location.reload()
        })
        .catch(() => {
          reject()
        })
    })
  },
  getMyBooks({ state, commit }): Promise<void> {
    return new Promise((resolve, reject) => {
      if (state.useMockBooksMap) {
        // 開発用
        resolve()
      } else {
        api
          .getMyBooks()
          .then(result => {
            if (result.data) {
              commit('setBooksMap', result.data)
              resolve()
            }
          })
          .catch(err => {
            resolve()
          })
          .finally(() => {
            commit('setBooksLoaded', true)
          })
      }
    })
  },
  getBookDetail({ commit }, { id }) {
    return new Promise((resolve, reject) => {
      api.getBookDetail(id).then(result => {
        if (result.data) {
          commit('setReadHistoriesMap', result.data)
          resolve(result.data)
        }
      })
    })
  },
  searchBooksByISBN(_, { isbn }): Promise<BookRecord[]> {
    return new Promise((resolve, reject) => {
      api.searchBooksByISBN(isbn).then(result => {
        if (result.data) {
          resolve(result.data)
        }
      })
    })
  },
  searchBooks(_, { search }): Promise<BookSimpleRecord[]> {
    return new Promise((resolve, reject) => {
      api.searchBooks(search).then(result => {
        if (result.data) {
          resolve(result.data)
        }
      })
    })
  },
  addNewBook(_, { book }): Promise<BookRecord[]> {
    return new Promise((resolve, reject) => {
      api.addNewBook(book).then(result => {
        if (result.data) {
          resolve(result.data)
        }
      })
    })
  },
  updateBook({ commit }, { book }): Promise<BookRecord> {
    return new Promise((resolve, reject) => {
      api.updateBook(book.id, book).then(result => {
        if (result.data) {
          commit('updateBook', { book: result.data })
          resolve(result.data)
        }
      })
    })
  },
  async deleteBook({ commit }, { id }): Promise<BookRecord[]> {
    return new Promise(async (resolve, reject) => {
      api.deleteBook(id).then(result => {
        if (result) {
          commit('deleteBook', id)
          resolve()
        }
      })
    })
  },
  getAllBookStats({ commit }): Promise<BookStats[]> {
    return new Promise((resolve, reject) => {
      api.getAllBookStats().then(result => {
        if (result.data) {
          commit('setBookStats', result.data)
          resolve(result.data)
        }
      })
    })
  },
  getShareURL(): Promise<string> {
    return new Promise((resolve, reject) => {
      api
        .getShareURL()
        .then(result => {
          if (result.data) {
            resolve(result.data.URL)
          }
        })
        .catch(reason => reject(reason))
    })
  }
}

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
}
