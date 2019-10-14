<template lang="pug">
  span.tweened-number
    | {{formattedNumber}}
</template>

<script lang="ts">
import { Vue, Component, Prop, Watch } from 'vue-property-decorator'
import * as TWEEN from '@tweenjs/tween.js'

const animate = () => {
  if (TWEEN.update()) {
    requestAnimationFrame(animate)
  }
}

@Component({
  components: {}
})
export default class TweenedNumber extends Vue {
  @Prop({ type: Number, required: true })
  private num!: number

  @Prop({ type: Boolean, default: false })
  private float!: boolean

  @Prop({ type: Number, default: 300 })
  private duration!: number

  @Prop({ type: Boolean, default: false })
  private formatLocal!: boolean

  tweenedNumberObject = { num: 0 }

  created() {
    this.tweenedNumberObject.num = this.num
  }

  get formattedNumber(): string {
    let num = this.tweenedNumberObject.num
    if (!this.float) {
      num = Math.round(num)
    }
    if (this.formatLocal) {
      return num.toLocaleString('ja-JP')
    } else {
      return num.toString()
    }
  }

  @Watch('num')
  onNumChanged(value: number) {
    const tween = new TWEEN.Tween(this.tweenedNumberObject).to(
      { num: value },
      this.duration
    )
    tween.easing(TWEEN.Easing.Quartic.InOut)
    tween.start()
    animate()
  }
}
</script>

<style lang="sass" scoped></style>
