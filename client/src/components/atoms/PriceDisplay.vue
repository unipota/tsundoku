<template lang="pug">
  .body(:style="bodyStyle")
    .currency-symbol ¥
    .price-body {{priceWithDelimiter}}
    .menu-button
      icon(name="down-arrow" :color="`var(--${currentColor})`")
</template>

<script lang="ts">
import { Vue, Component, Prop, Watch } from 'vue-property-decorator'
import TWEEN from '@tweenjs/tween.js'

import Icon from '@/components/assets/Icon.vue'

const animate = () => {
  if (TWEEN.update()) {
    requestAnimationFrame(animate)
  }
}

@Component({ components: { Icon } })
export default class PriceDisplay extends Vue {
  @Prop({ type: Number, default: 0 })
  readonly price!: number

  @Prop({ type: Boolean, default: false })
  readonly tsundoku!: boolean

  @Prop({ type: Boolean, default: false })
  readonly kidoku!: boolean

  tweenedPriceNumberObject = { price: 0 }

  get priceWithDelimiter(): string {
    return `${
      this.tsundoku && this.tweenedPriceNumberObject.price !== 0 ? '-' : ''
    }${Math.floor(this.tweenedPriceNumberObject.price).toLocaleString()}`
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

  @Watch('price')
  onPriceChanged(val: number) {
    const tween = new TWEEN.Tween(this.tweenedPriceNumberObject).to(
      { price: val },
      300
    )
    tween.easing(TWEEN.Easing.Quartic.InOut)
    tween.start()
    animate()
  }

  created() {
    this.tweenedPriceNumberObject.price = this.price
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

.currency-symbol
  display: flex
  align-items: center
  height: 100%
  flex-grow: 1

.price-body
  position: relative
  overflow: hidden
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
