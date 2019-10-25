// ______________________________________________________
//
// state
export interface S {
  stack: string[]
}
// ______________________________________________________
//
// getters
export interface G {
  currentModalName: string
}
// root getters has no namespace, so we can write root getter names like this
export interface RG {
  'modal/currentModalName': G['currentModalName']
}
// ______________________________________________________
//
// mutations
export interface M {
  push: { name: string }
  pop: {}
  popAll: {}
}
export interface RM {
  'modal/push': M['push']
  'modal/pop': M['pop']
  'modal/popAll': M['popAll']
}
// ______________________________________________________
//
// actions
export interface A {}
export interface RA {}
