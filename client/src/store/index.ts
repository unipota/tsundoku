import Vue from 'vue'
import Vuex, { Store, RootState, ExStore } from 'vuex'

import general from './general'

Vue.use(Vuex)

const store: ExStore = new Store<RootState>({
  ...general,
  modules: {},
} as any)

export default store
