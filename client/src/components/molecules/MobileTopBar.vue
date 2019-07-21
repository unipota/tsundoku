<template lang="pug">
  .mobile-top-bar
    price-display(:price="114514")
    user-icon
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import PriceDisplay from '@/components/atoms/PriceDisplay.vue'
import UserIcon from '@/components/atoms/UserIcon.vue'

@Component({
  components: { PriceDisplay, UserIcon }
})
export default class MobileTopBar extends Vue {}
</script>

<style lang="sass">
.mobile-top-bar
  display: flex
  justify-content: space-between
  padding:
    top: 24px
    left: 16px
    right: 16px
</style>
