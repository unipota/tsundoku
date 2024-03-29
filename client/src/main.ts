import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import i18n from './i18n'

import PortalVue from 'portal-vue'
import VueMeta from 'vue-meta'
import VTooltip from 'v-tooltip'

Vue.use(PortalVue)
Vue.use(VueMeta)
Vue.use(VTooltip, {
  defaultDelay: 500
})

Vue.config.productionTip = false

new Vue({
  router,
  store,
  i18n,
  // eslint-disable-next-line @typescript-eslint/explicit-function-return-type
  render: h => h(App)
}).$mount('#app')
