// ______________________________________________________
//
// state
export interface S {
  userId: string
  locale: string
}
// ______________________________________________________
//
// getters
export interface G {
  getUserId: string
  getLocale: string
}
export interface RG {
  getUserId: G['getUserId']
  getLocale: G['getLocale']
}
// ______________________________________________________
//
// mutations
export interface M {
  setUserId: string
  setLocale: string
}
export interface RM {
  setUserId: M['setUserId']
  setLocale: M['setLocale']
}
// ______________________________________________________
//
// actions
// eslint-disable-next-line
export interface A {}
// eslint-disable-next-line
export interface RA {}
