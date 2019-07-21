<template lang="pug">
  section.main-nav
    .brand
      Icon-Logo
    desktop-tab(:selected-tab="$route.name")
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import DesktopTab from '@/components/atoms/DesktopTab.vue'
import IconLogo from '@/components/assets/IconLogo.vue'

@Component({
  components: {
    DesktopTab,
    IconLogo
  }
})
export default class DesktopTabBar extends Vue {}
</script>

<style lang="sass" scoped>
.main-nav
  width: 320px
  height: 100%
  padding: 24px
.brand
  display: flex
  justify-content: flex-start
  align-items: center
  height: 40px
  width: 100%
  margin-bottom: 24px
  padding: 0 32px
</style>
