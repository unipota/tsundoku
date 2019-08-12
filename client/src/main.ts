import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import i18n from './i18n'

import PortalVue from 'portal-vue'
import vClickOutside from 'v-click-outside'
import VueMeta from 'vue-meta'

Vue.use(PortalVue)
Vue.use(vClickOutside)
Vue.use(VueMeta)

Vue.config.productionTip = false

new Vue({
  router,
  store,
  i18n,
  // eslint-disable-next-line @typescript-eslint/explicit-function-return-type
  render: h => h(App)
}).$mount('#app')
