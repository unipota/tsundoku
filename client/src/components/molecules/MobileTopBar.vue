<template lang="pug">
  .mobile-top-bar
    .price-display
      portal-target(name="priceDisplay")
    user-button.user
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import UserIcon from '@/components/atoms/UserIcon.vue'
import UserButton from '@/components/atoms/UserButton.vue'

@Component({
  components: { UserIcon, UserButton }
})
export default class MobileTopBar extends Vue {
  public $store!: ExStore
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
