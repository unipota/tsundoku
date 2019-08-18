<template lang="pug">
  section.main-nav
    .brand
      icon(name="logo")
    .tab-wrap
      desktop-tab(:selected-tab="selectedPath" :price="price")
    .button-wrap
      router-link(:to="`${firstRouteName}/add-books-search`")
        desktop-nav-button(:text="$t('add_by_searching')")
          icon(name="search" color="white" :height="24" :width="24")
    .button-wrap
      router-link(:to="`${firstRouteName}/add-books-scan`")
        desktop-nav-button(:text="$t('add_by_scanning')")
          icon(name="scanner" color="white" :height="24" :width="24")
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

import DesktopTab from '@/components/atoms/DesktopTab.vue'
import DesktopNavButton from '@/components/atoms/DesktopNavButton.vue'
import Icon from '@/components/assets/Icon.vue'

@Component({
  components: {
    Icon,
    DesktopTab,
    DesktopNavButton
  }
})
export default class DesktopNav extends Vue {
  get firstRouteName() {
    return this.$route.matched[0].path
  }
  get selectedPath() {
    return this.firstRouteName === ''
      ? 'tsundoku'
      : this.firstRouteName.slice(1)
  }
  get price(): number {
    return this.selectedPath === 'tsundoku'
      ? this.$store.getters.tsundokuPrice
      : this.selectedPath === 'kidoku'
      ? this.$store.getters.kidokuPrice
      : 0
  }
}
</script>

<style lang="sass" scoped>
.main-nav
  width: 100%
  height: 100%
.brand
  display: flex
  justify-content: flex-start
  align-items: center
  height: 40px
  width: 100%
  margin-bottom: 24px
  padding: 0 32px
.tab-wrap
  margin:
    top: 32px
    bottom: 48px
.button-wrap
  padding: 0 24px
  margin: 16px 0
</style>
