<template lang="pug">
  .progress-knob(
    :style="bodyStyle"
    @mouseleave="handleMouseleave"
    @mousemove.prevent="throttledMousemove"
    @touchmove.prevent="throttledTouchmove"
    @mouseup="handleMouseup"
    @touchend="handleMouseup"
    )
    .progress-knob__body(
      :style="positionByRatio"
      @mousedown.prevent="handleMousedown"
      @touchstart.prevent="handleTouchstart"
      )
      icon(name="bookmark" :width="24" :height="24")
</template>

<script lang="ts">
import { Vue, Component, Prop, Watch } from 'vue-property-decorator'
import { throttle } from 'lodash'

import Icon from '@/components/assets/Icon.vue'

@Component({
  components: { Icon }
})
export default class ProgressKnob extends Vue {
  @Prop({ type: Number, required: true })
  private readonly editedReadPages!: number

  @Prop({ type: Number, required: true })
  private readonly totalPages!: number

  @Prop({ type: Number, required: true })
  private readonly progressBarWidth!: number

  moveX: number =
    this.progressBarWidth * (this.editedReadPages / this.totalPages)
  isDrugging: boolean = false
  currentMouseX: number = 0

  get positionByRatio() {
    return {
      transform: `translateX(-50%) translateX(${this.limitedPosX}px)`
    }
  }

  get limitedPosX() {
    return Math.min(Math.max(this.moveX, 0), this.progressBarWidth)
  }

  get bodyStyle() {
    return {
      cursor: this.isDrugging ? 'ew-resize' : 'auto'
    }
  }

  handleTouchstart(e: TouchEvent) {
    this.isDrugging = true
    this.currentMouseX = e.touches[0].clientX
  }

  handleMousedown(e: MouseEvent) {
    this.isDrugging = true
    this.currentMouseX = e.x
  }

  handleMousemove(e: MouseEvent) {
    if (!this.isDrugging) return
    this.moveX += e.x - this.currentMouseX
    this.currentMouseX = e.x
  }

  get throttledMousemove() {
    return throttle(this.handleMousemove, 100)
  }

  handleTouchmove(e: TouchEvent) {
    if (!this.isDrugging) return
    this.moveX += e.touches[0].clientX - this.currentMouseX
    this.currentMouseX = e.touches[0].clientX
  }

  get throttledTouchmove() {
    return throttle(this.handleTouchmove, 100)
  }

  handleMouseup() {
    this.isDrugging = false
  }

  handleMouseleave() {
    this.isDrugging = false
  }

  @Watch('editedReadPages')
  onChangeEditedReadPages(value: number) {
    this.moveX = this.progressBarWidth * (value / this.totalPages)
  }

  @Watch('limitedPosX')
  onChangeMoveX(value: number) {
    this.$emit(
      'update:editedReadPages',
      Math.round((value / this.progressBarWidth) * this.totalPages)
    )
  }
}
</script>

<style lang="sass" scoped>
$triangle-height: 12px
$triangle-width: 6px

.progress-knob
  display: flex
  padding:
    left: calc(25% + 2px)
    right: 25%

  &__body
    display: flex
    align-items: center
    justify-content: center
    flex:
      shrink: 1
    position: relative
    min-width: 40px
    min-height: 40px
    background: var(--kidoku-blue)
    border:
      radius: 8px
    padding: 4px
    margin:
      bottom: $triangle-height - 2px
    cursor: ew-resize
    color: white
    box-shadow: 0 2px 6px -1px var(--kidoku-blue)
    transition: transform .1s linear

    &::before
      position: absolute
      bottom: 0
      transform: translateY(100%)
      left: 0
      right: 0
      content: ''
      display: block
      margin: auto
      width: 0
      height: 0
      border:
        style: solid
        width: $triangle-height
        color: var(--kidoku-blue)
        bottom:
          color: transparent
        left:
          width: $triangle-width
          color: transparent
        right:
          width: $triangle-width
          color: transparent
</style>
