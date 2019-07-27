<template lang="pug">
  .body(:style="bodyStyle")
    .currency-symbol
      | ¥
    .price-body
      | {{priceWithDelimiter}}
    .menu-button
      icon(name="down-arrow" :color="`var(--${currentColor})`")
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

import Icon from '@/components/assets/Icon.vue'

@Component({ components: { Icon } })
export default class PriceDisplay extends Vue {
  @Prop({ type: Number, default: 0 })
  readonly price!: number

  @Prop({ type: Boolean, default: false })
  readonly tsundoku!: boolean

  @Prop({ type: Boolean, default: false })
  readonly kidoku!: boolean

  get priceWithDelimiter(): string {
    return `${this.tsundoku ? '-' : ''}${this.price.toLocaleString()}`
  }

  get bodyStyle() {
    return {
      background: `var(--${this.currentColor}-bg)`,
      color: `var(--${this.currentColor})`
    }
  }

  get currentColor(): string | undefined {
    return this.tsundoku
      ? 'tsundoku-red'
      : this.kidoku
      ? 'kidoku-blue'
      : undefined
  }
}
</script>

<style lang="sass" scoped>
.body
  display: flex
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
  transition: background .5s, color .5s

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
  transition: all .5s
</style>
