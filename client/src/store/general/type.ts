// ______________________________________________________
//
// state
export interface S {
  userId: string
}
// ______________________________________________________
//
// getters
export interface G {
  getUserId: string
}
export interface RG {
  getUserId: G['getUserId']
}
// ______________________________________________________
//
// mutations
export interface M {
  setUserId: string
}
export interface RM {
  setUserId: M['setUserId']
}
// ______________________________________________________
//
// actions
// eslint-disable-next-line
export interface A {}
// eslint-disable-next-line
export interface RA {}
