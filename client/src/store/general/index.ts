/* tslint:disable:no-shadowed-variable */

import { Getters, Mutations, Actions } from 'vuex'
import { S, G, M, A } from './type'
import i18n from '@/i18n'
import { BookRecord } from '@/types/Book'
// ______________________________________________________
//
export const state = (): S => ({
  userId: '',
  locale: 'ja',
  books: []
})
// ______________________________________________________
//
export const getters: Getters<S, G> = {
  getUserId(state): string {
    return state.userId
  },
  getLocale(state): string {
    return state.locale
  },
  getTsundoku(state): BookRecord[] {
    return state.books.filter(
      (book): boolean => book.readPages < book.totalPages
    )
  },
  getKidoku(state): BookRecord[] {
    return state.books.filter(
      (book): boolean => book.readPages >= book.totalPages
    )
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
  }
}
// ______________________________________________________
//
export const actions: Actions<S, A, G, M> = {}

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
}
