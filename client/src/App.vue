<template lang="pug">
  #app
    component(:is="$store.state.viewType + '-template'")
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import MobileTemplate from '@/components/templates/MobileTemplate.vue'
import DesktopTemplate from '@/components/templates/DesktopTemplate.vue'
import { ExStore } from 'vuex'

// TODO: move this to better position
const mobileThreshould = 750

@Component({
  components: {
    MobileTemplate,
    DesktopTemplate
  }
})
export default class App extends Vue {
  private windowWidth = 0
  public $store!: ExStore

  private handleResizeWindow() {
    const width = window.innerWidth
    if (this.$store.state.viewType !== 'mobile' && width < mobileThreshould) {
      this.$store.commit('setViewType', 'mobile')
    } else if (
      this.$store.state.viewType !== 'desktop' &&
      width >= mobileThreshould
    ) {
      this.$store.commit('setViewType', 'desktop')
    }
    this.windowWidth = width
  }

  private setLocale() {
    // ブラウザの言語設定を取得
    this.$store.commit('setLocale', window.navigator.language.split('-')[0])
  }

  mounted() {
    this.$nextTick(() => {
      this.handleResizeWindow()
    })
    this.setLocale()
    window.addEventListener('resize', this.handleResizeWindow)
    window.addEventListener('orientationchange', this.handleResizeWindow)
  }
}
</script>

<style lang="sass">
@import '@/style/global.sass'
@import '@/style/reset.sass'
</style>
