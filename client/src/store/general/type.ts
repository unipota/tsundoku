import { BookRecord } from '@/types/Book'

type ViewType = 'mobile' | 'desktop'

// ______________________________________________________
//
// state
export interface S {
  userId: string
  locale: string
  viewType: ViewType
  books: BookRecord[]
}
// ______________________________________________________
//
// getters
export interface G {
  getUserId: string
  getLocale: string
  getTsundoku: BookRecord[]
  getKidoku: BookRecord[]
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
export interface A {}
// eslint-disable-next-line
export interface RA {}
