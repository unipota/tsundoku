/* eslint-disable @typescript-eslint/explicit-function-return-type */
/* tslint:disable:no-shadowed-variable */

// import Vue from 'vue'
import { Getters, Mutations, Actions } from 'vuex'
import { S, G, M, A } from './type'
// ______________________________________________________
//
export const state = (): S => ({
  stack: []
})
// ______________________________________________________
//
export const getters: Getters<S, G> = {
  currentModalName(state) {
    return state.stack.length > 0 ? state.stack[state.stack.length - 1] : ''
  }
}

// ______________________________________________________
//
export const mutations: Mutations<S, M> = {
  push(state, { name }) {
    state.stack.push(name)
  },
  pop(state) {
    state.stack.pop()
  },
  popAll(state) {
    state.stack = []
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
