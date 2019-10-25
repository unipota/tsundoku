<template lang="pug">
  .mobile-top-bar
    .price-display
      portal-target(name="priceDisplay")
    setting-button.setting
    user-button.user(v-if="!userLogined")
    .user-icon-wrap(v-else)
      user-icon(:src="userIconUrl")
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import UserIcon from '@/components/atoms/UserIcon.vue'
import SettingButton from '@/components/atoms/SettingButton.vue'
import UserButton from '@/components/atoms/UserButton.vue'

@Component({
  components: { UserIcon, SettingButton, UserButton }
})
export default class MobileTopBar extends Vue {
  public $store!: ExStore

  get userLogined(): boolean {
    return this.$store.state.userLogined
  }
  get userIconUrl(): string | undefined {
    return this.$store.state.userIconUrl
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

.setting, .user
  display: flex
  align-items: center
  margin: 0 4px

.user-icon-wrap
  width: 28px
  height: 28px
  margin:
    left: 6px
    right: 2px
</style>
