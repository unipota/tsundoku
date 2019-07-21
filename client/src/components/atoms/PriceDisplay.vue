<template lang="pug">
  .body
    .currency-symbol
      | ¥
    .price-body
      | {{priceWithDelimiter}}
    .menu-button
      icon-down-arrow
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import IconDownArrow from '@/components/assets/IconDownArrow.vue'

@Component({ components: { IconDownArrow } })
export default class PriceDisplay extends Vue {
  @Prop({ type: Number, default: 0 }) readonly price!: number
  get priceWithDelimiter(): string {
    return this.price.toLocaleString()
  }
}
</script>

<style lang="sass" scoped>
.body
  display: flex
  box-sizing: border-box
  min-width: 180px
  max-width: 40vw
  height: 48px
  border:
    radius: 100vw
  background: var(--tsundoku-red-bg)
  color: var(--tsundoku-red)
  font:
    weight: bold
    size: 22px
  padding:
    left: 24px
    right: 18px

.currency-symbol
  display: flex
  align-items: center
  height: 100%
  flex-grow: 1

.price-body
  display: flex
  align-items: center
  height: 100%

.menu-button
  cursor: pointer
  height: 100%
  display: flex
  align-items: center
  padding:
    left: 8px
</style>
