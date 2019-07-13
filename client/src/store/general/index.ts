/* tslint:disable:no-shadowed-variable */

import { Getters, Mutations, Actions } from 'vuex'
import { S, G, M, A } from './type'
// ______________________________________________________
//
export const state = (): S => ({
  userId: ''
})
// ______________________________________________________
//
export const getters: Getters<S, G> = {
  getUserId(state) {
    return state.userId
  }
}
// ______________________________________________________
//
export const mutations: Mutations<S, M> = {
  setUserId(state, userId) {
    state.userId = userId
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
