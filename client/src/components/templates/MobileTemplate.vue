<template lang="pug">
  div
    .top-bar-wrap
      mobile-top-bar
    .content-wrap
      keep-alive
        router-view
    .bottom-bar-wrap
      mobile-tab-bar
    floating-add-tsundoku-button(v-if="selectedPath === 'tsundoku'")
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import MobileTabBar from '@/components/atoms/MobileTabBar.vue'
import MobileTopBar from '@/components/molecules/MobileTopBar.vue'
import FloatingAddTsundokuButton from '@/components/atoms/FloatingAddTsundokuButton.vue'

@Component({
  components: {
    MobileTabBar,
    MobileTopBar,
    FloatingAddTsundokuButton
  }
})
export default class MobileTemplate extends Vue {
  get firstRouteName() {
    return this.$route.matched[0].path
  }
  get selectedPath() {
    return this.firstRouteName === ''
      ? 'tsundoku'
      : this.firstRouteName.slice(1)
  }
}
</script>

<style lang="sass">
.top-bar-wrap
  position: fixed
  z-index: 9999
  top: 0
  width: 100%

.content-wrap
  padding:
    top: 80px
  margin:
    bottom: 90px

.bottom-bar-wrap
  position: fixed
  z-index: 9999
  bottom: 0
  width: 100%
</style>
