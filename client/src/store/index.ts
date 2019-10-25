import Vue from 'vue'
import Vuex, { Store, RootState, ExStore } from 'vuex'

import general from './general'
import modal from './modal'

Vue.use(Vuex)

const store: ExStore = new Store<RootState>({
  ...general,
  modules: { modal }
} as any) // eslint-disable-line @typescript-eslint/no-explicit-any

export default store
