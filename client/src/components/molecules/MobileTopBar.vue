<template lang="pug">
  .mobile-top-bar
    portal-target(name="priceDisplay")
    user-icon
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import UserIcon from '@/components/atoms/UserIcon.vue'

@Component({
  components: { UserIcon }
})
export default class MobileTopBar extends Vue {
  public $store!: ExStore
}
</script>

<style lang="sass">
.mobile-top-bar
  display: flex
  justify-content: space-between
  padding:
    top: 8px
    left: 16px
    right: 16px
    bottom: 8px
  background: rgba(255,255,255,1)
  box-shadow: 0 0px 6px -4px var(--text-gray)
  // backdrop-filter: blur(2px)
</style>
