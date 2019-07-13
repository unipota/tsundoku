/* tslint:disable:no-empty-interface */

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
export interface A {}
export interface RA {}
