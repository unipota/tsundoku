<template lang="pug">
  .mobile-top-bar
    .price-display
      portal-target(name="priceDisplay")
    .right-side-wrapper
      .logo-icon(@click="openAboutModal")
        icon(name="logo" :width="28" :height="28")
      user-button.user
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import Icon from '@/components/assets/Icon.vue'
import UserIcon from '@/components/atoms/UserIcon.vue'
import UserButton from '@/components/atoms/UserButton.vue'

@Component({
  components: { Icon, UserIcon, UserButton }
})
export default class MobileTopBar extends Vue {
  public $store!: ExStore

  openAboutModal() {
    this.$store.commit('modal/push', { name: 'about' })
  }
}
</script>

<style lang="sass">
.mobile-top-bar
  display: flex
  justify-content: flex-end
  align-items: center
  height: 64px
  padding:
    top: 8px
    left: 16px
    right: 16px
    bottom: 8px
  background: rgba(255,255,255,1)
  box-shadow: 0 0px 6px -4px var(--text-gray)

.price-display
  margin:
    right: auto

.right-side-wrapper
  display: flex
  align-items: center

.logo-icon
  width: 28px
  height: 28px

.user
  display: flex
  align-items: center
  margin:
    left: 16px
    right: 4px
</style>
